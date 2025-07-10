package main

import (
	"github.com/charmbracelet/log"
	"github.com/emenesism/Decentralized-voting-backend/config"
	"github.com/emenesism/Decentralized-voting-backend/router/http"
	"github.com/emenesism/Decentralized-voting-backend/service"
	"github.com/emenesism/Decentralized-voting-backend/models"
)

func main() {
	config.Init()
	log.Info("Config loaded")
	models.Init()
	log.Info("Model init successfuly")
	service.InitContractService()
	http.Init()
}

