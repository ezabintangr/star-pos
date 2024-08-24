package outletModel

import (
	userModel "star-pos/features/user/model"
	"time"
)

type Outlet struct {
	ID          string         `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserID      string         `gorm:"type:char(36)" json:"user_id" form:"user_id"`
	OutletName  string         `json:"outlet_name" form:"outlet_name"`
	Address     string         `json:"address" form:"address"`
	PhoneOutlet string         `json:"phone_outlet" form:"phone_outlet"`
	CreatedAt   time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" form:"created_at"`
	User        userModel.User `gorm:"ForeignKey:UserID"`
}
