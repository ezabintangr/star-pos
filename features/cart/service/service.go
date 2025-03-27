package service

import (
	cartModel "star-pos/features/cart/model"
	cartRepository "star-pos/features/cart/repository"
	"star-pos/features/cart_product/repository"
)

func AddProductToCart(input cartModel.AddProductToCart) (string, error) {
	var idCart string = input.CartId
	var err error
	if input.CartId == "" {
		idCart, err = createCart(input)
		if err != nil {
			return "", err
		}
	}

	idCreated, err := addProductToCart(idCart, input)
	if err != nil {
		return "", err
	}
	return idCreated, nil
}

func createCart(input cartModel.AddProductToCart) (string, error) {
	cart := cartModel.Cart{
		UserID:   "test",
		RefID:    "test",
		OutletID: input.OutletId,
	}

	return cartRepository.Create(cart)
}

func addProductToCart(idCart string, input cartModel.AddProductToCart) (string, error) {
	cartProduct, err := repository.AddProductToCart(idCart, input)
	if err != nil {
		return "", err
	}

	return cartProduct.ID, nil
}
