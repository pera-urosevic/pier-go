package sensor

import (
	"time"

	"pier/database"
)

func heartbeat() {
	timestamp := time.Now().UnixNano() / 1000

	db := database.Connect()
	db.Del(database.Ctx, "monitor:heartbeat")
	db.HSet(database.Ctx, "monitor:heartbeat", "timestamp", timestamp)
}

func Heartbeat() {
	heartbeat()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		heartbeat()
	}
}
