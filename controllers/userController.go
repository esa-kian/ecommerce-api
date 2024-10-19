package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid input"})
		return
	}

	if err := uc.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
