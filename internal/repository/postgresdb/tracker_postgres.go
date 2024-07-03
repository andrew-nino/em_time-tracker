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
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var track_id int
	var statusTask string

	checkQuery := fmt.Sprintf("SELECT status FROM %s WHERE id = ($1)", taskTable)
	statusRow := tx.QueryRow(checkQuery, task_id)
	err = statusRow.Scan(&statusTask)
	if err != nil {
		tx.Rollback()
		return 0, err
	} else if statusTask != "planed" {
		tx.Rollback()
		return 0, fmt.Errorf("status task is not a planed status")
	}

	createItemQuery := fmt.Sprintf("INSERT INTO %s (task_id, people_id) values ($1, $2) RETURNING id", trackerTable)
	row := tx.QueryRow(createItemQuery, task_id, user_id)
	err = row.Scan(&track_id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to start task: %w", err)
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET status = 'accepted' WHERE id = $1", taskTable)
	_, err = tx.Exec(updateQuery, task_id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to update task status when running task: %w", err)
	}

	return track_id, tx.Commit()
}

func (t *TrackerToPostgres) StopTask(user_id, task_id string) error {

	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	createItemQuery := fmt.Sprintf("UPDATE %s SET finished_at = now() WHERE task_id = $1 AND people_id = $2", trackerTable)
	_, err = tx.Exec(createItemQuery, task_id, user_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating tracker: %w", err)
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET status = 'completed' WHERE id = $1", taskTable)
	_, err = tx.Exec(updateQuery, task_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update task status when running task: %w", err)
	}
	return tx.Commit()
}
