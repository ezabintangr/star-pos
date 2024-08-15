package repository

import (
	"star-pos/features/outlet/repository"
	userData "star-pos/features/user/repository"
	"time"
)

type Cart struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	RefID     string
	OutletID  uint
	Date      time.Time `gorm:"type:date"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      userData.User     `gorm:"ForeignKey:UserID"`
	Outlet    repository.Outlet `gorm:"ForeignKey:OutletID"`
}
