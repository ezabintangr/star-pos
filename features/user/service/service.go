package service

import (
	"errors"
	"star-pos/features/user"
	encrypts "star-pos/utils"
)

type userService struct {
	ur user.RepositoryInterface
	uh encrypts.HashInterface
}

func New(ur user.RepositoryInterface, uh encrypts.HashInterface) user.ServiceInterface {
	return &userService{
		ur: ur,
		uh: uh,
	}
}

// Create implements user.ServiceInterface.
func (u *userService) Create(input user.UserCore) error {
	if input.PhoneNumber == "" {
		return errors.New("phone number is required")
	} else if input.Password == "" {
		return errors.New("password is required")
	} else if input.PasswordConfirm == "" {
		return errors.New("password confirm is required")
	}

	if input.PasswordConfirm != input.Password {
		return errors.New("password doesn't match")
	}

	hashed, errhash := u.uh.HashPassword(input.Password)
	if errhash != nil {
		return errhash
	}

	input.Password = hashed

	err := u.ur.Insert(input)
	if err != nil {
		return err
	}

	return nil
}
