package database

import (
	"pier/storage"
)

func SetCache(key string, value string) error {
	db := storage.DB()

	_, err := db.Exec("INSERT INTO `database_cache` (`key`, `value`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `value`=?", key, value, value)
	if err != nil {
		return err
	}

	return nil
}
