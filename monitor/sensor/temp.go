package sensor

import (
	"fmt"
	"time"

	"github.com/pera-urosevic/pier/database"
	"github.com/pera-urosevic/pier/monitor/alert"
	"github.com/pera-urosevic/pier/notify"

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
