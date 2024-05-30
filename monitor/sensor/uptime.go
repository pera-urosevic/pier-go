package sensor

import (
	"time"

	"github.com/pera-urosevic/pier/database"
	"github.com/pera-urosevic/pier/notify"

	statsHost "github.com/shirou/gopsutil/v3/host"
)

const secondsMinute = 60
const secondsHour = secondsMinute * 60
const secondsDay = secondsHour * 24

func sensorUptime() {
	uptime, err := statsHost.Uptime()
	if err != nil {
		notify.ErrorAlert("monitor", "get uptime", err)
		return
	}
	days := uptime / secondsDay
	hours := (uptime % secondsDay) / secondsHour

	db := database.Connect()
	db.Del(database.Ctx, "monitor:uptime")
	db.HSet(database.Ctx, "monitor:uptime", "days", days, "hours", hours)
}

func Uptime() {
	sensorUptime()
	tickerUptime := time.NewTicker(1 * time.Hour)
	for range tickerUptime.C {
		sensorUptime()
	}
}
