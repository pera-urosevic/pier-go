package database

import (
	"pier/storage"
)

func RemoveTVShow(id int64) error {
	db := storage.DB()

	_, err := db.Exec("DELETE FROM `tvshows` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return nil
}
