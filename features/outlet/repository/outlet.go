package repository

import (
	userModel "star-pos/features/user/model"
	"time"
)

type Outlet struct {
	ID          string `gorm:"type:char(36);primarykey"`
	UserID      string `gorm:"type:char(36)"`
	OutletName  string
	Address     string
	PhoneOutlet string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        userModel.User `gorm:"ForeignKey:UserID"`
}
