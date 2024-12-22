package database

import (
	"errors"
	"pier/api/seeker/types"
	"pier/database"

	_ "modernc.org/sqlite"
)

func GetTargets() ([]types.Target, error) {
	var targets = []types.Target{}
	db := database.Connect()
	rows, err := db.Query("SELECT * FROM `seeker` ORDER BY `release` ASC, `title` ASC")
	if err != nil {
		return targets, err
	}
	for rows.Next() {
		var target types.Target
		err := rows.Scan(&target.Title, &target.Sources, &target.Release, &target.Checked, &target.Note)
		if err != nil {
			return targets, err
		}
		targets = append(targets, target)
	}
	return targets, nil
}

func CreateTarget(target types.Target) error {
	db := database.Connect()
	_, err := db.Exec("INSERT INTO `seeker` (`title`, `sources`, `release`, `checked`,  `note`) VALUES(?, ?, ?, ?, ?)", target.Title, target.Sources, target.Release, target.Checked, target.Note)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTarget(title string, target types.Target) error {
	db := database.Connect()
	_, err := db.Exec("UPDATE `seeker` SET `title`=?, `sources`=?, `release`=?, `checked`=?, `note`=? WHERE `title`=?", target.Title, target.Sources, target.Release, target.Checked, target.Note, title)
	if err != nil {
		return err
	}
	return nil
}

func RemoveTarget(title string) error {
	db := database.Connect()
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
