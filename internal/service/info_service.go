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
