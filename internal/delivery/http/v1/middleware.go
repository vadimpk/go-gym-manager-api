package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
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

func (h *Handler) getManagerID(c *gin.Context) (int, error) {
	id, ok := c.Get(managerCtx)
	if !ok {
		newResponse(c, http.StatusInternalServerError, "manager id not found")
		return -1, errors.New("manager id not found")
	}

	idToStr, ok := id.(string)
	if !ok {
		newResponse(c, http.StatusInternalServerError, "manager id not found")
		return -1, errors.New("manager id not found")
	}

	idToInt, err := strconv.Atoi(idToStr)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "manager id not found")
		return -1, errors.New("manager id not found")
	}

	return idToInt, nil
}
