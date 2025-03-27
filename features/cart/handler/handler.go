package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	cartModel "star-pos/features/cart/model"
	cartRepository "star-pos/features/cart/repository"
	"star-pos/features/cart/service"
	cartProductRepository "star-pos/features/cart_product/repository"
	"star-pos/utils/response"
)

func Get(c echo.Context) error {
	id := c.Param("id")
	cart, err := cartRepository.Get(id)
	if err != nil {
		return response.HandleError(c, err)
	}

	cartProduct, err := cartProductRepository.GetProductsFromCart(cart.ID)
	if err != nil {
		return response.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, cart)
}

func AddProductToCartHandler(c echo.Context) error {
	request := cartModel.AddProductToCart{}
	errBind := c.Bind(&request)
	if errBind != nil {
		return response.HandleBindError(c, errBind)
	}

	idCreated, err := service.AddProductToCart(request)
	if err != nil {
		return response.HandleError(c, err)
	}

	return c.JSON(http.StatusCreated, response.ResponseCreate{ID: idCreated})
}
