package postgresdb

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/andrew-nino/em_time-tracker/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// We create a new user in the database and return his ID or the error [ErrNoRows] if it does not work.
func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (passportSerie, passportNumber) values ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.PassportSerie, user.PassportNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// We make a request to the database about the user. An error is returned if the result set is empty.
func (r *AuthPostgres) GetUser(serie, number string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE passportSerie=$1 AND passportNumber=$2", usersTable)
	err := r.db.Get(&user, query, serie, number)

	return user, err
}
