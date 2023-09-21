package router

import (
	"auth_api/config"
	"auth_api/core"
	"auth_api/core/application"
	"auth_api/core/domain"

	"github.com/gin-gonic/gin"
)

func authRoutes(runner *gin.Engine) {
	const prefix = "/auth"

	runner.GET(prefix+"/login", func(context *gin.Context) {
		var credentials core.CredentialsDto
		context.BindHeader(&credentials)
		var authRepository application.DefaultAuthRepository

		authRepository = application.DefaultAuthRepository{
			Database: config.Database,
		}

		// TODO: Refactor this to not use a repository directly inside the controller
		user, error := authRepository.SignIn(domain.Credentials{
			Username: credentials.Username,
			Password: credentials.Password,
		})

		if error != nil {
			context.JSON(400, gin.H{
				"error":   error,
				"success": false,
			})
			return
		}

		context.JSON(200, gin.H{
			"user":    user,
			"success": true,
		})
	})

	runner.POST(prefix+"/signup", func(context *gin.Context) {

	})
}
