package repository

import (
	"star-pos/app/databases"
	cartModel "star-pos/features/cart/model"

	"github.com/google/uuid"
)

func Get(id string) (*cartModel.Cart, error) {
	cart := cartModel.Cart{}
	tx := databases.DB.Where("id = ?", id).First(&cart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &cart, nil
}

func Create(input cartModel.Cart) (string, error) {
	input.ID = uuid.NewString()
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return "", tx.Error
	}

	return input.ID, nil
}
