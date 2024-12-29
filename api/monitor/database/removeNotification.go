package database

import (
	"pier/storage"
)

func RemoveNotification(id int64) error {
	db := storage.DB()

	_, err := db.Exec("DELETE FROM `notify` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return nil
}
