package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func DiscardFeed(name string) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Save(&model.Article{FeedName: name, Discarded: true})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
