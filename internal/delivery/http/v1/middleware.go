package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	managerCtx          = "managerID"
)

func (h *Handler) managerIdentity(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		h.handleErrors(c, err, err.Error())
		return
	}

	c.Set(managerCtx, id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New(domain.ErrEmptyAuthHeader)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New(domain.ErrInvalidAuthHeader)
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New(domain.ErrEmptyToken)
	}

	return h.tokenManager.Parse(headerParts[1])
}

func (h *Handler) getManagerID(c *gin.Context) (int, error) {
	id, ok := c.Get(managerCtx)
	if !ok {
		newResponse(c, http.StatusBadRequest, domain.ErrManagerIdNotFoundMessage)
		return -1, errors.New(domain.ErrManagerIdNotFound)
	}

	idToInt, err := strconv.Atoi(id.(string))
	if err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrManagerIdNotFoundMessage)
		return -1, errors.New(domain.ErrManagerIdNotFound)
	}

	return idToInt, nil
}
