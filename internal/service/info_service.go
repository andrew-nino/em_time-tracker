package service

import (
	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
)

type InfoService struct {
	repo postgresdb.InfoRepository
}

func NewInfoService(repo postgresdb.InfoRepository) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) GetUserInfo(serie, number string) (p entity.People, err error) {

	serie, err = generatePasswordHash(serie)
	if err != nil {
		return p, err
	}
	number, err = generatePasswordHash(number)
	if err != nil {
		return p, err
	}

	return s.repo.GetUserInfo(serie, number)
}

func (s *InfoService) GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limit string) ([]entity.People, error) {
	return s.repo.GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limit)
}

func (s *InfoService) GetUserEffort(user_id, beginningPeriod, endPeriod string) ([]entity.Effort, entity.People, error) {

	// TODO Добавить проверку валидности строк с периодом выборки данных.
	return s.repo.GetUserEffort(user_id, beginningPeriod, endPeriod)
}
