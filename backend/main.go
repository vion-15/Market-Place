package main

import (
	"backend/handlers"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// depedency injection
	userRepo := repositories.NewUserRepositories()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/register", userHandler.RegisterHandler)

	router.Run("localhost:8080")
}
