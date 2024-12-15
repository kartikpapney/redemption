package validators

import (
	"net/http"
	requestModel "redemption/models/request"
	responseModel "redemption/models/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateExerciseValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createExerciseRequest requestModel.CreateExerciseRequest
		_ = c.ShouldBindBodyWith(&createExerciseRequest, binding.JSON)

		if err := createExerciseRequest.Validate(); err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func UpdateExerciseValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateExerciseRequest requestModel.UpdateExerciseRequest
		_ = c.ShouldBindBodyWith(&updateExerciseRequest, binding.JSON)

		if err := updateExerciseRequest.Validate(); err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}
