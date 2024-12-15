package controllers

import (
	"github.com/gin-gonic/gin"
)

// CreateNewNote godoc
// @Summary      Get All Tag
// @Description  get all tag
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.PaginatedRequest true "Get All Tag Request"
// @Success      200  {object}  responseModel.PaginatedResponse
// @Failure      400  {object}  responseModel.PaginatedResponse
// @Router       /tag [get]
// @Security     ApiKeyAuth
func GetAllTag(c *gin.Context) {

}


// CreateNewNote godoc
// @Summary      Create A Tag
// @Description  create a tag
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.PaginatedRequest true "Create Tag Request"
// @Success      200  {object}  responseModel.PaginatedResponse
// @Failure      400  {object}  responseModel.PaginatedResponse
// @Router       /tag [post]
// @Security     ApiKeyAuth
func CreateTag(c *gin.Context) {

}

// CreateNewNote godoc
// @Summary      Get All Equipment
// @Description  get all equipment
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.PaginatedRequest true "Get All Equipment Request"
// @Success      200  {object}  responseModel.PaginatedResponse
// @Failure      400  {object}  responseModel.PaginatedResponse
// @Router       /tag [put]
// @Security     ApiKeyAuth
func UpdateTag(c *gin.Context) {

}

// CreateNewNote godoc
// @Summary      Get All Equipment
// @Description  get all equipment
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.PaginatedRequest true "Get All Equipment Request"
// @Success      200  {object}  responseModel.PaginatedResponse
// @Failure      400  {object}  responseModel.PaginatedResponse
// @Router       /tag [delete]
// @Security     ApiKeyAuth
func DeleteTag(c *gin.Context) {

}
