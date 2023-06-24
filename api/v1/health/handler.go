package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "Ok"})
}
