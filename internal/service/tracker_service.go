package service

import (
	"github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
)

type TrackerService struct {
	repo postgresdb.TrackerRepository
}

func NewTrackerService(repo postgresdb.TrackerRepository) *TrackerService {
	return &TrackerService{repo: repo}
}

func (t *TrackerService) StartTask(user_id, task_id string) (int, error) {
	return t.repo.StartTask(user_id, task_id)
}

func (t *TrackerService) StopTask(user_id, task_id string) error {
	return t.repo.StopTask(user_id, task_id)
}
