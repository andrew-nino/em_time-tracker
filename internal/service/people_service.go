package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
)

type PeopleService struct {
	repo postgresdb.PeopleRepository
}

func NewPeopleService(repo postgresdb.PeopleRepository) *PeopleService {
	return &PeopleService{repo: repo}
}

func (s *PeopleService) CreatePerson(managerID int, passport string) (int, error) {

	serie, number, err := ProcessingPassportData(passport)
	if err != nil {
		return 0, fmt.Errorf("failed to parse passport data: %w", err)
	}
	return s.repo.CreatePerson(managerID, serie, number)
}

func (s *PeopleService) UpdatePerson(passport string, newData entity.People) (int, error) {

	serie, number, err := ProcessingPassportData(passport)
	if err != nil {
		return 0, fmt.Errorf("failed to parse passport data: %w", err)
	}
	return s.repo.UpdatePerson(serie, number, newData)
}

func (s *PeopleService) DeletePerson(managerID int, passport string) error {

	serie, number, err := ProcessingPassportData(passport)
	if err != nil {
		return fmt.Errorf("failed to parse passport data: %w", err)
	}
	return s.repo.DeletePerson(managerID, serie, number)

}

func ProcessingPassportData(passport string) (string, string, error) {

	var isDigit = regexp.MustCompile(`^[0-9]*$`).MatchString

	value := strings.Split(passport, " ")
	if len(value) == 2 {
		for _, v := range value {
			if !isDigit(v) {
				return "", "", fmt.Errorf("invalid string content")
			}
		}
	} else {
		return "", "", fmt.Errorf("invalid passport string")
	}

	serie, err := generatePasswordHash(value[0])
	if err != nil {
		return "", "", err
	}
	number, err := generatePasswordHash(value[1])
	if err != nil {
		return "", "", err
	}
	return serie, number, nil
}
