package routes

import (
	"redemption/controllers"
	"redemption/middlewares/validators"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/user", handlers...)
	{
		user.GET(
			"",
			controllers.GetUser,
		)

		user.PUT(
			"",
			validators.UpdateUserValidator(),
			controllers.UpdateUser,
		)

		user.PUT(
			"measurement",
			validators.UpdateMeasurementValidator(),
			controllers.UpdateMeasurement,
		)
	}
}
