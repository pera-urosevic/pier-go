package database

import (
	"pier/storage"
)

func UpdateFacets(database string, collection string, facets string) error {
	db := storage.DB()

	_, err := db.Exec("UPDATE `databases` SET `facets`=? WHERE `database`=? AND `collection`=?", facets, database, collection)
	if err != nil {
		return err
	}

	return nil
}
