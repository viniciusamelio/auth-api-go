package main

import (
	"auth_api/config"
	"auth_api/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
		panic("failed to load env")
	}

	config.InitializeDatabase()
	router.Init()
}
