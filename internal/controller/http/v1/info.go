package v1

import (
	"net/http"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//	@Summary		Get user info
//	@Description	A method for obtaining information about a user using his passport data. Defined in the task conditions.
//	@Security		ApiKeyAuth
//	@Tags			info
//	@ID				ger-user-info
//	@Accept			json
//	@Produces		json
//	@Param			passportSerie	query		int				true	"passportSerie"
//	@Param			passportNumber	query		int				true	"passportNumber"
//	@Success		200				{object}	entity.People	"Ok"
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/v1/info [get]
func (h *Handler) getUserInfo(c *gin.Context) {

	var out entity.People

	serie := c.Query("passportSerie")
	number := c.Query("passportNumber")

	out, err := h.services.GetUserInfo(serie, number)
	if err != nil {
		log.Debugf("error while getting user information: %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, out)
}

//	@Summary		Receiving data about all users with filtering and sorting.
//	@Description	Filtering is possible by the fields "surname", "name", "patronymic", "address". The field(s) for sorting and its direction are specified.
//	@Security		ApiKeyAuth
//	@Tags			info
//	@ID				get-all-user-info
//	@Accept			json
//	@Produces		json
//	@Param			filter			query		string			false	"filtering by fields: surname, name, patronymic, address""
//	@Param			sortProperty	query		string			false	"sorting by fields: surname, name, patronymic, address"
//	@Param			sortDirection	query		string			false	"sorting direction DESC and ASC"
//	@Param			limit			query		int				true	"output limit - maximum value is 10."
//	@Success		200				{object}	[]entity.People	"Ok"
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/v1/info/all [get]
func (h *Handler) getAllUsersInfo(c *gin.Context) {

	filterUsers := c.Query("filter")
	sortProperty := c.Query("sortProperty")
	sortDirection := c.Query("sortDirection")
	limit := c.Query("limit")

	responce, err := h.services.GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limit)
	if err != nil {
		log.Debugf("error while getting all users information: %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, responce)
}

//	@Summary		Obtaining labor costs for a user in a certain period.
//	@Description	Obtaining labor costs by user for a period/ task is the sum of hours and minutes, sorted from the highest cost to the least
//	@Security		ApiKeyAuth
//	@Tags			info
//	@ID				get-user-effort
//	@Accept			json
//	@Produces		json
//	@Param			user_id		query		string						true	"user id"
//	@Param			beginning	query		string						false	"beginning of period in format 2024-07-03"
//	@Param			end			query		string						false	"end of period in format 2024-07-04"
//	@Success		200			{object}	v1.getUserEffort.responce	"Ok"
//	@Failure		400,404		{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Failure		default		{object}	errorResponse
//	@Router			/api/v1/info/effort [get]
func (h *Handler) getUserEffort(c *gin.Context) {

	user_id := c.Query("user_id")
	beginningPeriod := c.Query("beginning")
	endPeriod := c.Query("end")

	if beginningPeriod == "" || endPeriod == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid value for sample period parameters")
		return
	}

	effort, user, err := h.services.GetUserEffort(user_id, beginningPeriod, endPeriod)
	if err != nil {
		log.Debugf("error while getting user effort information: %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type responce struct {
		Description string          `json:"description"`
		Surname     string          `json:"surname"`
		Name        string          `json:"name"`
		Effort      []entity.Effort `json:"effort"`
	}

	fullResponce := responce{
		Description: "The sample was obtained for the period from " + beginningPeriod + " to " + endPeriod + "",
		Surname:     user.Surname,
		Name:        user.Name,
		Effort:      effort,
	}

	c.JSON(http.StatusOK, fullResponce)
}
