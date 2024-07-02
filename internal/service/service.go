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
	CreatePerson(managerID int, passport string) error
	UpdatePerson(passport string, newData entity.People) error
	DeletePerson(managerID int, passport string) error
}

type Info interface {
	GetUserInfo(serie, number string) (entity.People, error)
}

type Service struct {
	Authorization
	People
	Info
}

func NewService(reposPG *postgres.PG_Repository) *Service {
	return &Service{
		Authorization: NewAuthService(reposPG.Authorization),
		People:        NewPeopleService(reposPG.PeopleRepository),
		Info:          NewInfoService(reposPG.InfoRepository),
	}
}
