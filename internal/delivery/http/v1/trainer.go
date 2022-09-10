package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"net/http"
	"strconv"
)

// @Summary Create trainer
// @Security ManagerAuth
// @Tags trainers
// @Description trainer creation
// @ModuleID trainerCreateNew
// @Accept  json
// @Produce  json
// @Param input body domain.TrainerCreateInput true "trainer info"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/trainers/create [post]
func (h *Handler) trainerCreateNew(c *gin.Context) {
	var input domain.TrainerCreateInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err)
		return
	}

	id, err := h.services.Trainers.CreateNew(input)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newDataResponse(c, http.StatusOK, dataResponse{
		Message: "trainer added successfully",
		Data: map[string]int{
			"id": id,
		},
	})
}

// @Summary Get trainer By ID
// @Security ManagerAuth
// @Tags trainers
// @Description get trainer by id
// @ModuleID trainerGetByID
// @Produce  json
// @Param id path int true "trainer ID"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/trainers/get/{id} [get]
func (h *Handler) trainerGetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	trainer, err := h.services.Trainers.GetByID(id)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newDataResponse(c, http.StatusOK, dataResponse{
		Message: "trainer found",
		Data:    trainer,
	})
}

// @Summary Update Trainer By ID
// @Security ManagerAuth
// @Tags trainers
// @Description update trainer by id with json body
// @ModuleID trainerUpdateByID
// @Accept json
// @Produce  json
// @Param id path int true "Trainer ID"
// @Param input body domain.TrainerUpdateInput true "trainer update info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/trainers/update/{id} [put]
func (h *Handler) trainerUpdateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	var input domain.TrainerUpdateInput
	if err := c.BindJSON(&input); err != nil {
		h.handleErrors(c, err)
		return
	}

	err = h.services.Trainers.UpdateByID(id, input)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "trainer updated successfully")
}

// @Summary Delete Trainer By ID
// @Security ManagerAuth
// @Tags trainers
// @Description delete trainer by id
// @ModuleID trainerDeleteByID
// @Produce  json
// @Param id path int true "Trainer ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/trainers/delete/{id} [delete]
func (h *Handler) trainerDeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	err = h.services.Trainers.DeleteByID(id)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "trainer deleted successfully")
}

// @Summary Add Trainer Visit
// @Security ManagerAuth
// @Tags trainers
// @Description add trainer visit
// @ModuleID trainerArrived
// @Produce  json
// @Param id path int true "Trainer ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/trainers/arrived/{id} [post]
func (h *Handler) trainerArrived(c *gin.Context) {
	managerID, err := h.getManagerID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	err = h.services.Trainers.SetNewVisit(id, managerID)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "visit set successfully")
}

// @Summary End Trainer Visit
// @Security ManagerAuth
// @Tags trainers
// @Description end trainer visit
// @ModuleID trainerLeft
// @Produce  json
// @Param id path int true "Trainer ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router       /managers/trainers/left/{id} [post]
func (h *Handler) trainerLeft(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrors(c, errors.New(errBadRequest))
		return
	}

	err = h.services.Trainers.EndVisit(id)
	if err != nil {
		h.handleErrors(c, err)
		return
	}
	newResponse(c, http.StatusOK, "visit ended successfully")
}
