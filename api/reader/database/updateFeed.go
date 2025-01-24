package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func UpdateFeed(name string, feed model.Feed) (model.Feed, error) {
	db, err := storage.DB()
	if err != nil {
		return feed, err
	}

	res := db.Where("name = ?", name).Save(&feed)
	if res.Error != nil {
		return feed, res.Error
	}

	return feed, nil
}
