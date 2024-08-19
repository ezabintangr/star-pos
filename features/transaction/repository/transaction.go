package repository

import (
	outletData "star-pos/features/outlet/repository"
	userModel "star-pos/features/user/model"
	"time"
)

type Transaction struct {
	ID            string `gorm:"type:char(36);primarykey"`
	UserID        string `gorm:"type:char(36)"`
	OutletID      string `gorm:"type:char(36)"`
	GrandTotal    float64
	PaymentStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          userModel.User    `gorm:"ForeignKey:UserID"`
	Outlet        outletData.Outlet `gorm:"ForeignKey:OutletID"`
}
