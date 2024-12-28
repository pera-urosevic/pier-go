package storage

import "pier/storage"

func Reload() bool {
	db := storage.DB()
	row := db.QueryRow("SELECT COUNT(`name`) FROM `reader_feeds` WHERE `disabled` = 0 AND `updated` = 0")
	var value int
	err := row.Scan(&value)
	if err == nil {
		return value > 0
	}
	return false
}
