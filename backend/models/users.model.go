package models

import (
	"time"
)

type Users struct {
	ID              int64      `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	Email           string     `json:"email" db:"email"`
	Phone           string     `json:"phone" db:"phone"`
	Password        string     `json:"-" db:"password"`
	PhotoProfile    string     `json:"link" db:"link"`
	IsEmailVerified bool       `json:"is_email_verified" db:"is_email_verified"`
	IsPhoneVerified bool       `json:"is_phone_verified" db:"is_phone_verified"`
	Role            string     `json:"role"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
	LastLogin       *time.Time `json:"las_login" db:"last_login"`
}
