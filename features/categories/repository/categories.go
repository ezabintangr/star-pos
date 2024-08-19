package repository

import (
	userModel "star-pos/features/user/model"
	"time"
)

type Categories struct {
	ID           string `gorm:"type:char(36);primarykey"`
	UserID       string `gorm:"type:char(36)"`
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         userModel.User `gorm:"ForeignKey:UserID"`
}
