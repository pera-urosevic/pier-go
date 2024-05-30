package sensor

import (
	"fmt"
	"time"

	"somnusalis.org/pier/database"
	"somnusalis.org/pier/monitor/alert"
	"somnusalis.org/pier/notify"

	statsHost "github.com/shirou/gopsutil/v3/host"
)

func temp() {
	temps, err := statsHost.SensorsTemperatures()
	if err != nil {
		notify.ErrorAlert("monitor", "get temperature", err)
		return
	}

	db := database.Connect()
	db.Del(database.Ctx, "monitor:temp")
	for _, temp := range temps {
		alert.Signal("temp value", 1, temp.Temperature > 70.0, fmt.Sprintf("%s = %f", temp.SensorKey, temp.Temperature))
		db.HSet(database.Ctx, "monitor:temp", temp.SensorKey, temp.Temperature)
	}
}

func Temp() {
	temp()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		temp()
	}
}
