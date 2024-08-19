package repository

import (
	outletData "star-pos/features/outlet/repository"
	productData "star-pos/features/product/repository"
	userModel "star-pos/features/user/model"
	"time"
)

type Discount struct {
	ID           string `gorm:"type:char(36);primarykey"`
	UserID       string `gorm:"type:char(36)"`
	OutletID     string `gorm:"type:char(36)"`
	ProductID    string `gorm:"type:char(36)"`
	DiscountName string
	Amount       float64
	DiscountType string
	DiscountMax  float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         userModel.User      `gorm:"ForeignKey:UserID"`
	Outlet       outletData.Outlet   `gorm:"ForeignKey:OutletID"`
	Product      productData.Product `gorm:"ForeignKey:ProductID"`
}
