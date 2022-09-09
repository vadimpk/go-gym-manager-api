package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
)

func (h *Handler) initManagerRoutes(api *gin.RouterGroup) {
	managers := api.Group("/manager")
	{
		managers.POST("/sign-in", h.managerSignIn)
		managers.POST("/auth/refresh", h.managerRefresh)
	}
}

// @Summary Manager Sign In
// @Tags manager-auth
// @Description manager sign in
// @ModuleID managerSignIn
// @Accept  json
// @Produce  json
// @Param input body domain.SignInInput true "sign up info"
// @Success 200 {object} service.Tokens
// @Router       /manager/sign-in [post]
func (h *Handler) managerSignIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
	}

	res, err := h.services.Managers.SignIn(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Manager Refresh Tokens
// @Tags manager-auth
// @Description manager refresh
// @ModuleID managerRefresh
// @Accept  json
// @Produce  json
// @Param input body domain.RefreshInput true "refresh info"
// @Success 200 {object} service.Tokens
// @Router       /manager/auth/refresh [post]
func (h *Handler) managerRefresh(c *gin.Context) {
	var input domain.RefreshInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
	}

	res, err := h.services.Managers.RefreshTokens(input.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
