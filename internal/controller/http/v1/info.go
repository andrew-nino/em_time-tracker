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
