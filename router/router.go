package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	runner := gin.Default()
	runner.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	authRoutes(runner)
	runner.Run()
}
