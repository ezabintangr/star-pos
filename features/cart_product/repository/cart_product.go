package repository

import (
	cartData "star-pos/features/cart/repository"
	productData "star-pos/features/product/repository"
	"time"
)

type CartProduct struct {
	ID        string `gorm:"type:char(36);primarykey"`
	CartID    string `gorm:"type:char(36)"`
	ProductID string `gorm:"type:char(36)"`
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Cart      cartData.Cart       `gorm:"ForeignKey:CartID"`
	Product   productData.Product `gorm:"ForeignKey:ProductID"`
}
