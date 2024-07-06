package controller

import (
	"golang-BE/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (us *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	return response.ErrorResponse(c,20003,"")
	// }

	// return response.SuccessResponse(c,20001,[]string{"hello"})
}