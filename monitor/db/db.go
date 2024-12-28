package db

import (
	"fmt"
	"pier/storage"
)

func Get(key string) string {
	db := storage.DB()
	row := db.QueryRow("SELECT `value` FROM `monitor` WHERE `key` = ?", key)
	var value string
	err := row.Scan(&value)
	if err == nil {
		return value
	}
	return ""
}

func Set(key string, value interface{}) {
	db := storage.DB()
	valueString := fmt.Sprint(value)
	db.Exec("INSERT INTO `monitor` (`key`, `value`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `value`=?", key, valueString, valueString)
}

func Del(key string) {
	db := storage.DB()
	db.Exec("DELETE FROM `monitor` WHERE `key` like ?", key)
}
