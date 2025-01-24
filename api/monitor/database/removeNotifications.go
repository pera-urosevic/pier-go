package database

import (
	"pier/api/monitor/database/model"
	"pier/storage"
)

func RemoveNotifications(channel string) error {
	db, err := storage.DB()
	if err != nil {
		return err
	}

	res := db.Where("channel = ?", channel).Delete(&model.Notification{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
