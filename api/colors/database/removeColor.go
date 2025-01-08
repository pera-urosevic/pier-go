package database

import (
	"pier/storage"
)

func RemoveColor(name string) error {
	db := storage.DB()

	_, err := db.Exec("DELETE FROM `colors` WHERE `name` = ?", name)
	if err != nil {
		return err
	}

	return nil
}
