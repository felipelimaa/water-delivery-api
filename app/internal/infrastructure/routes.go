package infrastructure

import (
	"felipelimaa/water-delivery-api/app/internal/service"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	Echo *echo.Echo
}

func NewRoutes() *Routes {
	return &Routes{Echo: echo.New()}
}

func (r *Routes) ConfigApi() {
	r.Echo.GET("/healthcheck", service.HandleHealthCheck)
}
