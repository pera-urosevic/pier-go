package database

import (
	"pier/api/monitor/database/model"
	"pier/storage"
)

func RemoveNotification(id int64) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Delete(&model.Notification{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
