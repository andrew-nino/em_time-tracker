package postgresdb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
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
		log.Debugf("repository.GetUserInfo - db.Get : %v", err)
		return response, err
	}
	return response, nil
}

var g_offsetAllInfo int
// Filters by fields and sorting with a given direction are used.
func (i *InfoFromPostgres) GetAllUsersInfo(filterUsers, sortProperty, sortDirection, limitStr string) ([]entity.People, error) {

	var response []entity.People

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Debugf("repository.GetAllUsersInfo - strconv.Atoi : %v", err)
		return response, err
	}
// Check and set default limit and reset global variable
	if limit > 10 || limit <= 0 {
		g_offsetAllInfo = 0
		limit = 10
	}
// Set the default limit for filtering and sorting
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
	err = i.db.Select(&response, queryPeopleStr, g_offsetAllInfo, limit)
	if err != nil {
		log.Debugf("repository.GetAllUsersInfo - db.Select : %v", err)
		return nil, err
	}

	if response == nil {
		g_offsetAllInfo = 0
		goto backward
	} else if len(response) <= limit {
		g_offsetAllInfo = 0
	} else {
		g_offsetAllInfo += limit
	}

	return response, nil
}
// Obtaining data on user labor costs for a specified period, sorted in descending order.
func (i *InfoFromPostgres) GetUserEffort(user_id, beginningPeriod, endPeriod string) ([]entity.Effort, entity.People, error) {

	// TODO Сделать проверку входящих данных

	effor := []entity.Effort{}
	queryPeopleStr := fmt.Sprintf(`SELECT trc.task_id, tsc.description, 
								   EXTRACT( EPOCH FROM trc.finished_at - trc.created_at) AS total_time 
								   FROM %s AS trc
								   INNER JOIN %s  AS tsc ON tsc.id = trc.task_id
								   WHERE trc.created_at >= to_timestamp($1, 'YYYY-MM-DD') AND trc.finished_at <= to_timestamp($2, 'YYYY-MM-DD') + interval '24 hour'
								   GROUP BY trc.people_id, trc.task_id, total_time, tsc.description
								   HAVING trc.people_id = $3
								   ORDER BY trc.people_id, total_time DESC`, trackerTable, taskTable)
	err := i.db.Select(&effor, queryPeopleStr, beginningPeriod, endPeriod, user_id)
	if err != nil {
		log.Debugf("repository.GetUserEffort - db.Select : %v", err)
		return nil, entity.People{}, err
	}

	for i := 0; i < len(effor); i++ {
		effor[i].TotalTime, _, _ = strings.Cut(effor[i].TotalTime, ".")
		h, err := strconv.Atoi(effor[i].TotalTime)
		if err != nil {
			log.Debugf("repository.GetUserEffort - strconv.Atoi : %v", err)
			return nil, entity.People{}, err
		}

		effor[i].TotalTime = fmt.Sprintf("%dh %dm", h/3600, h%3600/60)
	}

	var user entity.People
	query := fmt.Sprintf("SELECT surname, name FROM %s WHERE id = $1", peopleTable)
	err = i.db.Get(&user, query, user_id)
	if err != nil {
		log.Debugf("repository.GetUserEffort - db.Get: %v", err)
		return nil, entity.People{}, err
	}

	return effor, user, nil
}
