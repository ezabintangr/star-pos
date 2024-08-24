package discountModel

import (
	outletData "star-pos/features/outlet/model"
	productData "star-pos/features/product/model"
	userModel "star-pos/features/user/model"
	"time"
)

type Discount struct {
	ID           string              `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserID       string              `gorm:"type:char(36)" json:"user_id" form:"user_id"`
	OutletID     string              `gorm:"type:char(36)" json:"outlet_id" form:"outlet_id"`
	ProductID    string              `gorm:"type:char(36)" json:"product_id" form:"product_id"`
	DiscountName string              `json:"discount_name" form:"discount_name"`
	Amount       float64             `json:"amount" form:"amount"`
	DiscountType string              `json:"discount_type" form:"discount_type"`
	DiscountMax  float64             `json:"discount_max" form:"discount_max"`
	CreatedAt    time.Time           `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at" form:"updated_at"`
	User         userModel.User      `gorm:"ForeignKey:UserID"`
	Outlet       outletData.Outlet   `gorm:"ForeignKey:OutletID"`
	Product      productData.Product `gorm:"ForeignKey:ProductID"`
}
