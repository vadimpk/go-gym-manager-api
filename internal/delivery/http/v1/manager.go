package v1

import "github.com/gin-gonic/gin"

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
// @Router       /manager/ [post]
func (h *Handler) managerSignIn(c *gin.Context) {

}

func (h *Handler) managerRefresh(c *gin.Context) {

}
