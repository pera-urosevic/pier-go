package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func DiscardArticle(id string) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Save(&model.Article{ID: id, Discarded: true})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
