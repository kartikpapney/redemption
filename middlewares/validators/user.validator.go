package validators

import (
	"fmt"
	"net/http"
	requestModel "redemption/models/request"
	responseModel "redemption/models/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetUserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func UpdateUserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUserRequest requestModel.UserUpdateRequest
		err := c.ShouldBindBodyWith(&updateUserRequest, binding.JSON)
		fmt.Println(updateUserRequest.PreferredUnit.Height)
		if err != nil {

			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := updateUserRequest.Validate(); err != nil {

			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("Hi")
		c.Next()
	}
}

func UpdateMeasurementValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUserRequest requestModel.UserMeasurementTrackRequest
		err := c.ShouldBindBodyWith(&updateUserRequest, binding.JSON)

		if err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := updateUserRequest.Validate(); err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("Hi")
		c.Next()
	}
}
