package handler

import (
	"github.com/fakelozic/go-boilerplate/internal/server"
	"github.com/fakelozic/go-boilerplate/internal/service"
)

type Handlers struct {
	Health *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health: NewHealthHandler(s),
	}
}
