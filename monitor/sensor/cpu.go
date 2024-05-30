package sensor

import (
	"fmt"
	"time"

	"somnusalis.org/pier/database"
	"somnusalis.org/pier/monitor/alert"
	"somnusalis.org/pier/notify"

	statsCpu "github.com/shirou/gopsutil/v3/cpu"
)

func cpu() {
	usage, err := statsCpu.Percent(0, false)
	if err != nil {
		notify.ErrorAlert("monitor", "get cpu", err)
		return
	}

	alert.Signal("cpu usage", 3, usage[0] > 90.0, fmt.Sprintf("%f", usage[0]))

	db := database.Connect()
	db.Del(database.Ctx, "monitor:cpu")
	db.HSet(database.Ctx, "monitor:cpu", "usage", usage[0])
}

func Cpu() {
	cpu()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		cpu()
	}
}
