package main

import (
	"os"

	"somnusalis.org/pier/env"
	"somnusalis.org/pier/monitor"
	"somnusalis.org/pier/newsstand"
	"somnusalis.org/pier/notify"
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
