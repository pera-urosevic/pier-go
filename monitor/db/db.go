package db

import (
	"fmt"
	"pier/api/monitor/database/model"
	"pier/storage"
)

func Get(key string) string {
	db, err := storage.DB()
	if err != nil {
		return ""
	}

	var stat model.Stat
	res := db.Where("`key` = ?", key).Find(&stat)
	if res.Error != nil {
		return ""
	}

	return stat.Value
}

func Set(key string, value interface{}) {
	db, err := storage.DB()
	if err != nil {
		return
	}

	valueString := fmt.Sprint(value)
	db.Exec("INSERT INTO `monitor` (`key`, `value`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `value`=?", key, valueString, valueString)
}

func Del(key string) {
	db, err := storage.DB()
	if err != nil {
		return
	}

	db.Where("`key` like ?", key).Delete(&model.Stat{})
}
