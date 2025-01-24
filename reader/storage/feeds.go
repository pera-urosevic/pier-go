package storage

import (
	"pier/api/reader/database/model"
	"pier/notify"
	"pier/storage"
)

func Feeds() []*model.Feed {
	feeds := []*model.Feed{}

	db, err := storage.DB()
	if err != nil {
		notify.ErrorAlert("reader", "get feeds", err)
		return feeds
	}

	res := db.Find(&feeds)
	if res.Error != nil {
		notify.ErrorAlert("reader", "get feeds", res.Error)
		return feeds
	}

	return feeds
}
