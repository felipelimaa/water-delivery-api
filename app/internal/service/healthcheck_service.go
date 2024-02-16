package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleHealthCheck(c echo.Context) error {
	c.JSON(http.StatusOK, "The API is health")
	return nil
}
