package handler

import "time"

type CategoriesResponse struct {
	ID           string    `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserID       string    `gorm:"type:char(36)" json:"user_id" form:"user_id"`
	CategoryName string    `json:"category_name" form:"category_name"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"created_at"`
}
