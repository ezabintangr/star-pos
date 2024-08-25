package handler

import (
	"net/http"
	userModel "star-pos/features/user/model"
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

	requestAccount := userModel.User{
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

func GetAllProfile(c echo.Context) error {
	result, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get all user: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.WebJSONResponse("success get all user", result))
}

func GetProfile(c echo.Context) error {
	idParam := c.Param("id")
	result, err := service.GetProfile(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error get profile: "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get profile data: "+err.Error(), nil))
		}
	}

	responseProfile := ResponseUser{
		ID:          idParam,
		UserName:    result.UserName,
		PhoneNumber: result.PhoneNumber,
		Email:       result.Email,
		Role:        result.Role,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response.WebJSONResponse("success get profile", responseProfile))
}

func UpdateProfile(c echo.Context) error {
	idParam := c.Param("id")
	updateRequest := userModel.User{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateRequest.ID = idParam

	err := service.UpdateProfile(updateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "not change") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error update profile: "+err.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, response.WebJSONResponse("update successful", nil))
}

func DeleteAccount(c echo.Context) error {
	idParam := c.Param("id")
	err := service.Delete(idParam)
	if err != nil {
		if strings.ContainsAny(err.Error(), "first") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, response.WebJSONResponse("delete account successfully", nil))
}

func Login(c echo.Context) error {
	var LoginRequest userModel.User
	errBind := c.Bind(&LoginRequest)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	data, token, err := service.Login(LoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "fill") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else if strings.Contains(err.Error(), "password") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	responseLogin := map[string]any{
		"data":  data,
		"token": token,
	}

	return c.JSON(http.StatusOK, response.WebJSONResponse("login successful", responseLogin))

}
