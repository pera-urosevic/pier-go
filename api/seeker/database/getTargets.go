package database

import (
	"pier/api/seeker/database/model"
	"pier/storage"
)

func GetTargets() ([]model.Target, error) {
	var targets = []model.Target{}

	db, err := storage.DB()
	if err != nil {
		return nil, err
	}

	res := db.Order("`release` asc").Order("`title` asc").Find(&targets)
	if res.Error != nil {
		return nil, res.Error
	}

	return targets, nil
}
