package main

import (
	"os"

	"pier/env"
	"pier/monitor"
	"pier/notify"
	"pier/reader"
)

func main() {
	env.Load()

	notify.Info("pier", "starting")

	go monitor.Run()
	go reader.Run()

	done := make(chan os.Signal, 1)
	<-done

	notify.Info("pier", "ending")
}
