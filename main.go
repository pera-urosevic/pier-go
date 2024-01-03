package main

import (
	"localhost/pier/env"
	"localhost/pier/monitor"
	"localhost/pier/newsstand"
	"localhost/pier/notify"
	"os"
)

func main() {
	env.Load()

	notify.Info("pier", "Starting")

	go monitor.Run()
	go newsstand.Run()

	done := make(chan os.Signal, 1)
	<-done

	notify.Info("pier", "Ending")
}
