package middlewares

import (
	"net/http"
	db "redemption/models/db"
	requestModel "redemption/models/request"
	responseModel "redemption/models/response"
	"redemption/services"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenModel, err := services.VerifyToken(token, db.TokenTypeAccess)
		if err != nil {
			responseModel.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		reqMetadata, _ := requestModel.NewRequestMetadata(tokenModel.User)
		c.Set("reqMetadata", reqMetadata)

		c.Next()
	}
}
