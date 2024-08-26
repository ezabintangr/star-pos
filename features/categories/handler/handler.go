package handler

import (
	"net/http"
	categoriesModel "star-pos/features/categories/model"
	"star-pos/features/categories/service"
	"star-pos/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateCategories(c echo.Context) error {
	newCategory := categoriesModel.Categories{}
	errBind := c.Bind(&newCategory)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	idCreated, err := service.AddCategories(newCategory)
	if err != nil {
		if strings.Contains(err.Error(), "user id") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create category: "+err.Error(), nil))
		} else if strings.Contains(err.Error(), "name") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create category: "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error create category: "+err.Error(), nil))
		}
	}

	result := response.ResponseCreate{
		ID: idCreated,
	}

	return c.JSON(http.StatusCreated, result)
}

func GetAllCategories(c echo.Context) error {
	result, err := service.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get all categories: "+err.Error(), nil))
	}

	var responseCategories []CategoriesResponse

	for _, data := range result {
		responseCategories = append(responseCategories, CategoriesResponse{
			ID:           data.ID,
			UserID:       data.UserID,
			CategoryName: data.UserID,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, responseCategories)
}

func GetCurrentCategory(c echo.Context) error {
	idParam := c.Param("id")
	result, err := service.GetCategory(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error get category: "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get category: "+err.Error(), nil))
		}
	}

	responseCategory := CategoriesResponse{
		ID:           idParam,
		UserID:       result.UserID,
		CategoryName: result.CategoryName,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, responseCategory)
}

func UpdateCategory(c echo.Context) error {
	idParam := c.Param("id")
	updateRequest := categoriesModel.Categories{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateRequest.ID = idParam

	err := service.UpdateCurrentCategory(updateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteCategory(c echo.Context) error {
	idParam := c.Param("id")
	err := service.DeleteCategory(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	return c.NoContent(http.StatusNoContent)
}
