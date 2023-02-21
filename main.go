package main

import (
	"fmt"
	"localhost/pier/env"
	"localhost/pier/monitor"
	"localhost/pier/newsstand"
	"os"
)

func main() {
	fmt.Println("Pier starting")

	env.Load()
	go monitor.Run()
	go newsstand.Run()

	done := make(chan os.Signal, 1)
	<-done

	fmt.Println("Pier ending")
}
