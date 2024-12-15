package controllers

import (
	"net/http"
	responseModel "redemption/models/response"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      Ping
// @Description  check server
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseModel.Response
// @Router       /ping [get]
func Ping(c *gin.Context) {
	response := &responseModel.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Working!",
	}

	// timestamp := time.Now().Unix()
	// services.Set(context.Background(), "time", timestamp)
	// value, err := services.Get(c, "time")
	// if err != nil {
	// 	response.StatusCode = http.StatusInternalServerError
	// 	response.Success = false
	// 	response.Message = "Failed to set timestamp in cache"
	// 	response.SendResponse(c)
	// 	return
	// }

	response.SendResponse(c)
}
