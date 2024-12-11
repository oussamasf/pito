package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
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

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}

// Initialize initializes the database connection
func Initialize(config Config) error {
	var err error

	once.Do(func() {
		connStr := generateConnectionString(config)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return
		}

		// Set connection pool settings
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

		log.Println("Database connection established")
	})

	return err
}

// Close closes the database connection
func Close() {
	if db != nil {
		db.Close()
	}
}

func generateConnectionString(config Config) string {
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
