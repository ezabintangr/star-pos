package repository

import (
	"github.com/google/uuid"
	"star-pos/app/databases"
	cartModel "star-pos/features/cart/model"
	cartProductModel "star-pos/features/cart_product/model"
)

func AddProductToCart(idCart string, request cartModel.AddProductToCart) (*cartProductModel.CartProduct, error) {
	cartProduct := cartProductModel.CartProduct{
		ID:        uuid.New().String(),
		CartID:    idCart,
		ProductID: request.ProductId,
		Quantity:  request.Quantity,
	}

	tx := databases.DB.Create(&cartProduct)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &cartProduct, nil
}

func GetProductsFromCart(idCart string) ([]cartProductModel.CartProduct, error) {
	var cartProducts []cartProductModel.CartProduct
	tx := databases.DB.Model(&cartProductModel.CartProduct{}).Where("cart_id = ?", idCart).Find(&cartProducts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return cartProducts, nil
}
