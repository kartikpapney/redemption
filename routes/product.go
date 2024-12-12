package routes

import (
	"redemption/controllers"
	"redemption/middlewares/validators"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	product := router.Group("/product", handlers...)
	{
		product.POST(
			"",
			controllers.CreateProduct,
		)
		
		product.GET(
			"",
			validators.PaginationRequestValidator(),
			controllers.GetAllProducts,
		)

		product.DELETE(
			"",
			controllers.DeleteAllProduct,
		)

	}
}
