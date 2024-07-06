package service

import (
	"fmt"
	"regexp"

	log "github.com/sirupsen/logrus"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
)

type InfoService struct {
	repo postgresdb.InfoRepository
}

func NewInfoService(repo postgresdb.InfoRepository) *InfoService {
	return &InfoService{repo: repo}
}

// Checking the validity of the incoming string, processing the data into a hash and transmitting it to the database
func (s *InfoService) GetUserInfo(serie, number string) (ppl entity.People, err error) {

	var isDigit = regexp.MustCompile(`^[0-9]*$`).MatchString
	if !isDigit(serie) {
		return ppl, fmt.Errorf("invalid serie value")
	}
	if !isDigit(number) {
		return ppl, fmt.Errorf("invalid number value")
	}

	serie, err = generatePasswordHash(serie)
	if err != nil {
		return ppl, err
	}
	number, err = generatePasswordHash(number)
	if err != nil {
		return ppl, err
	}

	ppl, err = s.repo.GetUserInfo(serie, number)
	if err != nil {
		log.Errorf("InfoService.GetUserInfo - s.repo.GetUserInfo: %v", err)
		return ppl, err
	}

	return ppl, nil
}

func (s *InfoService) GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limit string) ([]entity.People, error) {
	return s.repo.GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limit)
}

func (s *InfoService) GetUserEffort(user_id, beginningPeriod, endPeriod string) ([]entity.Effort, entity.People, error) {

	// TODO Добавить проверку валидности строк с периодом выборки данных.
	return s.repo.GetUserEffort(user_id, beginningPeriod, endPeriod)
}
