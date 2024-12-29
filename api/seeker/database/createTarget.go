package database

import (
	"pier/api/seeker/types"
	"pier/storage"
)

func CreateTarget(target types.Target) error {
	db := storage.DB()

	_, err := db.Exec("INSERT INTO `seeker` (`title`, `sources`, `release`, `checked`,  `note`) VALUES(?, ?, ?, ?, ?)", target.Title, target.Sources, target.Release, target.Checked, target.Note)
	if err != nil {
		return err
	}

	return nil
}
