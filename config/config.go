package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// load config
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file : ", err)
	}
}
