package services

import (
	"context"
	"ecommerce-api/database"
	"ecommerce-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct{}

func (ps *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	collection := database.DB.Collection("products")
	result, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}
