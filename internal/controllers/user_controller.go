package controllers

import (
	"PrimeGin/internal/models"
	"PrimeGin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	result := utils.DB.Preload("Roles").Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
