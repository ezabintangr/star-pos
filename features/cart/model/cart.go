package cartModel

import (
	outletData "star-pos/features/outlet/model"
	userModel "star-pos/features/user/model"
	"time"
)

type Cart struct {
	ID        string            `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserID    string            `gorm:"type:char(36)" json:"user_id" form:"user_id"`
	RefID     string            `json:"ref_id" form:"ref_id"`
	OutletID  string            `gorm:"type:char(36)" json:"outlet_id" form:"outlet_id"`
	Date      time.Time         `gorm:"type:date" json:"date" form:"date"`
	CreatedAt time.Time         `json:"created_at" form:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" form:"updated_at"`
	User      userModel.User    `gorm:"ForeignKey:UserID"`
	Outlet    outletData.Outlet `gorm:"ForeignKey:OutletID"`
}
