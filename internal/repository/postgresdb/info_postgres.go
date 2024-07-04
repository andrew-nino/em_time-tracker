package postgresdb

import (
	"fmt"
	"strconv"

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

var offsetAllInfo int

func (i *InfoFromPostgres) GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limitStr string) ([]entity.People, error) {
	limit, _ := strconv.Atoi(limitStr)
	var responce []entity.People

	if filterUsers == "" {
		filterUsers = "surname, name, patronymic, address"
	}
	if sortProperty == "" {
		sortProperty = filterUsers
	}
	if sortDirection == "" {
		sortDirection = "ASC"
	}

backward:
	queryPeopleStr := fmt.Sprintf("SELECT %s FROM %s WHERE id > $1 ORDER BY %s %s LIMIT $2", filterUsers, peopleTable, sortProperty, sortDirection)
	err := i.db.Select(&responce, queryPeopleStr, offsetAllInfo, limit)
	if err != nil {
		return nil, err
	}

	if responce == nil {
		offsetAllInfo = 0
		goto backward
	} else if len(responce) < limit {
		offsetAllInfo = 0
	} else {
		offsetAllInfo += limit
	}

	return responce, nil
}
