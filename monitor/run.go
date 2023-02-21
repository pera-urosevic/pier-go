package monitor

import (
	"localhost/pier/monitor/sensor"
	"os"
)

func Run() {
	if os.Getenv("RUN_MONITOR") != "true" {
		return
	}

	go sensor.Uptime()
	go sensor.Storage()
	go sensor.Temp()
	go sensor.Cpu()
	go sensor.Mem()
	go sensor.Swap()
	go sensor.Heartbeat()
}
