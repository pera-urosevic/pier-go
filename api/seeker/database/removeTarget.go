package database

import (
	"pier/api/seeker/database/model"
	"pier/storage"
)

func RemoveTarget(title string) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Where("title = ?", title).Delete(&model.Target{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
