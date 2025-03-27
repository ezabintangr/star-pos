package repository

import (
	"github.com/google/uuid"
	"star-pos/app/databases"
	cartModel "star-pos/features/cart/model"
)

func Get(id string) (*cartModel.Cart, error) {
	cart := cartModel.Cart{}
	tx := databases.DB.Where("id = ?", id).First(&cart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &cart, nil
}

func Create(cartId string, input cartModel.Cart) (string, error) {
	if cartId == "" {
		input.ID = uuid.NewString()
	} else {
		input.ID = cartId
	}
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return "", tx.Error
	}

	return input.ID, nil
}
