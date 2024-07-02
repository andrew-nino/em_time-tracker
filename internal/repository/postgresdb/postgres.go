package postgresdb

import (
	"github.com/jmoiron/sqlx"

	"github.com/andrew-nino/em_time-tracker/entity"
)

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	GetManager(username, password string) (entity.Manager, error)
}

type PeopleRepository interface {
	CreatePerson(managerID int, serie, number string) error
	UpdatePerson(serie, number string, newData entity.People) error
	DeletePerson(managerID int, serie, number string) error
}

type InfoRepository interface {
	GetUserInfo(serie, number string) (entity.People, error)
}

type PG_Repository struct {
	Authorization
	PeopleRepository
	InfoRepository
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		Authorization: NewAuthPostgres(db),
		PeopleRepository: NewPeopleToPostgres(db),
		InfoRepository: NewInfoFromPostgres(db),
	}
}
