package service

import (
	"errors"
	"star-pos/app/middlewares"
	userModel "star-pos/features/user/model"
	"star-pos/features/user/repository"
	encrypts "star-pos/utils"
)

func Create(input userModel.User) error {
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

func GetProfile(id string) (*userModel.User, error) {
	if id == "" {
		return nil, errors.New("you must login first")
	}

	result, err := repository.ReadProfile(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateProfile(input userModel.User) error {
	if input.ID == "" {
		return errors.New("you must login first")
	}

	if input.UserName == "" && input.PhoneNumber == "" && input.Email == "" && input.Password == "" && input.Role == "" {
		return errors.New("this field is required")
	}

	err := repository.Update(input)
	if err != nil {
		return err
	}

	return nil
}

func Login(input userModel.User) (data *userModel.User, token string, err error) {
	if input.PhoneNumber == "" || input.Password == "" {
		return nil, "", errors.New("fill this field first")
	}

	result, err := repository.Login(input)
	if err != nil {
		return nil, "", err
	}

	isPasswordValid := encrypts.NewHashService().CheckPasswordHash(result.Password, input.Password)
	if !isPasswordValid {
		return nil, "", errors.New("incorrect password")
	}

	tokenCreated, errJWT := middlewares.CreateToken(result.ID)
	if errJWT != nil {
		return nil, "", errJWT
	}

	return result, tokenCreated, nil
}

func Delete(id string) error {
	if id == "" {
		return errors.New("you must login first")
	}

	err := repository.DeleteAccount(id)
	if err != nil {
		return err
	}

	return nil
}
