package controllers

import (
	"net/http"
	"redemption/models"
	"redemption/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewNote godoc
// @Summary      Get Logged In User
// @Description  get logged in user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /user [get]
// @Security     ApiKeyAuth
func GetUser(c *gin.Context) {
	
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	userId, exists := c.Get("userId")
	if !exists {
		response.Message = "cannot get user"
		response.SendResponse(c)
		return
	}

	user, err := services.FindUserById(userId.(primitive.ObjectID))
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
