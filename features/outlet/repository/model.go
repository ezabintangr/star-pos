package repository

import (
	"star-pos/features/user/repository"
	"time"
)

type Outlet struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint
	OutletName  string
	Address     string
	PhoneOutlet string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        repository.User `gorm:"ForeignKey:UserID"`
}
