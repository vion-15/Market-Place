package main

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/register", handlers.RegisterHandler)

	router.Run("localhost:8080")
}
