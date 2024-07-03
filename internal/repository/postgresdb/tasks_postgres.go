package postgresdb

import (
	"fmt"

	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/jmoiron/sqlx"
)

type TaskToPostgres struct {
	db *sqlx.DB
}

func NewTasksToPostgres(db *sqlx.DB) *TaskToPostgres {
	return &TaskToPostgres{db: db}
}

func (t *TaskToPostgres) CreateTask(task entity.Task) (int, error) {

	var taskID int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (name, importance, status, description) VALUES ($1, $2, $3, $4) RETURNING id", taskTable)
	row := t.db.QueryRow(createTaskQuery, task.Name, task.Importance, task.Status, task.Description)

	err := row.Scan(&taskID)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return taskID, nil
}

func (t *TaskToPostgres) GetTask(taskId int) (entity.Task, error) {

	var task entity.Task

	query := fmt.Sprintf("SELECT name, importance, status, description FROM %s WHERE id = $1", taskTable)
	err := t.db.Get(&task, query, taskId)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (t *TaskToPostgres) GetTasks() ([]entity.Task, error) {

	var tasks []entity.Task

	query := fmt.Sprintf("SELECT name, importance, status, description FROM %s", taskTable)
	err := t.db.Select(&tasks, query)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (t *TaskToPostgres) DeleteTask(taskId int) error {

	query := fmt.Sprintf("DELETE FROM %s  WHERE id = $1", taskTable)
	_, err := t.db.Exec(query, taskId)
	return err
}
