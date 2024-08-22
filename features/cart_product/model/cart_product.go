package cartProductModel

import (
	cartData "star-pos/features/cart/model"
	productData "star-pos/features/product/model"
	"time"
)

type CartProduct struct {
	ID        string              `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	CartID    string              `gorm:"type:char(36)" json:"cart_id" form:"cart_id"`
	ProductID string              `gorm:"type:char(36)" json:"product_id" form:"product_id"`
	Quantity  int64               `json:"quantity" form:"quantity"`
	CreatedAt time.Time           `json:"created_at" form:"created_at"`
	UpdatedAt time.Time           `json:"updated_at" form:"updated_at"`
	Cart      cartData.Cart       `gorm:"ForeignKey:CartID"`
	Product   productData.Product `gorm:"ForeignKey:ProductID"`
}
