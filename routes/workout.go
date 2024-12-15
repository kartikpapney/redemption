package routes

import (
	"redemption/controllers"

	"github.com/gin-gonic/gin"
)

func WorkoutRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/workout", handlers...)
	{
		user.GET(
			"",
			controllers.GetAllWorkout,
		)

		user.POST(
			"",
			controllers.CreateWorkout,
		)

		user.PUT(
			"",
			controllers.UpdateWorkout,
		)

		user.DELETE(
			"",
			controllers.DeleteWorkout,
		)

	}
}
