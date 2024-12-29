package database

import (
	"pier/api/seeker/types"
	"pier/storage"
)

func GetTargets() ([]types.Target, error) {

	db := storage.DB()

	var targets = []types.Target{}
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
