package database

import (
	"pier/storage"
)

func DiscardArticle(id string) error {
	db := storage.DB()

	_, err := db.Exec("UPDATE `reader_articles` SET `discarded` = 1 WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return nil
}
