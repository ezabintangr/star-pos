package repository

import (
	"time"
)

type User struct {
	ID          string `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserName    string `json:"user_name" form:"user_name"`
	PhoneNumber string `gorm:"unique" json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Role        string `json:"role" form:"role"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
