package database

import (
	"pier/api/colors/database/model"
	"pier/storage"
)

func GetColors() ([]model.Color, error) {
	var colors = []model.Color{}

	db, err := storage.DB()
	if err != nil {
		return nil, err
	}

	res := db.Order("name ASC").Find(&colors)
	if res.Error != nil {
		return colors, res.Error
	}

	return colors, nil
}
