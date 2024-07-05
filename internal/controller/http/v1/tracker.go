package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//	@Summary		Starts timing on a task for the user.
//	@Description	Starts counting down the task completion time for the user. Inside, the task status changes from "planned" to "accepted". If successful, we receive the tracker ID.
//	@Security		ApiKeyAuth
//	@Tags			tracker
//	@ID				start-tracker
//	@Accept			json
//	@Produces		json
//	@Param			task_id		query		string	true	"specify task ID"
//	@Param			people_id	query		string	true	"specify user ID"
//	@Success		200			{object}	v1.response
//	@Failure		400,404		{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Failure		default		{object}	errorResponse
//	@Router			/api/v1/tracker/start [post]
func (h *Handler) startTracker(c *gin.Context) {

	task_id := c.Query("task_id")
	user_id := c.Query("people_id")

	if user_id == "" || task_id == "" {
		newErrorResponse(c, http.StatusBadRequest, "people_id and task_id are required")
		return
	}

	tracker_id, err := h.services.StartTracker(user_id, task_id)
	if err != nil {
		log.Debugf("error when start tracker : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: tracker_id})
}

//	@Summary		Stops the execution time of a task for the user.
//	@Description	Ends the task timer for the user. Inside, the task status changes from “accepted” to “completed”.
//	@Security		ApiKeyAuth
//	@Tags			tracker
//	@ID				stop-tracker
//	@Accept			json
//	@Produces		json
//	@Param			task_id		query		string	true	"specify task ID"
//	@Param			people_id	query		string	true	"specify user ID"
//	@Success		200			{object}	v1.response
//	@Failure		400,404		{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Failure		default		{object}	errorResponse
//	@Router			/api/v1/tracker/stop [post]
func (h *Handler) stopTracker(c *gin.Context) {

	task_id := c.Query("task_id")
	user_id := c.Query("people_id")

	if user_id == "" || task_id == "" {
		newErrorResponse(c, http.StatusBadRequest, "people_id and task_id are required")
		return
	}

	if err := h.services.StopTracker(user_id, task_id); err != nil {
		log.Debugf("error when stop tracker : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Message: "success"})
}
