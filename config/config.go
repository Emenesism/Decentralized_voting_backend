package config 

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	Port string
}

var AppConfig Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using default values")
	}

	config := &Config{
		Port: getEnv("PORT", "8080"),
	}


	AppConfig = config
	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value 
	}

	return defaultValue
}
