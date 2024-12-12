package routes

import (
	"redemption/controllers"
	"redemption/middlewares/validators"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST(
			"/register",
			validators.RegisterValidator(),
			controllers.Register,
		)

		auth.GET(
			"/verify",
			validators.VerifyUserValidator(),
			controllers.VerifyUser,
		)

		auth.POST(
			"/login",
			validators.LoginValidator(),
			controllers.Login,
		)

		auth.POST(
			"/refresh",
			validators.RefreshValidator(),
			controllers.Refresh,
		)
	}
}
