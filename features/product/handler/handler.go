package handler

import (
	"net/http"
	productModel "star-pos/features/product/model"
	"star-pos/features/product/service"
	"star-pos/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	newProduct := productModel.Product{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	newID, err := service.AddProduct(newProduct)
	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	responseProduct := response.ResponseCreate{
		ID: newID,
	}

	return c.JSON(http.StatusCreated, responseProduct)
}

func GetAllProducts(c echo.Context) error {
	result, err := service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get all product: "+err.Error(), nil))
	}

	allProductResponse := []ProductResponse{}

	for _, data := range result {
		allProductResponse = append(allProductResponse, ProductResponse{
			ID:           data.ID,
			UserID:       data.UserID,
			ProductName:  data.ProductName,
			CategoriesID: data.CategoriesID,
			Stock:        data.Stock,
			Price:        data.Price,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, allProductResponse)
}

func GetProduct(c echo.Context) error {
	idParam := c.Param("id")
	result, err := service.GetProduct(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error get product: "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get product data: "+err.Error(), nil))
		}
	}

	responseProduct := ProductResponse{
		ID:           idParam,
		UserID:       result.UserID,
		ProductName:  result.ProductName,
		CategoriesID: result.CategoriesID,
		Stock:        result.Stock,
		Price:        result.Price,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, responseProduct)
}

func UpdateProduct(c echo.Context) error {
	idParam := c.Param("id")
	updateRequest := productModel.Product{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateRequest.ID = idParam

	err := service.UpdateProduct(updateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "didn't change") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error update product: "+err.Error(), nil))
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteProduct(c echo.Context) error {
	idParam := c.Param("id")
	err := service.DeleteProduct(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "first") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	return c.NoContent(http.StatusNoContent)
}
