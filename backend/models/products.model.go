package models

import "time"

type Product struct {
	ID          int       `json:"id_product" db:"id_product"`
	Name        string    `json:"name_product" db:"name_product"`
	Description string    `json:"desk_product" db:"desk_product"`
	Price       int       `json:"price_product" db:"price_product"`
	Stock       int       `json:"stock_product" db:"stock_product"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	IsActive    bool      `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
