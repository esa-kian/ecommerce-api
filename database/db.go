package database

import (
	"context"
	"ecommerce-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectDB(cfg config.Config) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("failed to connect to database")
	}

	DB = Client.Database(cfg.DBName)
}
