package controllers

import (
	"net/http"
	requestModel "redemption/models/request"
	responseModel "redemption/models/response"
	"redemption/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateNewNote godoc
// @Summary      Get Logged In User
// @Description  get logged in user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /user [get]
// @Security     ApiKeyAuth
func GetUser(c *gin.Context) {
	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)
	user, err := services.FindUserById(reqMetadata.UserId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"user": user,
	}
	response.SendResponse(c)
}

// CreateNewNote godoc
// @Summary      Update Logged In User
// @Description  update logged in user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.UserUpdateRequest true "Update User Request"
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /user [put]
// @Security     ApiKeyAuth
func UpdateUser(c *gin.Context) {

	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var requestBody requestModel.UserUpdateRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}
	err := services.UpdateUserById(reqMetadata.UserId, requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}

func UpdateMeasurement(c *gin.Context) {

	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var requestBody requestModel.UserMeasurementTrackRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}
	user, _ := services.FindUserById(reqMetadata.UserId)
	err := services.UpdateMeasurementById(reqMetadata.UserId, user.PreferredUnit, requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}
