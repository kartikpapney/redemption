package routes

import (
	"redemption/controllers"

	"github.com/gin-gonic/gin"
)

func TagRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/tag", handlers...)
	{
		user.GET(
			"",
			controllers.GetAllTag,
		)

		user.POST(
			"",
			controllers.CreateTag,
		)

		user.PUT(
			"",
			controllers.UpdateTag,
		)

		user.DELETE(
			"",
			controllers.DeleteTag,
		)

	}
}
