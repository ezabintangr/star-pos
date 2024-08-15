package repository

import (
	categoriesData "star-pos/features/categories/repository"
	userData "star-pos/features/user/repository"
	"time"
)

type Product struct {
	ID           uint `gorm:"primarykey"`
	UserID       uint
	ProductName  string
	CategoriesID uint
	Stock        float64
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         userData.User             `gorm:"ForeignKey:UserID"`
	Categories   categoriesData.Categories `gorm:"ForeignKey:CategoriesID"`
}
