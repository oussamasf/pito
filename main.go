package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oussamasf/pito/internal/config"
	db "github.com/oussamasf/pito/pkg/databases"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	env, err := config.Load()
	if err != nil {
		log.Print(err)
		log.Fatal("Failed to validate config file")
	}
	gin.SetMode(gin.ReleaseMode)

	dbConfig := config.GetDBConfig()

	dbErr := db.Initialize(&dbConfig)
	if dbErr != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	r := setupRouter()

	port := fmt.Sprintf("%s:%s", "localhost", env.Port)
	r.Run(port)
}
