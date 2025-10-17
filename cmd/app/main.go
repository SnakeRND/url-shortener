package main

import (
	"fmt"
	"log"
	"url-shortener/internal/config"

	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	cfg := config.MustLoad()

	fmt.Println(cfg)
}
