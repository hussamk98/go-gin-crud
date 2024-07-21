package main

import (
	"crud-app/internal/app/crud"
	"crud-app/pkg/config"
	"crud-app/pkg/db"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	port := flag.String("port", "8080", "specify server port")
	flag.Parse()

	// Initialize database connection
	database, err := db.Init()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize Gin router based on the environment
	isDevelopment := config.GetEnv("ENV") == "development"

	var router *gin.Engine
	if isDevelopment {
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	// Setup routes
	crud.SetupRoutes(router, database)

	// Start the server
	if err := router.Run(":" + *port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
