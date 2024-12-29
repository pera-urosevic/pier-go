package database

import (
	"pier/api/seeker/types"
	"pier/storage"
)

func UpdateTarget(title string, target types.Target) error {
	db := storage.DB()

	_, err := db.Exec("UPDATE `seeker` SET `title`=?, `sources`=?, `release`=?, `checked`=?, `note`=? WHERE `title`=?", target.Title, target.Sources, target.Release, target.Checked, target.Note, title)
	if err != nil {
		return err
	}

	return nil
}
