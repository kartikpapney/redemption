package routes

import (
	"net/http"
	"redemption/docs"
	"redemption/middlewares"
	responseModel "redemption/models/response"
	"redemption/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	r := gin.New()
	initRoute(r)

	r.Use(gin.LoggerWithWriter(middlewares.LogWriter()))
	r.Use(gin.CustomRecovery(middlewares.AppRecovery()))
	r.Use(middlewares.CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		PingRoute(v1)
		AuthRoute(v1)
		UserRoute(v1, middlewares.JWTMiddleware())
		// ProductRoute(v1, middlewares.JWTMiddleware())
		ExerciseRoute(v1, middlewares.JWTMiddleware())
		// EquipmentRoute(v1, middlewares.JWTMiddleware())
		// TagRoute(v1, middlewares.JWTMiddleware())
		// WorkoutRoute(v1, middlewares.JWTMiddleware())
	}

	docs.SwaggerInfo.BasePath = v1.BasePath()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func initRoute(r *gin.Engine) {
	_ = r.SetTrustedProxies(nil)
	r.RedirectTrailingSlash = false
	r.HandleMethodNotAllowed = true

	r.NoRoute(func(c *gin.Context) {
		responseModel.SendErrorResponse(c, http.StatusNotFound, c.Request.RequestURI+" not found")
	})

	r.NoMethod(func(c *gin.Context) {
		responseModel.SendErrorResponse(c, http.StatusMethodNotAllowed, c.Request.Method+" is not allowed here")
	})
}

func InitGin() {
	gin.DisableConsoleColor()
	gin.SetMode(services.Config.Mode)
}
