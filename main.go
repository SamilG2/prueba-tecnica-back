package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/", func(context *gin.Context) {
		// Return JSON response
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go",
		})
	})
	// Start the server on port 8080
	r.Run()
}
