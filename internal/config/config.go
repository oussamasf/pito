package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	db "github.com/oussamasf/pito/pkg/databases"
)

// ? Struct to define environment variables with validation tags
type Config struct {
	Port     string `validate:"required,numeric"`
	Env      string `validate:"required,oneof=development production testing"`
	Database struct {
		Host     string `validate:"required"`
		Port     string `validate:"required,numeric"`
		Username string `validate:"required"`
		Password string `validate:"required"`
		DBName   string `validate:"required"`
	} `validate:"required"`
}

func (c *Config) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("config validation error: %w", err)
	}

	return nil
}

// ? Load environment variables, map them to the struct, and validate
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables instead.")
	}

	config := Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
		Database: struct {
			Host     string `validate:"required"`
			Port     string `validate:"required,numeric"`
			Username string `validate:"required"`
			Password string `validate:"required"`
			DBName   string `validate:"required"`
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetDBConfig() db.Config {
	return db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}
}
