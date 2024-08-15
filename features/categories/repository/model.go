package repository

import (
	"star-pos/features/user/repository"
	"time"
)

type Categories struct {
	ID           uint `gorm:"primarykey"`
	UserID       uint
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         repository.User `gorm:"ForeignKey:UserID"`
}