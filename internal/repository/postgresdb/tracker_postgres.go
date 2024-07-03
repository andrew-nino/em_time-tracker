package postgresdb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TrackerToPostgres struct {
	db *sqlx.DB
}

func NewTrackerPostgres(db *sqlx.DB) *TrackerToPostgres {
	return &TrackerToPostgres{db: db}
}

func (t *TrackerToPostgres) StartTask(user_id, task_id string) (int, error) {

	var track_id int

	createItemQuery := fmt.Sprintf("INSERT INTO %s (task_id, people_id) values ($1, $2) RETURNING id", trackerTable)
	row := t.db.QueryRow(createItemQuery, task_id, user_id)
	err := row.Scan(&track_id)
	if err != nil {
		return 0, err
	}

	return track_id, nil
}

func (t *TrackerToPostgres) StopTask(user_id, task_id string) error {

	createItemQuery := fmt.Sprintf("UPDATE %s SET finished_at = now() WHERE task_id = $1 AND people_id = $2", trackerTable)
	_, err := t.db.Exec(createItemQuery, task_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
