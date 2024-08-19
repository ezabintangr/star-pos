package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase(database *gorm.DB) {
	db = database
}

func Insert(input *User) error {
	if db == nil {
		return errors.New("database connection is not initialized")
	}

	input.ID = uuid.New().String()
	tx := db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
