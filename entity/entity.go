package entity

import "time"

type Manager struct {
	Name        string `db:"name" json:"name" binding:"required"`
	Managername string `db:"managername" json:"managername" binding:"required"`
	Password    string `db:"password_hash" json:"password" binding:"required"`
	Role        string `db:"role" json:"role"`
}

type People struct {
	Surname    string `db:"surname" json:"surname"`
	Name       string `db:"name" json:"name"`
	Patronymic string `db:"patronymic" json:"patronymic"`
	Address    string `db:"address" json:"address"`
}

type Task struct {
	Name        string `db:"name" json:"name"`
	Importance  string `db:"importance" json:"importance"`
	Description string `db:"description" json:"description"`
}

type Tracker struct {
	TaskID      int       `db:"task_id" json:"task_id"`
	PeopleID    int       `db:"people_id" json:"people_id"`
	Created_at  time.Time `db:"created_at" json:"created_at"`
	Finished_at time.Time `db:"finished_at" json:"finished_at"`
}

type Effort struct {
	TaskID          string `db:"task_id" json:"task_id"`
	TaskDescription string `db:"description" json:"description"`
	TotalTime       string    `db:"total_time" json:"total_time"`
}
