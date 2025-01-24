package storage

import "pier/storage"

func Reload() bool {
	db, err := storage.DB()
	if err != nil {
		return false
	}

	var count int64
	res := db.Where("disabled = 0 AND updated = 0").Count(&count)
	if res.Error != nil {
		return false
	}

	return count > 0
}
