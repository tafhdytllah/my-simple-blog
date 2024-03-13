package handler

import (
	"my-simple-blog/dto"
	"my-simple-blog/errorhandler"
	"my-simple-blog/helper"
	"my-simple-blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

// Method
// Register
func (h *authHandler) Register(c *gin.Context) {

	var register dto.RegisterRequest

	// binding request body json to type
	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	// register service
	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	// response body
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register success, please login.",
	})

	c.JSON(http.StatusCreated, res)

}

// Login
func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	// binding request body json to type
	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	// login service
	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	// response body
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login Succesfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)

}
