package handler

import "time"

type ProductResponse struct {
	ID           string    `json:"id" form:"id"`
	UserID       string    `json:"user_id" form:"user_id"`
	ProductName  string    `json:"product_name" form:"product_name"`
	CategoriesID string    `json:"categories_id" form:"categories_id"`
	Stock        float64   `json:"stock" form:"stock"`
	Price        float64   `json:"price" form:"price"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}
