package v1

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidAuthHeader   = fmt.Errorf("invalid auth header")
	// ErrCannotParseToken    = fmt.Errorf("cannot parse token")
	ErrCannotParsePassport = fmt.Errorf("invalid passport string")
	// ErrCannotParseToken  = fmt.Errorf("cannot parse token")

)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
