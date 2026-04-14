package user

import (
	"backend/models"
)

type usersInterface interface {
	FindByEmail(email string) (*models.Users, error)
	Create(users *models.Users) error
}

// func FindByEmail(email string) usersInterface {

// }
