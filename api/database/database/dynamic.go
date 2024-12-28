package database

import (
	"database/sql"
	"errors"
	"fmt"
	"pier/api/database/types"
)

type ColumnInfo struct {
	Name string
	Type string
}

func ParseValue(val *interface{}, col ColumnInfo) (interface{}, error) {
	var value interface{}
	switch col.Type {
	case "VARCHAR":
		value = fmt.Sprintf("%s", *val)
	case "TEXT":
		value = fmt.Sprintf("%s", *val)
	case "DATE":
		value = *val
	case "DATETIME":
		value = *val
	case "TINYINT":
		value = *val
	case "BIGINT":
		value = *val
	case "INT":
		value = *val
	case "BLOB":
		value = *val
	default:
		return nil, errors.New("parse value unknown type: column [" + col.Name + "], type " + col.Type)
	}
	return value, nil
}

func MapDocs(rows *sql.Rows) ([]map[string]interface{}, error) {
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	colInfos := make([]ColumnInfo, len(colTypes))
	for i, colType := range colTypes {
		colInfos[i] = ColumnInfo{Name: colType.Name(), Type: colType.DatabaseTypeName()}
	}
	cols := make([]interface{}, len(colInfos))
	colPtrs := make([]interface{}, len(colInfos))
	for i := range cols {
		colPtrs[i] = &cols[i]
	}
	docs := []types.Doc{}
	for rows.Next() {
		if err := rows.Scan(colPtrs...); err != nil {
			return nil, err
		}
		doc := make(map[string]interface{})
		for i := range cols {
			val := colPtrs[i].(*interface{})
			value, err := ParseValue(val, colInfos[i])
			if err != nil {
				return nil, err
			}
			doc[colInfos[i].Name] = value
		}
		docs = append(docs, doc)
	}
	return docs, nil
}
