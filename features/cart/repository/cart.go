package repository

import (
	outletData "star-pos/features/outlet/repository"
	userData "star-pos/features/user/repository"
	"time"
)

type Cart struct {
	ID        string `gorm:"type:char(36);primarykey"`
	UserID    string `gorm:"type:char(36)"`
	RefID     string
	OutletID  string    `gorm:"type:char(36)"`
	Date      time.Time `gorm:"type:date"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      userData.User     `gorm:"ForeignKey:UserID"`
	Outlet    outletData.Outlet `gorm:"ForeignKey:OutletID"`
}
