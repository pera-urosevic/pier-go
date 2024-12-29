package database

import (
	"errors"
	"pier/storage"
)

func RemoveTarget(title string) error {
	db := storage.DB()

	res, err := db.Exec("DELETE FROM `seeker` WHERE `title`=?", title)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("failed to delete")
	}

	return nil
}
