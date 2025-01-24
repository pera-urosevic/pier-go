package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func GetFeed(name string) (model.Feed, error) {
	var feed = model.Feed{}

	db, err := storage.DB()
	if err != nil {
		return feed, err
	}

	res := db.Where("name = ?", name).First(&feed)
	if res.Error != nil {
		return feed, res.Error
	}

	return feed, nil
}
