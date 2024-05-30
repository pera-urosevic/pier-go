package sensor

import (
	"fmt"
	"time"

	"somnusalis.org/pier/database"
	"somnusalis.org/pier/monitor/alert"
	"somnusalis.org/pier/notify"

	statsDisk "github.com/shirou/gopsutil/v3/disk"
)

func storage() {
	partitions, err := statsDisk.Partitions(true)
	if err != nil {
		notify.ErrorAlert("monitor", "get partitions", err)
		return
	}

	db := database.Connect()
	db.Del(database.Ctx, "monitor:storage")
	for _, partition := range partitions {
		usage, err := statsDisk.Usage(partition.Mountpoint)
		if err != nil {
			notify.ErrorAlert("monitor", "get partition usage", err)
			continue
		}
		if usage.UsedPercent < 1 {
			continue
		}
		alert.Signal("storage usage", 1, usage.UsedPercent > 90.0, fmt.Sprintf("%s - %f", partition.Mountpoint, usage.UsedPercent))
		db.HSet(database.Ctx, "monitor:storage", partition.Mountpoint, usage.UsedPercent)
	}
}

func Storage() {
	storage()
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		storage()
	}
}
