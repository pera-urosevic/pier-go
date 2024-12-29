package database

import (
	"fmt"
	"pier/storage"
)

func GetDocsCount(database string, collection string, where string) (int, error) {
	db := storage.DB()

	var count int
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s", table, where)
	err := db.QueryRow(query).Scan(&count)

	return count, err
}
