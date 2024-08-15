package repository

import (
	cartData "star-pos/features/cart/repository"
	productData "star-pos/features/product/repository"
	"time"
)

type CartProduct struct {
	ID        uint `gorm:"primarykey"`
	CartID    uint
	ProductID uint
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Cart      cartData.Cart       `gorm:"ForeignKey:CartID"`
	Product   productData.Product `gorm:"ForeignKey:ProductID"`
}
