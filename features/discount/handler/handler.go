package handler

import (
	"net/http"
	discountModel "star-pos/features/discount/model"
	"star-pos/features/discount/service"
	"star-pos/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateDiscounts(c echo.Context) error {
	newDiscount := discountModel.Discount{}
	errBind := c.Bind(&newDiscount)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	idCreated, err := service.Add(newDiscount)
	if err != nil {
		if strings.Contains(err.Error(), "user id") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create Discount: "+err.Error(), nil))
		} else if strings.Contains(err.Error(), "name") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create Discount: "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error create Discount: "+err.Error(), nil))
		}
	}

	result := response.ResponseCreate{
		ID: idCreated,
	}

	return c.JSON(http.StatusCreated, result)
}

func GetAllDiscounts(c echo.Context) error {
	result, err := service.GetAllDiscounts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get all Discounts: "+err.Error(), nil))
	}

	var responseDiscounts []DiscountsResponse

	for _, data := range result {
		responseDiscounts = append(responseDiscounts, DiscountsResponse{
			ID:           data.ID,
			UserID:       data.UserID,
			OutletID:     data.OutletID,
			ProductID:    data.ProductID,
			DiscountName: data.DiscountName,
			Amount:       data.Amount,
			DiscountType: data.DiscountType,
			DiscountMax:  data.DiscountMax,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
			User:         data.User,
			Outlet:       data.Outlet,
			Product:      data.Product,
		})
	}

	return c.JSON(http.StatusOK, responseDiscounts)
}

func GetCurrentDiscount(c echo.Context) error {
	idParam := c.Param("id")
	result, err := service.GetDiscount(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error get Discount: "+err.Error(), nil))
		} else if err.Error() == "record not found" {
			return c.NoContent(http.StatusNotFound)
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get Discount: "+err.Error(), nil))
		}
	}

	responseDiscount := DiscountsResponse{
		ID:           result.ID,
		UserID:       result.UserID,
		OutletID:     result.OutletID,
		ProductID:    result.ProductID,
		DiscountName: result.DiscountName,
		Amount:       result.Amount,
		DiscountType: result.DiscountType,
		DiscountMax:  result.DiscountMax,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
		User:         result.User,
		Outlet:       result.Outlet,
		Product:      result.Product,
	}

	return c.JSON(http.StatusOK, responseDiscount)
}

func UpdateDiscount(c echo.Context) error {
	idParam := c.Param("id")
	updateRequest := discountModel.Discount{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateRequest.ID = idParam

	err := service.UpdateCurrentDiscount(updateRequest)
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

func DeleteDiscount(c echo.Context) error {
	idParam := c.Param("id")
	err := service.DeleteDiscount(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse(err.Error(), nil))
		}
	}

	return c.NoContent(http.StatusNoContent)
}
