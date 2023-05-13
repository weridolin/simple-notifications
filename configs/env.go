package config

import (
	"github.com/joho/godotenv"
)

func init() {
	var logger = GetLogger()
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatalln("Error loading .env file")
	} else {
		logger.Println("loading env from .env success")
	}
}
