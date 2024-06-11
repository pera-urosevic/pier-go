package storage

import (
	"pier/database"
)

func Reload() bool {
	db := database.Connect()
	row := db.QueryRow("SELECT COUNT(`name`) FROM `reader_feeds` WHERE `disabled` = 0 AND `updated` = 0")
	var value int
	err := row.Scan(&value)
	if err == nil {
		return value > 0
	}
	return false
}
