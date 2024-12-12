package controllers

import (
	"net/http"
	"redemption/models"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      Ping
// @Description  check server
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response
// @Router       /ping [get]
func Ping(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Working!",
	}

	response.SendResponse(c)
}
