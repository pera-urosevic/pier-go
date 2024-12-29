package database

import (
	"pier/storage"
)

func RemoveCache(key string) error {
	db := storage.DB()

	_, err := db.Exec("DELETE FROM `cache` WHERE `key`=?", key)
	if err != nil {
		return err
	}

	return nil
}
