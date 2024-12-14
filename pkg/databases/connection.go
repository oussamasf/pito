package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	models "github.com/oussamasf/pito/internal/models/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *sql.DB
	once sync.Once
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetDB() *sql.DB {
	return db
}

func Initialize(config *Config) error {
	var err error

	once.Do(func() {
		connStr := generateConnectionString(config)
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			return
		}

		err = db.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatalf("Failed to auto-migrate models: %v", err)
		}

		log.Println("Database connected and models migrated!")
	})

	return err
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func generateConnectionString(config *Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)
}
