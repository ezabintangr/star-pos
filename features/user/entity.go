package user

import "time"

type UserCore struct {
	ID              string    `json:"id"`
	UserName        string    `json:"user_name"`
	PhoneNumber     string    `json:"phone_number"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	Role            string    `json:"role"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ServiceInterface interface {
	Create(input UserCore) error
}

type RepositoryInterface interface {
	Insert(input UserCore) error
}
