package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
)

func (h *Handler) initManagerRoutes(api *gin.RouterGroup) {
	managers := api.Group("/managers")
	{
		managers.POST("/sign-in", h.managerSignIn)
		managers.POST("/auth/refresh", h.managerRefresh)

		manager := managers.Group("/manager", h.managerIdentity)
		{
			manager.POST("/create", h.managerCreateNew)
		}
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
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/sign-in [post]
func (h *Handler) managerSignIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err)
		return
	}

	res, err := h.services.Managers.SignIn(input)
	if err != nil {
		h.handleErrors(c, err)
		return
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
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/auth/refresh [post]
func (h *Handler) managerRefresh(c *gin.Context) {
	var input domain.RefreshInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, errBadRequestMessage)
		return
	}

	res, err := h.services.Managers.RefreshTokens(input.RefreshToken)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Create New Manager
// @Security ManagerAuth
// @Tags manager
// @Description manager creation
// @ModuleID managerCreateNew
// @Router       /managers/manager/create [post]
func (h *Handler) managerCreateNew(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
