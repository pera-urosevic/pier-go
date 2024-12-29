package database

import (
	"pier/storage"
)

func GetFacets(database string, collection string) (string, error) {
	var facets string

	db := storage.DB()

	err := db.
		QueryRow("SELECT `facets` FROM `databases` WHERE `database`=? AND `collection`=?", database, collection).
		Scan(&facets)

	return facets, err
}
