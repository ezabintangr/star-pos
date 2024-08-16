package repository

import "time"

type User struct {
	ID          string `gorm:"type:char(36);primarykey"`
	UserName    string
	PhoneNumber string `gorm:"unique"`
	Email       string
	Password    string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
