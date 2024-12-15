package routes

import (
	"redemption/controllers"
	"redemption/middlewares/validators"

	"github.com/gin-gonic/gin"
)

func ExerciseRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/exercise", handlers...)
	{
		user.GET(
			"",
			controllers.GetAllExercise,
		)

		user.POST(
			"",
			validators.CreateExerciseValidator(),
			controllers.CreateExercise,
		)

		user.GET(
			":id",
			validators.PathIdValidator(),
			controllers.GetExercise,
		)

		user.PUT(
			":id",
			validators.PathIdValidator(),
			controllers.UpdateExercise,
		)

		user.DELETE(
			":id",
			validators.PathIdValidator(),
			controllers.DeleteExercise,
		)

	}
}
