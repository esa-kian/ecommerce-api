package config

import (
	"os"
)

type Config struct {
	MongoURI  string
	DBName    string
	JWTSecret string
}

func LoadConfig() Config {
	return Config{
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:    getEnv("DB_NAME", "ecommerce"),
		JWTSecret: getEnv("JWT_SECRET", "your_secret_key"),
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
