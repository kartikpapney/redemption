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
// @Summary      Get All Exercise
// @Description  get all exercise
// @Tags         exercise
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.PaginatedRequest true "Get All Exercise Request"
// @Success      200  {object}  responseModel.PaginatedResponse
// @Failure      400  {object}  responseModel.PaginatedResponse
// @Router       /exercise [get]
// @Security     ApiKeyAuth
func GetAllExercise(c *gin.Context) {
	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var paginated requestModel.PaginatedRequest
	_ = c.ShouldBindQuery(&paginated)

	response := &responseModel.PaginatedResponse{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	exercise, total, err := services.GetAllExercise(reqMetadata.UserId, paginated.Limit, (paginated.Page-1)*paginated.Limit)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.Limit = paginated.Limit
	response.Page = paginated.Page
	response.Total = total
	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"exercise": exercise,
	}
	response.SendResponse(c)
}

// CreateNewNote godoc
// @Summary      Get Exercise
// @Description  get exercise
// @Tags         exercise
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Exercise ID"
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /exercise/{id} [get]
// @Security     ApiKeyAuth
func GetExercise(c *gin.Context) {
	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	reqContext, _ := c.Get("id")
	exerciseId := reqContext.(*requestModel.PathIdRequest)

	var paginated requestModel.PaginatedRequest
	_ = c.ShouldBindQuery(&paginated)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	exercise, err := services.GetExercise(reqMetadata.UserId, exerciseId.Id)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"exercise": exercise,
	}
	response.SendResponse(c)
}

// CreateNewNote godoc
// @Summary      Create Exercise
// @Description  create exercise
// @Tags         exercise
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.CreateExerciseRequest true "Create Exercise Request"
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /exercise [post]
// @Security     ApiKeyAuth
func CreateExercise(c *gin.Context) {
	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var requestQuery requestModel.CreateExerciseRequest
	_ = c.ShouldBindBodyWith(&requestQuery, binding.JSON)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	exercise, err := services.CreateExercise(reqMetadata.UserId, requestQuery.Name, requestQuery.Primary, requestQuery.Secondary, requestQuery.Other)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"exercise": exercise,
	}
	response.SendResponse(c)
}

// CreateNewNote godoc
// @Summary      Update Exercise
// @Description  update exercise
// @Tags         exercise
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Exercise ID"
// @Param        req  body      requestModel.UpdateExerciseRequest true "Update Exercise Request"
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /exercise/{id} [put]
// @Security     ApiKeyAuth
func UpdateExercise(c *gin.Context) {
	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	var requestQuery requestModel.UpdateExerciseRequest
	_ = c.ShouldBindBodyWith(&requestQuery, binding.JSON)

	reqPath := &requestModel.PathIdRequest{}
	if err := c.BindUri(reqPath); err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	exercise, err := services.UpdateExerciseById(reqMetadata.UserId, reqPath.Id, requestQuery.Name, requestQuery.Primary, requestQuery.Secondary, requestQuery.Other)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"exercise": exercise,
	}
	response.SendResponse(c)
}

// DeleteExercise godoc
// @Summary      Delete Exercise Request
// @Description  Delete exercise request
// @Tags         exercise
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Exercise ID"
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /exercise/{id} [delete]
// @Security     ApiKeyAuth
func DeleteExercise(c *gin.Context) {
	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	reqPath := &requestModel.PathIdRequest{}
	if err := c.BindUri(reqPath); err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	err := services.DeleteExerciseById(reqMetadata.UserId, reqPath.Id)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.Message = "exercise deleted"
	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}
