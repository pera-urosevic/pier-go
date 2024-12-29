package database

import (
	"fmt"
	"pier/storage"
)

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
