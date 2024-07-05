package v1

import (
	"net/http"
	"strconv"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//	@Summary		Create task
//	@Description	Creates a new task using the parameters and returns its ID on success. By default, a new task has an internal status of "planned"
//	@Security		ApiKeyAuth
//	@Tags			tasks
//	@ID				create-task
//	@Accept			json
//	@Produces		json
//	@Param			description	of			the	task	body	entity.Task	true	"You need to give the task a name, its importance (high or low(default)) and a description."
//	@Success		200			{object}	v1.response
//	@Failure		400,404		{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Failure		default		{object}	errorResponse
//	@Router			/api/v1/tasks/create [post]
func (h *Handler) createTask(c *gin.Context) {

	var input entity.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	taskID, err := h.services.CreateTask(input)
	if err != nil {
		log.Debugf("error when creating task  : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "task creation failed")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: taskID})
}

//	@Summary		Retrieving a task by its ID
//	@Description	We get the task by its ID in the request parameters.
//	@Security		ApiKeyAuth
//	@Tags			tasks
//	@ID				get-task
//	@Accept			json
//	@Produces		json
//	@Param			taskId	query		string	true	"Yes"
//	@Success		200		{object}	v1.response
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/v1/tasks/id [get]
func (h *Handler) getTask(c *gin.Context) {

	var out entity.Task

	taskID := c.Query("taskId")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to convert task ID")
		return
	}

	out, err = h.services.GetTask(id)
	if err != nil {
		log.Debugf("error when receiving task  : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "failed to get task")
		return
	}
	c.JSON(http.StatusOK, out)
}

//	@Summary		Get all tasks
//	@Description	We receive all the tasks and set a limit on page output.
//	@Security		ApiKeyAuth
//	@Tags			tasks
//	@ID				get-all-task
//	@Accept			json
//	@Produces		json
//	@Param			limit	query		string	true	"limit on page outpu"
//	@Success		200		{object}	[]entity.Task
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/v1/tasks/all [get]
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
		log.Debugf("error when receiving tasks : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "failed to get tasks")
		return
	}
	c.JSON(http.StatusOK, out)
}

//	@Summary		Deleting a task
//	@Description	Deleting a task using the ID obtained from the parameter.
//	@Security		ApiKeyAuth
//	@Tags			tasks
//	@ID				delete-task
//	@Accept			json
//	@Produces		json
//	@Param			taskID	query		string	true	"task ID to delete"
//	@Success		200		{object}	v1.response
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/v1/tasks/delete [delete]
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
		log.Debugf("error when deleting tasks : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete task")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success"})
}
