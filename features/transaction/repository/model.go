package repository

import (
	outletData "star-pos/features/outlet/repository"
	userData "star-pos/features/user/repository"
	"time"
)

type Transaction struct {
	ID            uint `gorm:"primarykey"`
	UserID        uint
	OutletID      uint
	GrandTotal    float64
	PaymentStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          userData.User     `gorm:"ForeignKey:UserID"`
	Outlet        outletData.Outlet `gorm:"ForeignKey:OutletID"`
}
