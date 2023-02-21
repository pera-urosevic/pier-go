package sensor

import (
	"fmt"
	"localhost/pier/database"
	"localhost/pier/monitor/alert"
	"time"

	statsMem "github.com/shirou/gopsutil/v3/mem"
)

func mem() {
	vm, err := statsMem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
		return
	}

	alert.Signal("mem usage", 2, vm.UsedPercent > 90.0, fmt.Sprintf("%f", vm.UsedPercent))

	db := database.Connect()
	db.Del(database.Ctx, "monitor:mem")
	db.HSet(database.Ctx, "monitor:mem", "usage", vm.UsedPercent)
}

func Mem() {
	mem()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		mem()
	}
}
