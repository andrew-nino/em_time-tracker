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
func (r *AuthPostgres) CreateManager(mng entity.Manager) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, managername, password_hash) values ($1, $2, $3) RETURNING id", managerTable)

	row := r.db.QueryRow(query, mng.Name, mng.Managername, mng.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// We make a request to the database about the user. An error is returned if the result set is empty.
func (r *AuthPostgres) GetManager(managerName, password string) (entity.Manager, error) {
	var user entity.Manager
	query := fmt.Sprintf("SELECT id FROM %s WHERE managername=$1 AND password_hash=$2", managerTable)
	err := r.db.Get(&user, query, managerName, password)

	return user, err
}
