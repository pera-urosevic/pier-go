package database

import (
	"pier/api/tvshows/database/model"
	"pier/storage"
)

func GetTVShows() ([]model.TVShow, error) {
	var tvshows []model.TVShow

	db, err := storage.DB()
	if err != nil {
		return nil, err
	}

	res := db.Find(&tvshows)
	if res.Error != nil {
		return nil, res.Error
	}

	return tvshows, nil
}
