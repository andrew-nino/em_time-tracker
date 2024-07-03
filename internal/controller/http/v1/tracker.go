package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) startTask(c *gin.Context) {

	user_id := c.Query("people_id")
	task_id := c.Query("task_id")

	if user_id == "" || task_id == "" {
		newErrorResponse(c, http.StatusBadRequest, "people_id and task_id are required")
		return
	}

	tracker_id, err := h.services.StartTask(user_id, task_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: tracker_id})
}

func (h *Handler) stopTask(c *gin.Context) {

	user_id := c.Query("people_id")
	task_id := c.Query("task_id")

	if user_id == "" || task_id == "" {
		newErrorResponse(c, http.StatusBadRequest, "people_id and task_id are required")
		return
	}

	if err := h.services.StopTask(user_id, task_id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success"})
}
