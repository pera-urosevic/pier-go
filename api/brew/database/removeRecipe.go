package database

import (
	"pier/storage"
)

func RemoveRecipe(id int64) error {
	db := storage.DB()

	_, err := db.Exec("DELETE FROM `brew` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return nil
}
