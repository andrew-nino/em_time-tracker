package v1

import (
	"net/http"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//	@Summary		SignUp
//	@Tags			auth
//	@Description	create account
//	@ID				create_account
//	@Accept			json
//	@Produce		json
//	@Param			input	body		entity.Manager	true	"account info"
//	@Success		200		{integer}	integer			1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {

	var input entity.Manager

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateManager(input)
	if err != nil {
		log.Debugf("error when registering manager : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	ManagerName string `json:"managername" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

//	@Summary		SignIn
//	@Tags			auth
//	@Description	login
//	@ID				login
//	@Accept			json
//	@Produce		json
//	@Param			input	body		signInInput	true	"credentials"
//	@Success		200		{string}	string		"token"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.SignIn(input.ManagerName, input.Password)
	if err != nil {
		log.Debugf("error during mamager verification : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
