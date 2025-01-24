package database

import (
	"pier/api/tvshows/database/model"
	"pier/storage"
)

func CreateTVShow(record model.TVShow) (model.TVShow, error) {
	db, err := storage.DB()
	if err != nil {
		return record, err
	}

	res := db.Create(&record)
	if res.Error != nil {
		return record, res.Error
	}

	return record, nil
}
