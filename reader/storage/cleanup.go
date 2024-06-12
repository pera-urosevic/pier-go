package storage

import (
	"pier/database"
	"pier/notify"
	"time"
)

func Cleanup() {
	db := database.Connect()
	threshold := time.Now().Add(-1 * time.Hour * 24 * 28).Format("2006-01-02")
	_, err := db.Exec("DELETE FROM `reader_articles` WHERE `discarded` = 1 AND `id` < ?", threshold)
	if err != nil {
		notify.Warn("reader", "Cleanup: "+err.Error())
	}
}
