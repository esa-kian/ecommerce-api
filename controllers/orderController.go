package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService services.OrderService
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	createdOrder, err := oc.OrderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}
