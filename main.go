package main

import (
	"os"

	"pier/env"
	"pier/monitor"
	"pier/newsstand"
	"pier/notify"
)

func main() {
	env.Load()

	notify.Info("pier", "starting")

	go monitor.Run()
	go newsstand.Run()

	done := make(chan os.Signal, 1)
	<-done

	notify.Info("pier", "ending")
}
