package v1

import (
	"errors"
	"net/http"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type userInput struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type response struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

//	@Summary		Create person
//	@Description	The format is specified in the task conditions. Creates a new user using passport data. Returns the ID of the created user. To further fill/change data, the Update method works.
//	@Security		ApiKeyAuth
//	@Tags			people
//	@ID				create-person
//	@Accept			json
//	@Produces		json
//	@Param			passportNumber	body		v1.userInput	true	"passportNumber"
//	@Success		200				{object}	v1.response
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/v1/people/create [post]
func (h *Handler) createPerson(c *gin.Context) {

	var input userInput
	// TODO снять коментанрий и отправлять значение из контекста
	var stub = 1

	// managerId, err := getManagerId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreatePerson(stub, input.PassportNumber)
	if err != nil {
		if errors.Is(err, ErrCannotParsePassport) {
			newErrorResponse(c, http.StatusBadRequest, "Invalid passport number")
			return
		}
		log.Debugf("error when create person : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: id})
}

//	@Summary		Updates information about a person
//	@Description	Updates the user data to the specified one. user verification is carried out using the passport data in the body of the request. If successful, returns a message and user ID.
//	@Security		ApiKeyAuth
//	@Tags			people
//	@ID				update-person
//	@Accept			json
//	@Produces		json
//	@Param			passportNumber	body		v1.userInput	true	"passportNumber"
//	@Param			surname			query		string			false	"surname user"
//	@Param			name			query		string			false	"name user"
//	@Param			patronymic		query		string			false	"patronymic user"
//	@Param			address			query		string			false	"address user"
//	@Success		200				{object}	v1.response
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/v1/people/update [patch]
func (h *Handler) updatePerson(c *gin.Context) {

	var input userInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	surname := c.Query("surname")
	name := c.Query("name")
	patronymic := c.Query("patronymic")
	address := c.Query("address")

	newData := entity.People{
		Surname:    surname,
		Name:       name,
		Patronymic: patronymic,
		Address:    address,
	}

	id, err := h.services.UpdatePerson(input.PassportNumber, newData)
	if err != nil {
		log.Debugf("error when update person : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: id})
}

//	@Summary		Delete information about a person
//	@Description	Based on accepted passport data, removes the user from the system.
//	@Security		ApiKeyAuth
//	@Tags			people
//	@ID				delete-person
//	@Accept			json
//	@Produces		json
//	@Param			passportNumber	body		v1.userInput	true	"passportNumber"
//	@Success		200				{object}	v1.response
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/v1/people/delete [delete]
func (h *Handler) deletePerson(c *gin.Context) {

	var input userInput
	// TODO снять коментанрий и отправлять значение из контекста
	var stub = 1

	// managerId, err := getManagerId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeletePerson(stub, input.PassportNumber); err != nil {
		log.Debugf("error when delete person : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}
