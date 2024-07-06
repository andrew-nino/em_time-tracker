package postgresdb

import (
	"fmt"

	"github.com/andrew-nino/em_time-tracker/entity"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TaskToPostgres struct {
	db *sqlx.DB
}

func NewTasksToPostgres(db *sqlx.DB) *TaskToPostgres {
	return &TaskToPostgres{db: db}
}

func (t *TaskToPostgres) CreateTask(task entity.Task) (int, error) {

	var taskID int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (name, importance, description) VALUES ($1, $2, $3) RETURNING id", taskTable)
	row := t.db.QueryRow(createTaskQuery, task.Name, task.Importance, task.Description)
	err := row.Scan(&taskID)
	if err != nil {
		log.Debugf("repository.CreateTask - row.Scan : %v", err)
		return 0, err
	}
	return taskID, nil
}

func (t *TaskToPostgres) GetTask(taskId int) (entity.Task, error) {

	var task entity.Task
	query := fmt.Sprintf("SELECT name, importance, description FROM %s WHERE id = $1", taskTable)
	err := t.db.Get(&task, query, taskId)
	if err != nil {
		log.Debugf("repository.GetTask - db.Get : %v", err)
		return task, err
	}
	return task, nil
}

var g_offsetGetTasks int

func (t *TaskToPostgres) GetTasks(limit int) ([]entity.Task, error) {

	var tasks []entity.Task

	// Check and set default limit and reset global variable
	if limit > 10 || limit <= 0 {
		g_offsetGetTasks = 0
		limit = 10
	}
backward:
	query := fmt.Sprintf("SELECT name, importance, description FROM %s WHERE id > $1 ORDER BY id ASC LIMIT $2", taskTable)
	err := t.db.Select(&tasks, query, g_offsetGetTasks, limit)
	if err != nil {
		log.Debugf("repository.GetTasks - db.Select : %v", err)
		return tasks, err
	}

	if tasks == nil {
		g_offsetGetTasks = 0
		goto backward
	} else if len(tasks) < limit {
		g_offsetGetTasks = 0
	} else {
		g_offsetGetTasks += limit
	}
	return tasks, nil
}

func (t *TaskToPostgres) DeleteTask(taskId int) error {

	query := fmt.Sprintf("DELETE FROM %s  WHERE id = $1", taskTable)
	_, err := t.db.Exec(query, taskId)
	if err != nil {
		log.Debugf("repository.DeleteTask - db.Exec : %v", err)
		return err
	}
	return nil
}
