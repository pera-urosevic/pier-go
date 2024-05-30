package sensor

import (
	"time"

	"pier/database"
	"pier/notify"

	statsMem "github.com/shirou/gopsutil/v3/mem"
)

func swap() {
	swap, err := statsMem.SwapMemory()
	if err != nil {
		notify.ErrorAlert("monitor", "get swap", err)
		return
	}

	db := database.Connect()
	db.Del(database.Ctx, "monitor:swap")
	db.HSet(database.Ctx, "monitor:swap", "usage", swap.UsedPercent)
}

func Swap() {
	swap()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		swap()
	}
}
