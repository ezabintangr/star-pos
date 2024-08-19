package service

import (
	"errors"
	"star-pos/features/user/repository"
	encrypts "star-pos/utils"
)

func Create(input repository.User) error {
	if input.PhoneNumber == "" {
		return errors.New("phone is required")
	}

	hashed, errhash := encrypts.NewHashService().HashPassword(input.Password)
	if errhash != nil {
		return errors.New("error hashing password")
	}

	input.Password = hashed

	err := repository.Insert(&input)
	if err != nil {
		return err
	}

	return nil
}
