package postgresdb

import (
	"fmt"
	"strings"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PeopleToPostgres struct {
	db *sqlx.DB
}

func NewPeopleToPostgres(db *sqlx.DB) *PeopleToPostgres {
	return &PeopleToPostgres{db: db}
}

func (p *PeopleToPostgres) CreatePerson(managerID int, serie, number string) (int, error) {

	var id int
	query := fmt.Sprintf("INSERT INTO %s (manager_id, passport_serie, passport_number) values ($1, $2, $3) RETURNING id", peopleTable)
	row := p.db.QueryRow(query, managerID, serie, number)
	err := row.Scan(&id)
	if err != nil {
		log.Debugf("repository.CreatePerson - row.Scan : %v", err)
		return 0, err
	}
	return id, nil
}

func (p *PeopleToPostgres) UpdatePerson(serie, number string, newData entity.People) (int, error) {

	var id int
	setValues := make([]string, 0)

	if newData.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname='%s'", newData.Surname))
	}

	if newData.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name='%s'", newData.Name))
	}

	if newData.Patronymic != "" {
		setValues = append(setValues, fmt.Sprintf("patronymic='%s'", newData.Patronymic))
	}

	if newData.Address != "" {
		setValues = append(setValues, fmt.Sprintf("address='%s'", newData.Address))
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s  WHERE passport_serie = $1 AND passport_number = $2 RETURNING id`, peopleTable, setQuery)

	row := p.db.QueryRow(query, serie, number)
	err := row.Scan(&id)
	if err != nil {
		log.Debugf("repository.UpdatePerson - row.Scan : %v", err)
		return 0, err
	}

	return id, err
}

func (p *PeopleToPostgres) DeletePerson(managerID int, serie, number string) error {

	query := fmt.Sprintf(`DELETE FROM %s WHERE passport_serie = $1 AND passport_number = $2`, peopleTable)
	_, err := p.db.Exec(query, serie, number)
	if err != nil {
		log.Debugf("repository.DeletePerson - db.Exec : %v", err)
		return err
	}
	return nil
}
