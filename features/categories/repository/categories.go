package repository

import (
	"star-pos/features/user/repository"
	"time"
)

type Categories struct {
	ID           string `gorm:"type:char(36);primarykey"`
	UserID       string `gorm:"type:char(36)"`
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         repository.User `gorm:"ForeignKey:UserID"`
}
