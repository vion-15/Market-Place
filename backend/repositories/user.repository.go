package repositories

import (
	"backend/models"
	"database/sql"
	"errors"
)

type UserRepositories interface {
	FindByEmail(email string) (*models.Users, error)
	FindByNoTelp(noTelp string) (*models.Users, error)
	Create(user *models.Users) error
}

type userPostgresRepositories struct {
	db *sql.DB
}

func (d *userPostgresRepositories) FindByEmail(email string) (*models.Users, error) {
	var user models.Users

	query := "SELECT id, name, email, phone, role FROM users WHERE email = ($1)"

	err := d.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Data tidak ditemukan")
		}
		return nil, err
	}
	return &user, nil
}

func (d *userPostgresRepositories) FindByNoTelp(noTelp string) (*models.Users, error) {

	var user models.Users

	query := "SELECT id, name, email, phone, role FROM users WHERE phone = ($1)"

	err := d.db.QueryRow(query, noTelp).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Data tidak ditemukan")
		}
		return nil, err
	}

	return &user, nil
}

func (d *userPostgresRepositories) Create(user *models.Users) error {
	_, err := d.db.Exec("INSERT INTO users (name, email, phone, password, is_email_verified, is_phone_verified, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", user.Name, user.Email, user.Phone, user.Password, user.IsEmailVerified, user.IsPhoneVerified, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// depedency injection
func NewUserPostgresRepositories(db *sql.DB) UserRepositories {
	return &userPostgresRepositories{
		db: db,
	}
}
