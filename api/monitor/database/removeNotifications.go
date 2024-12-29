package database

import (
	"pier/storage"
)

func RemoveNotifications(channel string) error {
	db := storage.DB()
	_, err := db.Exec("DELETE FROM `notify` WHERE `channel` = ?", channel)
	if err != nil {
		return err
	}
	return nil
}
