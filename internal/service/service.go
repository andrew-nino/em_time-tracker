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

type Service struct {
	Authorization
}

func NewService(reposPG *postgres.PG_Repository) *Service {
	return &Service{
		Authorization: NewAuthService(reposPG.Authorization),
	}
}
