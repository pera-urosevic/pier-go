package main

import (
	"os"

	"github.com/pera-urosevic/pier/env"
	"github.com/pera-urosevic/pier/monitor"
	"github.com/pera-urosevic/pier/newsstand"
	"github.com/pera-urosevic/pier/notify"
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
