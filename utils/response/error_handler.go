package response

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	if strings.Contains(err.Error(), "login") {
		return c.JSON(http.StatusBadRequest, WebJSONResponse("error get Outlet: "+err.Error(), nil))
	} else if strings.Contains(err.Error(), "not found") {
		return c.NoContent(http.StatusNotFound)
	} else {
		return c.JSON(http.StatusInternalServerError, WebJSONResponse("error get Outlet: "+err.Error(), nil))
	}
}

func HandleBindError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, WebJSONResponse("error bind data: "+err.Error(), nil))
}
