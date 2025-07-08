package config

import (
	"context"
	"log"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type ConfigS struct {
    Port  int `env:"PORT" default:"8080"`
	Host  string `env:"HOST" default:"localhost"`
}

var AppConfig ConfigS

func Init() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on reading configuration from .env file", "error", err.Error())
	}

	err = envconfig.Process(ctx, &AppConfig)
	if err != nil {
		log.Fatal("Error on loading environment variables", "error", err.Error())
	}
}
