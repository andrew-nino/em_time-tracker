package v1

import (
	"io"
	"log"
	"os"

	"github.com/andrew-nino/em_time-tracker/internal/service"

	_ "github.com/andrew-nino/em_time-tracker/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// v1 := router.Group("/api/v1", h.userIdentity)
	v1 := router.Group("/api/v1")

	{
		info := v1.Group("/info")
		{
			info.GET("/", h.getUserInfo)
			info.GET("/all", h.getAllUsersInfo)
			info.GET("/effort", h.getUserEffort)
		}

		people := v1.Group("/people")
		{
			people.POST("/create", h.createPerson)
			people.PATCH("/update", h.updatePerson)
			people.DELETE("/delete", h.deletePerson)
		}

		tasks := v1.Group("/tasks")
		{
			tasks.POST("/create", h.createTask)
			tasks.GET("/id", h.getTask)
			tasks.GET("/all", h.getTasks)
			tasks.DELETE("/delete", h.deleteTask)
		}
		tracker := v1.Group("/tracker")
		{
			tracker.POST("/start", h.startTracker)
			tracker.POST("/stop", h.stopTracker)
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
