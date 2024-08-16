package handler

import (
	"net/http"
	"star-pos/features/user"
	"star-pos/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService user.ServiceInterface
}

func New(us user.ServiceInterface) *userHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response.WebJSONHelloResponse("hello world"))
}

func (uh *userHandler) CreateAccount(c echo.Context) error {
	newRequest := RequestAccount{}
	errbind := c.Bind(&newRequest)
	if errbind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errbind.Error(), nil))
	}

	newAccountCore := user.UserCore{
		PhoneNumber:     newRequest.PhoneNumber,
		Password:        newRequest.Password,
		PasswordConfirm: newRequest.PasswordConfirm,
	}

	errCreate := uh.userService.Create(newAccountCore)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "phone") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create account: "+errCreate.Error(), nil))
		} else if strings.Contains(errCreate.Error(), "password") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create account: "+errCreate.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error create account:"+errCreate.Error(), nil))
		}
	}

	return c.JSON(http.StatusCreated, response.WebJSONResponse("success create account", nil))
}
