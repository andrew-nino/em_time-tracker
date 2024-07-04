package v1

import (
	"net/http"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserInfo(c *gin.Context) {

	var out entity.People

	serie := c.Query("passportSerie")
	number := c.Query("passportNumber")

	out, err := h.services.GetUserInfo(serie, number)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, out)
}

func (h *Handler) getAllUsersInfo(c *gin.Context) {

	filterUsers := c.Query("filter")
	sortProperty := c.Query("sortProperty")
	sortDirection := c.Query("sortDirection")
	limit := c.Query("limit")

	responce, err := h.services.GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, responce)
}

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
