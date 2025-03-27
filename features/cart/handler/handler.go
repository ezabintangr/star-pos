package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"star-pos/app/middlewares"
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

	return c.JSON(http.StatusOK, cartProduct)
}

func AddProductToCartHandler(c echo.Context) error {
	id := c.Param("id")
	userId := middlewares.ExtractTokenUserId(c)
	request := cartModel.AddProductToCart{}
	errBind := c.Bind(&request)
	if errBind != nil {
		return response.HandleBindError(c, errBind)
	}

	idCreated, err := service.AddProductToCart(id, request, userId)
	if err != nil {
		return response.HandleError(c, err)
	}

	return c.JSON(http.StatusCreated, response.ResponseCreate{ID: idCreated})
}
