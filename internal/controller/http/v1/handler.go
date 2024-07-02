package v1

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/andrew-nino/em_time-tracker/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	gin.DisableConsoleColor()
	f := setLogsFile()
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	v1 := router.Group("/api/v1")
	{
		info := v1.Group("/info")
		{
			info.GET("/", h.getUserInfo)
		}

		people := v1.Group("/people")
		{
			people.POST("/create", h.createPerson)
			people.PATCH("/update", h.updatePerson)
			people.DELETE("/delete", h.deletePerson)
		}
	}

	return router
}

func setLogsFile() *os.File {
	file, err := os.OpenFile("./logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
