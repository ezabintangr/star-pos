package repository

import (
	"star-pos/features/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.RepositoryInterface.
func (u *userQuery) Insert(input user.UserCore) error {
	accountGorm := User{
		ID:          uuid.New().String(),
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	tx := u.db.Create(&accountGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
