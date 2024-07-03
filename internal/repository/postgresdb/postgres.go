package postgresdb

import (
	"github.com/jmoiron/sqlx"

	"github.com/andrew-nino/em_time-tracker/entity"
)

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	GetManager(username, password string) (entity.Manager, error)
}

type PeopleRepository interface {
	CreatePerson(managerID int, serie, number string) error
	UpdatePerson(serie, number string, newData entity.People) error
	DeletePerson(managerID int, serie, number string) error
}

type InfoRepository interface {
	GetUserInfo(serie, number string) (entity.People, error)
}

type TasksRepository interface {
	CreateTask(task entity.Task) (int, error)
	GetTask(taskId int) (entity.Task, error)
	GetTasks(limit int) ([]entity.Task, error)
	DeleteTask(taskId int) error
}

type TrackerRepository interface {
	StartTask(user_id, task_id string) (int, error)
	StopTask(user_id, task_id string) error
}

type PG_Repository struct {
	Authorization
	PeopleRepository
	InfoRepository
	TasksRepository
	TrackerRepository
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		Authorization:     NewAuthPostgres(db),
		PeopleRepository:  NewPeopleToPostgres(db),
		InfoRepository:    NewInfoFromPostgres(db),
		TasksRepository:   NewTasksToPostgres(db),
		TrackerRepository: NewTrackerPostgres(db),
	}
}
