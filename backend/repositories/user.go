package repositories

import (
	"backend/models"
	"errors"
)

type UserRepositories interface {
	FindByEmail(email string) (*models.Users, error)
	FindByNoTelp(noTelp string) (*models.Users, error)
	Create(user *models.Users) error
}

type userRepositories struct {
	data []*models.Users
}

func (d *userRepositories) FindByEmail(email string) (*models.Users, error) {
	for _, u := range d.data {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("Data tidak ditemukan")
}

func (d *userRepositories) FindByNoTelp(noTelp string) (*models.Users, error) {
	for _, u := range d.data {
		if u.Phone == noTelp {
			return u, nil
		}
	}
	return nil, errors.New("Data tidak ditemukan")
}

func (d *userRepositories) Create(user *models.Users) error {
	d.data = append(d.data, user)
	return nil
}

// depedency injection
func NewUserRepositories() UserRepositories {
	return &userRepositories{}
}
