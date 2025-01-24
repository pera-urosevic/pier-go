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

	res := db.Model(&model.Article{}).Where("id = ?", id).Update("discarded", "1")
	if res.Error != nil {
		return res.Error
	}

	return nil
}
