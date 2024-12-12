package routes

import (
	"redemption/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/user", handlers...)
	{
		user.GET(
			"",
			controllers.GetUser,
		)
	}
}
