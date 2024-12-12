package validators

import (
	"github.com/gin-gonic/gin"
)

func GetUserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
