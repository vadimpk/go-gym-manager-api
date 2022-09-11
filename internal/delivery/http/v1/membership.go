package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
	"strconv"
)

// @Summary Create Membership
// @Security ManagerAuth
// @Tags memberships
// @Description membership creation
// @ModuleID membershipCreateNew
// @Accept  json
// @Produce  json
// @Param input body domain.MembershipCreateInput true "membership info"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/memberships/create [post]
func (h *Handler) membershipCreateNew(c *gin.Context) {
	var input domain.MembershipCreateInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err, domain.ErrBadRequest)
		return
	}

	id, err := h.services.Memberships.CreateNew(input)
	if err != nil {
		h.handleErrors(c, err, err.Error())
		return
	}
	newDataResponse(c, http.StatusOK, dataResponse{
		Message: domain.MessageMembershipCreated,
		Data: map[string]int{
			"id": id,
		},
	})
}

// @Summary Get Membership By ID
// @Security ManagerAuth
// @Tags memberships
// @Description get membership by id
// @ModuleID membershipGetByID
// @Produce  json
// @Param id path int true "Membership ID"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/memberships/get/{id} [get]
func (h *Handler) membershipGetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, err, domain.ErrBadRequest)
		return
	}

	membership, err := h.services.Memberships.GetByID(id)
	if err != nil {
		h.handleErrors(c, err, err.Error())
		return
	}
	newDataResponse(c, http.StatusOK, dataResponse{
		Message: domain.MessageMembershipFound,
		Data:    membership,
	})
}

// @Summary Update Membership By ID
// @Security ManagerAuth
// @Tags memberships
// @Description update membership by id with json body
// @ModuleID membershipUpdateByID
// @Accept json
// @Produce  json
// @Param id path int true "Membership ID"
// @Param input body domain.MembershipUpdateInput true "membership update info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/memberships/update/{id} [put]
func (h *Handler) membershipUpdateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, err, domain.ErrBadRequest)
		return
	}

	var input domain.MembershipUpdateInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err, err.Error())
		return
	}

	err = h.services.Memberships.UpdateByID(id, input)
	if err != nil {
		h.handleErrors(c, err, err.Error())
		return
	}
	newResponse(c, http.StatusOK, domain.MessageMembershipUpdated)
}

// @Summary Delete Membership By ID
// @Security ManagerAuth
// @Tags memberships
// @Description delete membership by id
// @ModuleID membershipDeleteByID
// @Produce  json
// @Param id path int true "Membership ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/memberships/delete/{id} [delete]
func (h *Handler) membershipDeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, err, domain.ErrBadRequest)
		return
	}

	err = h.services.Memberships.DeleteByID(id)
	if err != nil {
		h.handleErrors(c, err, err.Error())
		return
	}
	newResponse(c, http.StatusOK, domain.MessageMembershipDeleted)
}
