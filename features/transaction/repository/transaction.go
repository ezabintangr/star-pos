package repository

import (
	outletData "star-pos/features/outlet/repository"
	userData "star-pos/features/user/repository"
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
	User          userData.User     `gorm:"ForeignKey:UserID"`
	Outlet        outletData.Outlet `gorm:"ForeignKey:OutletID"`
}
