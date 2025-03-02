package handler

import (
	"net/http"
	outletModel "star-pos/features/outlet/model"
	"star-pos/features/outlet/service"
	"star-pos/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateOutlet(c echo.Context) error {
	newOutlet := outletModel.Outlet{}
	errBind := c.Bind(&newOutlet)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	idCreated, err := service.Add(newOutlet)
	if err != nil {
		if strings.Contains(err.Error(), "user id") || strings.Contains(err.Error(), "name") || strings.Contains(err.Error(), "address") || strings.Contains(err.Error(), "phone") {
			return c.JSON(http.StatusBadRequest, response.WebJSONResponse("error create Outlet: "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error create Outlet: "+err.Error(), nil))
		}
	}

	result := response.ResponseCreate{
		ID: idCreated,
	}

	return c.JSON(http.StatusCreated, result)
}

func GetAllOutlets(c echo.Context) error {
	allOutlets, err := service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.WebJSONResponse("error get all outlets: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, allOutlets)
}

func GetOutlet(c echo.Context) error {
	idParam := c.Param("id")
	outlet, err := service.GetById(idParam)
	if err != nil {
		return response.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, outlet)
}

func UpdateOutlet(c echo.Context) error {
	idParam := c.Param("id")
	updateRequest := outletModel.Outlet{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return response.HandleBindError(c, errBind)
	}

	updateRequest.ID = idParam

	err := service.Update(updateRequest)
	if err != nil {
		return response.HandleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteOutlet(c echo.Context) error {
	idParam := c.Param("id")
	err := service.Delete(idParam)

	if err != nil {
		return response.HandleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
