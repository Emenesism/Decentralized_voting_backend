package main

import (
	"github.com/charmbracelet/log"
	"github.com/emenesism/Decentralized-voting-backend/config"
	"github.com/emenesism/Decentralized-voting-backend/router/http"
	"github.com/emenesism/Decentralized-voting-backend/service"
)

func main() {
	config.Init()
	log.Info("Config loaded")

	service.InitContractService()

	http.Init()
}
