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
// @Summary      Get All Products
// @Description  creates a new note
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.PaginatedRequest true "Get All Products Request"
// @Success      200  {object}  responseModel.PaginatedResponse
// @Failure      400  {object}  responseModel.PaginatedResponse
// @Router       /product [get]
// @Security     ApiKeyAuth
func GetAllProducts(c *gin.Context) {

	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var paginated requestModel.PaginatedRequest
	_ = c.ShouldBindQuery(&paginated)

	response := &responseModel.PaginatedResponse{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	products, total, err := services.GetAllProducts(reqMetadata.UserId, paginated.Limit, (paginated.Page-1)*paginated.Limit)

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
		"products": products,
	}
	response.SendResponse(c)
}

// CreateNewProduct godoc
// @Summary      Create Product
// @Description  creates a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        req  body      requestModel.CreateProductRequest true "Create Product Request"
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /product [post]
// @Security     ApiKeyAuth
func CreateProduct(c *gin.Context) {

	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var requestQuery requestModel.CreateProductRequest
	_ = c.ShouldBindBodyWith(&requestQuery, binding.JSON)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	product, err := services.CreateProduct(requestQuery.Name, requestQuery.Manufacturer, reqMetadata.UserId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"product": product,
	}
	response.SendResponse(c)
}

// DeleteAllProduct godoc
// @Summary      Delete All Products
// @Description  deletes all product of user
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseModel.Response
// @Failure      400  {object}  responseModel.Response
// @Router       /product [delete]
// @Security     ApiKeyAuth
func DeleteAllProduct(c *gin.Context) {

	reqMetadataContext, _ := c.Get("reqMetadata")
	var reqMetadata *requestModel.RequestMetadata = reqMetadataContext.(*requestModel.RequestMetadata)

	var requestQuery requestModel.CreateProductRequest
	_ = c.ShouldBindBodyWith(&requestQuery, binding.JSON)

	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	err := services.DeleteAllProducts(reqMetadata.UserId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.Message = "all products deleted"
	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}
