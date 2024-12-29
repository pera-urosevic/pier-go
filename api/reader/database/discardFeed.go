package database

import (
	"pier/storage"
)

func DiscardFeed(name string) error {
	db := storage.DB()

	_, err := db.Exec("UPDATE `reader_articles` SET `discarded` = 1 WHERE `feed_name` = ?", name)
	if err != nil {
		return err
	}

	return nil
}
