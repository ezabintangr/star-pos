package repository

import (
	"errors"
	"star-pos/app/databases"
	userModel "star-pos/features/user/model"

	"github.com/google/uuid"
)

func Insert(input *userModel.User) error {
	if databases.DB == nil {
		return errors.New("database connection is not initialized")
	}

	input.ID = uuid.New().String()
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadAllProfile() ([]userModel.User, error) {
	var allUser []userModel.User
	tx := databases.DB.Model(&userModel.User{}).Find(&allUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return allUser, nil
}

func ReadProfile(id string) (*userModel.User, error) {
	var ProfileUser userModel.User
	tx := databases.DB.Where("id = ?", id).First(&ProfileUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &ProfileUser, nil
}

func Update(input userModel.User) error {
	tx := databases.DB.Where("id = ?", input.ID).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func Login(input userModel.User) (*userModel.User, error) {
	var dataLogin userModel.User
	tx := databases.DB.Where("phone_number = ?", input.PhoneNumber).First(&dataLogin)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dataLogin, nil
}

func DeleteAccount(id string) error {
	tx := databases.DB.Delete(&userModel.User{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
