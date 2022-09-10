package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
	"strconv"
)

// @Summary Create Member
// @Security ManagerAuth
// @Tags members
// @Description member creation
// @ModuleID memberCreateNew
// @Accept  json
// @Produce  json
// @Param input body domain.MemberCreateInput true "member info"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/create [post]
func (h *Handler) memberCreateNew(c *gin.Context) {
	var input domain.MemberCreateInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err)
		return
	}

	id, err := h.services.Members.CreateNew(input)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newDataResponse(c, http.StatusOK, dataResponse{
		Message: "Member added successfully",
		Data: map[string]int{
			"id": id,
		},
	})
}

// @Summary Get Member By ID
// @Security ManagerAuth
// @Tags members
// @Description get member by id
// @ModuleID memberGetByID
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/get/{id} [get]
func (h *Handler) memberGetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	member, err := h.services.Members.GetByID(id)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newDataResponse(c, http.StatusOK, dataResponse{
		Message: "Member found",
		Data:    member,
	})
}

// @Summary Update Member By ID
// @Security ManagerAuth
// @Tags members
// @Description update member by id with json body
// @ModuleID memberUpdateByID
// @Accept json
// @Produce  json
// @Param id path int true "Member ID"
// @Param input body domain.MemberUpdateInput true "member update info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/update/{id} [put]
func (h *Handler) memberUpdateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	var input domain.MemberUpdateInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err)
		return
	}

	err = h.services.Members.UpdateByID(id, input)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "Member updated successfully")
}

// @Summary Delete Member By ID
// @Security ManagerAuth
// @Tags members
// @Description delete member by id
// @ModuleID memberDeleteByID
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/delete/{id} [delete]
func (h *Handler) memberDeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	err = h.services.Members.DeleteByID(id)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "Member deleted successfully")
}

// @Summary Set Membership
// @Security ManagerAuth
// @Tags members
// @Description set membership for member
// @ModuleID memberSetMembership
// @Produce  json
// @Param id path int true "Member ID"
// @Param membership_id path int true "Membership ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/set_membership/{id}/{membership_id} [post]
func (h *Handler) memberSetMembership(c *gin.Context) {
	memberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	membershipID, err := strconv.Atoi(c.Param("membership_id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	if err := h.services.Members.SetMembership(memberID, membershipID); err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "Membership set successfully")
}

// @Summary Get Member's Membership
// @Security ManagerAuth
// @Tags members
// @Description get membership from member
// @ModuleID memberGetMembership
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/get_membership/{id} [get]
func (h *Handler) memberGetMembership(c *gin.Context) {
	memberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	res, err := h.services.Members.GetMembership(memberID)
	if err != nil {
		h.handleErrors(c, err)
		return
	}

	newDataResponse(c, http.StatusOK, dataResponse{
		Message: "Membership found",
		Data:    res,
	})
}

// @Summary Delete Member's Membership
// @Security ManagerAuth
// @Tags members
// @Description delete membership from member
// @ModuleID memberDeleteMembership
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/delete_membership/{id} [delete]
func (h *Handler) memberDeleteMembership(c *gin.Context) {
	memberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	if err := h.services.Members.DeleteMembership(memberID); err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}
	newResponse(c, http.StatusOK, "Membership deleted successfully")
}

// @Summary Add Member Visit
// @Security ManagerAuth
// @Tags members
// @Description add member visit
// @ModuleID memberArrived
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/arrived/{id} [post]
func (h *Handler) memberArrived(c *gin.Context) {
	managerID, err := h.getManagerID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	err = h.services.Members.SetNewVisit(id, managerID)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "visit set successfully")
}

// @Summary End Member Visit
// @Security ManagerAuth
// @Tags members
// @Description end member visit
// @ModuleID memberLeft
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/members/left/{id} [post]
func (h *Handler) memberLeft(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	err = h.services.Members.EndVisit(id)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "visit ended successfully")
}
