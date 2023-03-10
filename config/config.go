package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// API port
	ApiPort string

	// Database user
	DbUser string
	// Database password
	DbPassword string
	// Database host
	DbHost string
	// Database name
	DbName string

	// JWT Secret key
	SecretKey []byte
)

func EnvironmentConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	ApiPort = os.Getenv("PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbName = os.Getenv("DB_NAME")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	return nil
}
