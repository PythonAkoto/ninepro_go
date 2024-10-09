package main

import (
	"log"

	"github.com/PythonAkoto/ninepro_go/pkg/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Gin
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Define routes
	r.GET("/", handler.Index)
	r.GET("/home", handler.Home)
	r.POST("/contact", handler.Contact)

	// Run the server
	r.Run(":8080")
}
