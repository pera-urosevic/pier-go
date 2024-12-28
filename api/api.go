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
	router = brew.Routes(router)
	router = database.Routes(router)
	router = monitor.Routes(router)
	router = proxy.Routes(router)
	router = reader.Routes(router)
	router = seeker.Routes(router)
	router = subtler.Routes(router)
	router = tvshows.Routes(router)
	router.Run(addr)
}
