package services

import (
	"context"
	"ecommerce-api/database"
	"ecommerce-api/models"
)

type UserService struct{}

func (us *UserService) CreateUser(user *models.User) error {
	collection := database.DB.Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}
