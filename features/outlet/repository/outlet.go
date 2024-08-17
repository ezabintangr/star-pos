package repository

import (
	"star-pos/features/user/repository"
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
	User        repository.User `gorm:"ForeignKey:UserID"`
}
