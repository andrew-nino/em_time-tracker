package v1

import (
	"net/http"
	"strconv"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {

	var input entity.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	taskID, err := h.services.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "task creation failed")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: taskID})
}

func (h *Handler) getTask(c *gin.Context) {

	var out entity.Task

	taskID := c.Param("taskId")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to convert task ID")
		return
	}

	out, err = h.services.GetTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get task")
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *Handler) getTasks(c *gin.Context) {

	var out []entity.Task
	var limit int

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to convert task ID")
		return
	}

	out, err = h.services.GetTasks(limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get tasks")
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *Handler) deleteTask(c *gin.Context) {

	taskID := c.Query("taskID")
	if taskID == "" {
		newErrorResponse(c, http.StatusBadRequest, "task ID is required")
		return
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to convert task ID")
		return
	}

	err = h.services.DeleteTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete task")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success"})
}
