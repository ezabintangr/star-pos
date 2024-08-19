package repository

import (
	categoriesData "star-pos/features/categories/repository"
	userModel "star-pos/features/user/model"
	"time"
)

type Product struct {
	ID           string `gorm:"type:char(36);primarykey"`
	UserID       string `gorm:"type:char(36)"`
	ProductName  string
	CategoriesID string `gorm:"type:char(36)"`
	Stock        float64
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         userModel.User            `gorm:"ForeignKey:UserID"`
	Categories   categoriesData.Categories `gorm:"ForeignKey:CategoriesID"`
}
