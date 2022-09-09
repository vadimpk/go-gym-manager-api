package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	errNotInDB           = "sql: no rows in result set"
	errBadRequest        = "bad request"
	errEmptyAuthHeader   = "empty auth header"
	errInvalidAuthHeader = "invalid auth header"
	errEmptyToken        = "token is empty"

	errNotAuthMessage        = "Not authorized"
	errBadRequestMessage     = "Input data is incorrect"
	errNotInDBMessage        = "No results found in database"
	errInternalServerMessage = "Server is not responding at the moment"
)

func (h *Handler) handleErrors(ctx *gin.Context, err error) {

	log.Println(err)

	switch err.Error() {
	case errNotInDB:
		newResponse(ctx, http.StatusNotFound, errNotInDBMessage)
	case errBadRequest:
		newResponse(ctx, http.StatusBadRequest, errBadRequestMessage)
	case errEmptyAuthHeader, errInvalidAuthHeader, errEmptyToken:
		newResponse(ctx, http.StatusUnauthorized, errNotAuthMessage)
	default:
		newResponse(ctx, http.StatusInternalServerError, errInternalServerMessage)
	}
}
