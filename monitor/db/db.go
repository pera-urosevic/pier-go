package db

import (
	"fmt"
	"pier/database"
)

func Get(key string) string {
	db := database.Connect()
	row := db.QueryRow("SELECT `value` FROM `monitor` WHERE `key` = ?", key)
	var value string
	err := row.Scan(&value)
	if err == nil {
		return value
	}
	return ""
}

func Set(key string, value interface{}) {
	db := database.Connect()
	valueString := fmt.Sprint(value)
	db.Exec("INSERT INTO `monitor` (`key`, `value`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `value`=?", key, valueString, valueString)
}

func Del(key string) {
	db := database.Connect()
	db.Exec("DELETE FROM `monitor` WHERE `key` like ?", key)
}
