package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var json RegisterRequest

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Request Valid",
		"data":   json,
	})
}
