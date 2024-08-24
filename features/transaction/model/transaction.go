package transactionModel

import (
	outletData "star-pos/features/outlet/model"
	userModel "star-pos/features/user/model"
	"time"
)

type Transaction struct {
	ID            string            `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserID        string            `gorm:"type:char(36)" json:"user_id" form:"user_id"`
	OutletID      string            `gorm:"type:char(36)" json:"outlet_id" form:"outlet_id"`
	GrandTotal    float64           `json:"grand_total" form:"grand_total"`
	PaymentStatus string            `json:"payment_status" form:"payment_status"`
	CreatedAt     time.Time         `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at" form:"updated_at"`
	User          userModel.User    `gorm:"ForeignKey:UserID"`
	Outlet        outletData.Outlet `gorm:"ForeignKey:OutletID"`
}
