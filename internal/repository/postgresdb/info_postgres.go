package postgresdb

import (
	"fmt"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/jmoiron/sqlx"
)

type InfoFromPostgres struct {
	db *sqlx.DB
}

func NewInfoFromPostgres(db *sqlx.DB) *InfoFromPostgres {
	return &InfoFromPostgres{db: db}
}

func (i *InfoFromPostgres) GetUserInfo(serie, number string) (entity.People, error) {

	var response entity.People

	query := fmt.Sprintf("SELECT surname, name, patronymic, address FROM %s WHERE passport_serie = $1 AND passport_number = $2", peopleTable)

	err := i.db.Get(&response, query, serie, number)

	if err != nil {
		return response, err
	}
	return response, nil
}
