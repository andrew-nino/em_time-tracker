package service

import (
	"github.com/andrew-nino/em_time-tracker/entity"
	postgres "github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
)

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	SignIn(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type People interface {
	CreatePerson(managerID int, passport string) (int, error)
	UpdatePerson(passport string, newData entity.People) (int, error)
	DeletePerson(managerID int, passport string) error
}

type Info interface {
	GetUserInfo(serie, number string) (entity.People, error)
}

type Tasks interface {
	CreateTask(task entity.Task) (int, error)
	DeleteTask(taskId int) error
	GetTask(taskId int) (entity.Task, error)
	GetTasks(limit int) ([]entity.Task, error)
}

type Tracker interface {
	StartTask(user_id, task_id string) (int, error)
	StopTask(user_id, task_id string) error
}
type Service struct {
	Authorization
	People
	Info
	Tasks
	Tracker
}

func NewService(reposPG *postgres.PG_Repository) *Service {
	return &Service{
		Authorization: NewAuthService(reposPG.Authorization),
		People:        NewPeopleService(reposPG.PeopleRepository),
		Info:          NewInfoService(reposPG.InfoRepository),
		Tasks:         NewTasksService(reposPG.TasksRepository),
		Tracker:       NewTrackerService(reposPG.TrackerRepository),
	}
}
