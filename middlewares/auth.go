package middlewares

import (
	"net/http"
	"redemption/models"
	db "redemption/models/db"
	"redemption/services"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenModel, err := services.VerifyToken(token, db.TokenTypeAccess)
		if err != nil {
			models.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		
		reqMetadata, _ := models.NewRequestMetadata(tokenModel.User)
		c.Set("reqMetadata", reqMetadata)

		c.Next()
	}
}
