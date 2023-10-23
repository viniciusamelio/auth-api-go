package main

import (
	"auth_api/internal/config"
	"auth_api/internal/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
		panic("failed to load env")
	}
	if os.Getenv("ENV") == "Prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	config.InitializeDatabase()
	router.Init()
}
