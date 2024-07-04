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

func (i *InfoFromPostgres) GetUserEffort(user_id, beginningPeriod, endPeriod string) ([]entity.Effort, entity.People, error) {

	effor := []entity.Effort{}
	queryPeopleStr := fmt.Sprintf(`SELECT task_id, to_char( finished_at - created_at, 'HH24:MI') AS total_time 
								   FROM %s
								   WHERE created_at >= to_timestamp($1, 'YYYY-MM-DD') AND finished_at <= to_timestamp($2, 'YYYY-MM-DD') + interval '24 hour'
								   GROUP BY people_id, task_id, total_time
								   HAVING people_id = $3
								   ORDER BY people_id, total_time DESC`, trackerTable)
	err := i.db.Select(&effor, queryPeopleStr, beginningPeriod, endPeriod, user_id)
	if err != nil {
		return nil, entity.People{}, err
	}

	var user entity.People
	query := fmt.Sprintf("SELECT surname, name FROM %s WHERE id = $1", peopleTable)
	err = i.db.Get(&user, query, user_id)
	if err != nil {
		return nil, entity.People{}, err
	}

	return effor, user, nil
}
