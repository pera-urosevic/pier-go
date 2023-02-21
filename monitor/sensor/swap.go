package sensor

import (
	"fmt"
	"localhost/pier/database"
	"time"

	statsMem "github.com/shirou/gopsutil/v3/mem"
)

func swap() {
	swap, err := statsMem.SwapMemory()
	if err != nil {
		fmt.Println(err)
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
