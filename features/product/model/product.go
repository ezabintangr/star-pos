package productModel

import (
	categoriesData "star-pos/features/categories/model"
	userModel "star-pos/features/user/model"
	"time"
)

type Product struct {
	ID           string                    `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserID       string                    `gorm:"type:char(36)" json:"user_id" form:"user_id"`
	ProductName  string                    `validate:"required" json:"product_name" form:"product_name"`
	CategoriesID string                    `gorm:"type:char(36)" json:"categories_id" form:"categories_id"`
	Stock        float64                   `validate:"required" json:"stock" form:"stock"`
	Price        float64                   `validate:"required" json:"price" form:"price"`
	CreatedAt    time.Time                 `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time                 `json:"updated_at" form:"updated_at"`
	User         userModel.User            `gorm:"ForeignKey:UserID"`
	Categories   categoriesData.Categories `gorm:"ForeignKey:CategoriesID"`
}
