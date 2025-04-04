package controllers

import (
	"PrimeGin/internal/services"
	"PrimeGin/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.UserService.GetUsers()
	utils.Response(c, users, err)
}
