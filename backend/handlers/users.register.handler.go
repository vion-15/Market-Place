package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Phone    string `json:"phone" binding:"required,min=9,max=15"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

func RegisterHandler(c *gin.Context) {
	// menginisialisasi struct RegisterRequest
	var json RegisterRequest

	// menangkap dan memasukan data dari body dan disimpan kedalam variable json,
	// jika error return kode, status dan message
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format data tidak valid, pastikan field diisi dengan sesuai",
			"dbg_msg": err.Error(),
		})
		return
	}

	// mapping handler -> services
	req := services.RegisterRequest{
		Email:    json.Email,
		Name:     json.Name,
		Phone:    json.Phone,
		Password: json.Password,
	}

	// memasukan mapping kedalam fungsi
	// !mengikat handler kedalam struct yang memiliki depedency ke service (interface)
	result, err := services.Register(req)

	// error handler mapping
	// !logic mengecek status error dari service
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// success handler mapping
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "register berhasil",
		"data":    result,
	})
}
