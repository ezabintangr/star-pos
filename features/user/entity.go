package user

import "time"

type UserCore struct {
	ID              string
	UserName        string
	PhoneNumber     string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ServiceInterface interface {
	Create(input UserCore) error
}

type RepositoryInterface interface {
	Insert(input UserCore) error
}
