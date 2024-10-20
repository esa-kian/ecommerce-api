package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.ProductService
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	createdProduct, err := pc.ProductService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}
