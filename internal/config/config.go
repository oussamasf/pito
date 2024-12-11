package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// ? Struct to define environment variables with validation tags
type Config struct {
	Port     string `validate:"required,numeric"`
	Env      string `validate:"required,oneof=development production testing"`
	Database string `validate:"required,url"`
}

// ? Load environment variables, map them to the struct, and validate
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables instead.")
	}

	config := Config{
		Port:     os.Getenv("PORT"),
		Env:      os.Getenv("ENV"),
		Database: os.Getenv("DATABASE_URL"),
	}

	validate := validator.New()

	//? Validate the struct
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return &config, nil
}
