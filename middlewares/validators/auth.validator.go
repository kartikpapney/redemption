package validators

import (
	"net/http"
	"redemption/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var registerRequest models.RegisterRequest
		_ = c.ShouldBindBodyWith(&registerRequest, binding.JSON)

		if err := registerRequest.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginRequest models.LoginRequest
		_ = c.ShouldBindBodyWith(&loginRequest, binding.JSON)

		if err := loginRequest.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func RefreshValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var refreshRequest models.RefreshRequest
		_ = c.ShouldBindBodyWith(&refreshRequest, binding.JSON)

		if err := refreshRequest.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func VerifyUserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var verifyRequest models.VerifyRequest
		_ = c.ShouldBindQuery(&verifyRequest)
		
		if err := verifyRequest.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}
