package services

import (
	"context"
	"ecommerce-api/database"
	"ecommerce-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct{}

func (os *OrderService) CreateOrder(order *models.Order) (*models.Order, error) {
	collection := database.DB.Collection("orders")
	result, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}
	order.ID = result.InsertedID.(primitive.ObjectID) // Assign the ID
	return order, nil
}
