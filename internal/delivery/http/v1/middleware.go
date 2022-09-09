package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	managerCtx          = "managerID"
)

func (h *Handler) managerIdentity(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		h.handleErrors(c, err)
		return
	}

	c.Set(managerCtx, id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New(errEmptyAuthHeader)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New(errInvalidAuthHeader)
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New(errEmptyToken)
	}

	return h.tokenManager.Parse(headerParts[1])
}
