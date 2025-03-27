package service

import (
	cartModel "star-pos/features/cart/model"
	cartRepository "star-pos/features/cart/repository"
	"star-pos/features/cart_product/repository"
)

func AddProductToCart(cartId string, input cartModel.AddProductToCart, userId string) (string, error) {
	var err error
	cart, err := cartRepository.Get(cartId)
	if err != nil && err.Error() != "record not found" {
		return "", err
	}
	if cart == nil {
		_, err = createCart(cartId, input, userId)
		if err != nil {
			return "", err
		}
	}

	idCreated, err := addProductToCart(cartId, input)
	if err != nil {
		return "", err
	}
	return idCreated, nil
}

func createCart(cartId string, input cartModel.AddProductToCart, userId string) (string, error) {
	cart := cartModel.Cart{
		UserID:   userId,
		OutletID: input.OutletId,
	}

	return cartRepository.Create(cartId, cart)
}

func addProductToCart(idCart string, input cartModel.AddProductToCart) (string, error) {
	cartProduct, err := repository.AddProductToCart(idCart, input)
	if err != nil {
		return "", err
	}

	return cartProduct.ID, nil
}
