package router

import (
	"github.com/gin-gonic/gin"
)

func authRoutes(runner *gin.Engine) {
	const prefix = "/auth"

	runner.GET(prefix+"/login", func(context *gin.Context) {
		// credentials := core.Credentials{
		// 	Username: context.GetHeader("username"),
		// 	Password: context.GetHeader("password"),
		// }
		context.JSON(200, gin.H{
			// "credentials": credentials,
			"success": true,
		})
	})
}
