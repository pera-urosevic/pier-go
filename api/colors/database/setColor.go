package database

import (
	"pier/api/colors/database/model"
	"pier/storage"
)

func SetColor(color model.Color) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Create(&color)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
