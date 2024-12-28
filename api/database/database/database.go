package database

import (
	"fmt"
	"pier/api/database/types"
	"pier/lib"
	"pier/storage"
	"strings"
)

func GetDatabaseIndex() ([]types.DatabaseIndex, error) {
	db := storage.DB()
	rows, err := db.Query("SELECT * FROM `databases` ORDER BY `database`, `collection` ASC")
	if err != nil {
		return nil, err
	}
	var databaseIndex []types.DatabaseIndex
	for rows.Next() {
		var database types.DatabaseIndex
		err := rows.Scan(&database.Database, &database.Collection, &database.Facets)
		if err != nil {
			return nil, err
		}
		databaseIndex = append(databaseIndex, database)
	}
	return databaseIndex, nil
}

func GetDocs(database string, collection string, where string, order string, offset int) ([]types.Doc, error) {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s ORDER BY %s LIMIT %d OFFSET %d", table, where, order, 100, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	docs, err := MapDocs(rows)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func GetFacets(database string, collection string) (string, error) {
	db := storage.DB()
	var facets string
	err := db.QueryRow("SELECT `facets` FROM `databases` WHERE `database`=? AND `collection`=?", database, collection).Scan(&facets)
	return facets, err
}

func UpdateFacets(database string, collection string, facets string) error {
	db := storage.DB()
	_, err := db.Exec("UPDATE `databases` SET `facets`=? WHERE `database`=? AND `collection`=?", facets, database, collection)
	if err != nil {
		return err
	}
	return err
}

func GetDocsCount(database string, collection string, where string) (int, error) {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s", table, where)
	var count int
	err := db.QueryRow(query).Scan(&count)
	return count, err
}

func CreateDoc(database string, collection string, record types.Doc) (int64, error) {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	docKeys := []string{}
	values := []interface{}{}
	for key := range record {
		docKeys = append(docKeys, "`"+key+"`")
		value := fmt.Sprintf("%v", record[key])
		values = append(values, value)
	}
	columns := strings.Join(docKeys, ", ")
	qms := []string{}
	for range docKeys {
		qms = append(qms, "?")
	}
	placeholders := strings.Join(qms, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	res, err := db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAutocompletes(database string, collection string, fields types.AutocompleteFields) (types.Autocompletes, error) {
	db := storage.DB()
	autocompletes := types.Autocompletes{}
	table := fmt.Sprintf("database_%s_%s", database, collection)
	for _, field := range fields.Selects {
		query := fmt.Sprintf("SELECT DISTINCT(`%s`) FROM `%s` ORDER BY `%s` ASC", field, table, field)
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		var values []string
		for rows.Next() {
			var value string
			err := rows.Scan(&value)
			if err != nil {
				return nil, err
			}
			if value != "" {
				values = append(values, value)
			}
		}
		autocompletes[field] = values
	}
	for _, field := range fields.Tags {
		query := fmt.Sprintf("SELECT DISTINCT tags.tag from `%s` cross join json_table(`%s`.`%s`, '$[*]' COLUMNS (tag varchar(255) path '$')) tags order by tags.tag ASC", table, table, field)
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		var values []string
		for rows.Next() {
			var value string
			err := rows.Scan(&value)
			if err != nil {
				return nil, err
			}
			if value != "" {
				values = append(values, value)
			}
		}
		autocompletes[field] = values
	}
	return autocompletes, nil
}

func GetDoc(database string, collection string, id int64) (types.Doc, error) {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT * FROM %s WHERE `id`=?", table)
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	docs, err := MapDocs(rows)
	if err != nil {
		return nil, err
	}
	return docs[0], nil
}

func GetDuplicate(database string, collection string, id int64, name string) (bool, error) {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE `id`!=? AND `name`=?", table)
	row := db.QueryRow(query, id, name)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func UpdateDoc(database string, collection string, id int64, doc types.Doc) error {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	sets := []string{}
	values := []interface{}{}
	for key := range doc {
		sets = append(sets, "`"+key+"`=?")
		value := doc[key]
		valueString := fmt.Sprintf("%v", value)
		if strings.HasPrefix(valueString, "🕸") && strings.HasSuffix(valueString, "🕸") {
			url := valueString[4 : len(valueString)-4]
			image, err := lib.Download(url)
			if err != nil {
				return err
			}
			value = lib.ResizeImage(image, 256, 256, 70)
		}
		values = append(values, value)
	}
	set := strings.Join(sets, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE `id`=%d", table, set, id)
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func RemoveDoc(database string, collection string, id int64) error {
	db := storage.DB()
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", table)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func SetCache(key string, value string) error {
	db := storage.DB()
	_, err := db.Exec("INSERT INTO `database_cache` (`key`, `value`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `value`=?", key, value, value)
	if err != nil {
		return err
	}
	return nil
}

func GetCache(key string) (string, error) {
	db := storage.DB()
	var value string
	err := db.QueryRow("SELECT `value` FROM `database_cache` WHERE `key`=?", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func RemoveCache(key string) error {
	db := storage.DB()
	_, err := db.Exec("DELETE FROM `cache` WHERE `key`=?", key)
	if err != nil {
		return err
	}
	return nil
}
