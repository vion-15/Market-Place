package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaimsCustom struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, role string) (string, error) {

	claims := &JwtClaimsCustom{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		return "", errors.New("Kode Rahasia Tidak Ada")
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", errors.New("Gagal Membuat Token")
	}

	return tokenString, nil
}
