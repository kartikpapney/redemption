package controllers

import (
	"net/http"
	"redemption/models"
	db "redemption/models/db"
	"redemption/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary      Register
// @Description  registers a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      models.RegisterRequest true "Register Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/register [post]
func Register(c *gin.Context) {
	var requestBody models.RegisterRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	err := services.CheckUserMail(requestBody.Email)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	requestBody.Name = strings.TrimSpace(requestBody.Name)
	user, err := services.CreateUser(requestBody.Name, requestBody.Email, requestBody.Password, requestBody.Location)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	emailToken, err := services.GenerateVerifyUserToken(user)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.Message = "verification link sent to your email"
	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{
		"user": user,
		"token": gin.H{
			"email": emailToken.GetResponseJson(),
		},
	}
	response.SendResponse(c)
}



// Login godoc
// @Summary      Login
// @Description  login a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      models.LoginRequest true "Login Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var requestBody models.LoginRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	user, err := services.FindUserByEmail(requestBody.Email)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		response.Message = "email and password don't match"
		response.SendResponse(c)
		return
	}

	accessToken, refreshToken, err := services.GenerateAccessTokens(user)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}
	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"user": user,
		"token": gin.H{
			"access":  accessToken.GetResponseJson(),
			"refresh": refreshToken.GetResponseJson(),
		},
	}
	response.SendResponse(c)
}

// Refresh godoc
// @Summary      Refresh
// @Description  refreshes a user token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      models.RefreshRequest true "Refresh Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/refresh [post]
func Refresh(c *gin.Context) {
	var requestBody models.RefreshRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// check token validity
	token, err := services.VerifyToken(requestBody.Token, db.TokenTypeRefresh)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	user, err := services.FindUserById(token.User)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	// delete old token
	err = services.DeleteTokenById(token.ID)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	accessToken, refreshToken, err := services.GenerateAccessTokens(user)
	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"user": user,
		"token": gin.H{
			"access":  accessToken.GetResponseJson(),
			"refresh": refreshToken.GetResponseJson()},
	}
	response.SendResponse(c)
}



// Register godoc
// @Summary      Verify
// @Description  verifies user email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  body      models.VerifyRequest true "Verify Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /auth/register [post]
func VerifyUser(c *gin.Context) {
	var requestQuery models.VerifyRequest
	_ = c.ShouldBindQuery(&requestQuery)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	token, err := services.VerifyToken(requestQuery.Token, db.TokenTypeEmail)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}
	
	user, err := services.VerifyUserEmail(token.User);
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	accessToken, refreshToken, err := services.GenerateAccessTokens(user)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.Message = "user verified"
	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{
		"user": user,
		"token": gin.H{
			"access":  accessToken.GetResponseJson(),
			"refresh": refreshToken.GetResponseJson(),
		},
	}
	response.SendResponse(c)
}

