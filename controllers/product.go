package controllers

import (
	"net/http"
	"redemption/models"
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
// @Param        req  body      models.PaginatedRequest true "Get All Products Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /product [get]
// @Security     ApiKeyAuth
func GetAllProducts(c *gin.Context) {

	reqMetadataContext, _ := c.Get("reqMetadata");
	var reqMetadata *models.RequestMetadata = reqMetadataContext.(*models.RequestMetadata)

	var paginated models.PaginatedRequest
	_ = c.ShouldBindQuery(&paginated)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	products, err := services.GetAllProducts(reqMetadata.UserId, paginated.Limit, (paginated.Page-1)*paginated.Limit)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

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
// @Param        req  body      models.CreateProductRequest true "Create Product Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /product [post]
// @Security     ApiKeyAuth
func CreateProduct(c *gin.Context) {
	
	reqMetadataContext, _ := c.Get("reqMetadata");
	var reqMetadata *models.RequestMetadata = reqMetadataContext.(*models.RequestMetadata)

	var requestQuery models.CreateProductRequest
	_ = c.ShouldBindBodyWith(&requestQuery, binding.JSON)

	response := &models.Response{
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
