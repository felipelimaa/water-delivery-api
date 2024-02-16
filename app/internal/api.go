package internal

import (
	"felipelimaa/water-delivery-api/app/internal/infrastructure"
	"fmt"
	"log/slog"
	"net/http"
)

func StartApi() {
	slog.Info("Starting API")

	routes := infrastructure.NewRoutes()
	routes.ConfigApi()

	s := http.Server{
		Addr: fmt.Sprintf(":%s", infrastructure.PORT),
		Handler: routes.Echo,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error(fmt.Sprintf("Something went wrong when try to start API using PORT=%s", infrastructure.PORT), err)
	}
}
