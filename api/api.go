package api

import (
	"os"
	"pier/api/brew"
	"pier/api/database"
	"pier/api/monitor"
	"pier/api/proxy"
	"pier/api/reader"
	"pier/api/seeker"
	"pier/api/subtler"
	"pier/api/tvshows"
	"pier/lib"
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

	router := lib.CustomGin()

	brew.Routes(router)
	database.Routes(router)
	monitor.Routes(router)
	proxy.Routes(router)
	reader.Routes(router)
	seeker.Routes(router)
	subtler.Routes(router)
	tvshows.Routes(router)

	router.Run(addr)
}
