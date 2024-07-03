package service

import (
	"github.com/andrew-nino/em_time-tracker/entity"
	postgresdb "github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
)

type TasksService struct {
	repo postgresdb.TasksRepository
}

func NewTasksService(repo postgresdb.TasksRepository) *TasksService {
	return &TasksService{repo: repo}
}

func (t *TasksService) CreateTask(task entity.Task) (int, error) {
	return t.repo.CreateTask(task)
}

func (t *TasksService) GetTask(taskID int) (entity.Task, error) {
    return t.repo.GetTask(taskID)
}

func (t *TasksService) GetTasks() ([]entity.Task, error) {
    return t.repo.GetTasks()
}

func (t *TasksService) DeleteTask(taskID int) error {
	return t.repo.DeleteTask(taskID)
}
