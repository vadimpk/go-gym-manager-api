package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) handleErrors(ctx *gin.Context, err error, errorMessage string) {

	log.Println(err)

	switch errorMessage {
	case domain.ErrNotInDB:
		newResponse(ctx, http.StatusNotFound, domain.ErrNotInDBMessage)
	case domain.ErrBadRequest:
		newResponse(ctx, http.StatusBadRequest, domain.ErrBadRequestMessage)
	case domain.ErrEmptyAuthHeader, domain.ErrInvalidAuthHeader, domain.ErrEmptyToken:
		newResponse(ctx, http.StatusUnauthorized, domain.ErrNotAuthMessage)
	case domain.ErrStillInGym:
		newResponse(ctx, http.StatusForbidden, domain.ErrStillInGymMessage)
	case domain.ErrIsNotInGym:
		newResponse(ctx, http.StatusForbidden, domain.ErrIsNotInGymMessage)
	case domain.ErrDoesntHaveMembership:
		newResponse(ctx, http.StatusNotFound, domain.ErrDoesntHaveMembershipMessage)
	case domain.ErrExpiredMembership:
		newResponse(ctx, http.StatusForbidden, domain.ErrExpiredMembershipMessage)
	case domain.ErrExpiredToken:
		newResponse(ctx, http.StatusUnauthorized, domain.ErrExpiredTokenMessage)
	default:
		if strings.Contains(errorMessage, "duplicate key") {
			newResponse(ctx, http.StatusConflict, domain.ErrConflictMessage)
			return
		}
		newResponse(ctx, http.StatusInternalServerError, domain.ErrInternalServerMessage)
	}
}
