package postgresdb

import (
	"fmt"
	"strings"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/jmoiron/sqlx"
)

type PeopleToPostgres struct {
	db *sqlx.DB
}

func NewPeopleToPostgres(db *sqlx.DB) *PeopleToPostgres {
	return &PeopleToPostgres{db: db}
}

func (p *PeopleToPostgres) CreatePerson(managerID int, serie, number string) error {

	query := fmt.Sprintf("INSERT INTO %s (manager_id, passport_serie, passport_number) values ($1, $2, $3)", peopleTable)

	_, err := p.db.Exec(query, managerID, serie, number)
	if err != nil {
		return err
	}

	return nil
}

func (p *PeopleToPostgres) UpdatePerson(serie, number string, newData entity.People) error {

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

	query := fmt.Sprintf(`UPDATE %s SET %s  WHERE passport_serie = $1 AND passport_number = $2`, peopleTable, setQuery)

	_, err := p.db.Exec(query, serie, number)

	return err
}

func (p *PeopleToPostgres) DeletePerson(managerID int, serie, number string) error {

	query := fmt.Sprintf(`DELETE FROM %s WHERE passport_serie = $1 AND passport_number = $2`, peopleTable)

	_, err := p.db.Exec(query, serie, number)

	return err
}
