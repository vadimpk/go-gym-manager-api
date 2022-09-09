package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/service"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, manager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: manager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initManagerRoutes(v1)
	}
}
