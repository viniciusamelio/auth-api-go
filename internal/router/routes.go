package router

import (
	"auth_api/internal/core"
	dependencies "auth_api/internal/di"

	"github.com/gin-gonic/gin"
)

func authRoutes(runner *gin.Engine) {
	const prefix = "/auth"

	runner.GET(prefix+"/login", func(context *gin.Context) {
		var credentials core.CredentialsDto
		context.BindHeader(&credentials)
		authService := dependencies.NewAuthenticationService()

		user, error := authService.SignIn(credentials)

		if error != nil {
			context.JSON(400, gin.H{
				"error":   error.Error(),
				"success": false,
			})
			return
		}

		context.JSON(200, gin.H{
			"session": user,
			"success": true,
		})
	})

	runner.POST(prefix+"/signup", func(context *gin.Context) {
		var data core.SignUpDto
		context.BindJSON(&data)
		authService := dependencies.NewAuthenticationService()
		user, error := authService.SignUp(data)

		if error != nil {
			context.JSON(400, gin.H{
				"error":   error.Error(),
				"success": false,
			})
			return
		}

		context.JSON(200, gin.H{
			"user":    user,
			"success": true,
		})

	})
}

func sessionRoutes(runner *gin.Engine) {
	const prefix = "/session"

	runner.GET(prefix+"/:id", func(context *gin.Context) {
		sessionService := dependencies.NewSessionService()
		session, error := sessionService.GetSession(context.Param("id"))

		if error != nil {
			context.JSON(400, gin.H{
				"error":   error.Error(),
				"success": false,
			})
			return
		}

		context.JSON(200, gin.H{
			"session": session,
			"success": true,
		})
	})
}
