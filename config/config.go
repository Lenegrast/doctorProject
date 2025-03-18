package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Port string
}

func ReadConfig() Config {
	err := godotenv.Load(".env") //получение токена
	if err != nil {
		logrus.Fatalf("Failed to read .env file. ERROR: %s", err)
	}
	return Config{
		Port: os.Getenv("PORT"),
	}
}
