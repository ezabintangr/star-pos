package repository

import "time"

type User struct {
	ID          uint `gorm:"primarykey"`
	UserName    string
	PhoneNumber string
	Email       string
	Password    string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
