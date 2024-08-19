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
