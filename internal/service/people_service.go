package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
	log "github.com/sirupsen/logrus"
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

	id, err := s.repo.CreatePerson(managerID, serie, number)
	if err != nil {
		log.Errorf("PeopleService.CreatePerson - s.repo.CreatePerson: %v", err)
		return 0, err
	}
	return id, nil
}

func (s *PeopleService) UpdatePerson(passport string, newData entity.People) (int, error) {

	serie, number, err := ProcessingPassportData(passport)
	if err != nil {
		return 0, fmt.Errorf("failed to parse passport data: %w", err)
	}

	id, err := s.repo.UpdatePerson(serie, number, newData)
	if err != nil {
		log.Errorf("PeopleService.UpdatePerson - s.repo.UpdatePerson: %v", err)
		return 0, err
	}
	return id, nil
}

func (s *PeopleService) DeletePerson(managerID int, passport string) error {

	serie, number, err := ProcessingPassportData(passport)
	if err != nil {
		return fmt.Errorf("failed to parse passport data: %w", err)
	}

	err = s.repo.DeletePerson(managerID, serie, number)
	if err != nil {
		log.Errorf("PeopleService.DeletePerson - s.repo.DeletePerson: %v", err)
		return err
	}
	return nil
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
