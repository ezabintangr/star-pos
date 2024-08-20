package handler

import "time"

type ResponseUser struct {
	ID          string    `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	UserName    string    `json:"user_name" form:"user_name"`
	PhoneNumber string    `gorm:"unique" json:"phone_number" form:"phone_number"`
	Email       string    `json:"email" form:"email"`
	Role        string    `json:"role" form:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
