package database

import (
	"pier/api/colors/database/model"
	"pier/storage"
)

func RemoveColor(name string) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Where("name = ?", name).Delete(&model.Color{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
