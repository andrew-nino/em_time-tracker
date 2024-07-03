package v1

import (
	"errors"
	"net/http"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/gin-gonic/gin"
)

type userInput struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type response struct {
	Message string `json:"message"`
	ID  int    `json:"id"`
}

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

	if err := h.services.CreatePerson(stub, input.PassportNumber); err != nil {
		if errors.Is(err, ErrCannotParsePassport) {
			newErrorResponse(c, http.StatusBadRequest, "Invalid passport number")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success"})
}

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

	if err := h.services.UpdatePerson(input.PassportNumber, newData); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success"})
}

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
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}
