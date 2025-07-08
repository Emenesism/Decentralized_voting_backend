package main

import (
	"github.com/emenesism/Decentralized-voting-backend/config"
	"github.com/emenesism/Decentralized-voting-backend/router/http"

	"github.com/charmbracelet/log"
)


func main () {
	config.Init()
	log.Info("Config loaded")
	http.Init()
	log.Info("Server has been running")
}