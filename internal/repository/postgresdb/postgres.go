package postgresdb

import (
	"github.com/jmoiron/sqlx"

	"github.com/andrew-nino/em_time-tracker/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type PG_Repository struct {
	Authorization
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		Authorization: NewAuthPostgres(db),
	}
}
