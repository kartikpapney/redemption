package routes

import (
	"redemption/controllers"

	"github.com/gin-gonic/gin"
)

func EquipmentRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	user := router.Group("/equipment", handlers...)
	{
		user.GET(
			"",
			controllers.GetAllEquipment,
		)

		user.POST(
			"",
			controllers.CreateEquipment,
		)

		user.PUT(
			"",
			controllers.UpdateEquipment,
		)

		user.DELETE(
			"",
			controllers.DeleteEquipment,
		)

	}
}
