package handler

import (
	"net/http"
	"star-pos/features/user/repository"
	"star-pos/features/user/service"
	"star-pos/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response.WebJSONHelloResponse("hello world"))
}

func CreateAccount(c echo.Context) error {
	newRequest := RequestAccount{}
	errbind := c.Bind(&newRequest)
	if errbind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errbind.Error(), nil))
	}

	if newRequest.Password == "" || newRequest.PasswordConfirm == "" {
		return c.JSON(http.StatusBadRequest, response.WebJSONResponse("this field is required", nil))
	} else if newRequest.PasswordConfirm != newRequest.Password {
		return c.JSON(http.StatusBadRequest, response.WebJSONResponse("password confirm didn't match", nil))
	}

	requestAccount := repository.User{
		PhoneNumber: newRequest.PhoneNumber,
		Password:    newRequest.Password,
	}

	errCreate := service.Create(requestAccount)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "phone") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create account: "+errCreate.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error create account:"+errCreate.Error(), nil))
		}
	}

	return c.JSON(http.StatusCreated, response.WebJSONResponse("success create account", nil))
}
