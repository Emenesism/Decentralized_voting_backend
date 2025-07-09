package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/emenesism/Decentralized-voting-backend/config"
	"github.com/emenesism/Decentralized-voting-backend/router/http"
	"github.com/emenesism/Decentralized-voting-backend/service"
)

func main() {
	config.Init()
	log.Info("Config loaded")

	service.InitContractService()

	count, err := service.GetVotes("Alice")
	fmt.Println(count, err)
	if err != nil {
		log.Fatal("Failed to read votes:", err)
	}
	log.Info("Votes for Alice:", count)

	http.Init()
	log.Info("Server has been running")
}
