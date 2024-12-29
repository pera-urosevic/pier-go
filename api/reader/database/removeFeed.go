package database

import (
	"pier/storage"
)

func RemoveFeed(name string) error {
	db := storage.DB()

	_, err := db.Exec("DELETE FROM `reader_feeds` WHERE `name` = ?", name)
	if err != nil {
		return err
	}

	return nil
}
