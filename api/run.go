package api

import (
	"os"
	"pier/api/brew"
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
	router.Run(addr)
}
