package api

import (
	"os"
	"pier/api/brew"
	"pier/api/monitor"
	"pier/api/seeker"
	"pier/api/subtler"
	"pier/notify"
)

func Run() {
	if os.Getenv("RUN_API") != "true" {
		return
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "55555"
	}
	addr := ":" + port

	notify.Info("api", addr)
	router := customGin()
	router = brew.Routes(router)
	router = monitor.Routes(router)
	router = seeker.Routes(router)
	router = subtler.Routes(router)
	router.Run(addr)
}
