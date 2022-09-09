package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
)

func (h *Handler) initManagerRoutes(api *gin.RouterGroup) {
	users := api.Group("/manager")
	{
		users.POST("/sign-in", h.managerSignIn)
		users.POST("/auth/refresh", h.managerRefresh)
	}
}

// @Summary Manager SignIn
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

	res, err := h.services.Managers.SignIn(c, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) managerRefresh(c *gin.Context) {

}
