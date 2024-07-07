package entity

import "time"

type Manager struct {
	Name        string `db:"name" json:"name" binding:"required" example:"Andrew"`
	Managername string `db:"managername" json:"managername" binding:"required" example:"Manager"`
	Password    string `db:"password_hash" json:"password" binding:"required" example:"qwerty"`
	Role        string `db:"role" json:"-"`
}

type People struct {
	Surname    string `db:"surname" json:"surname" example:"Иванов"`
	Name       string `db:"name" json:"name" example:"Иван"`
	Patronymic string `db:"patronymic" json:"patronymic" example:"Иванович"`
	Address    string `db:"address" json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

type Task struct {
	Name        string `db:"name" json:"name" example:"T-001"`
	Importance  string `db:"importance" json:"importance" example:"low or high"`
	Description string `db:"description" json:"description" example:"A very important task"`
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
	TotalTime       string `db:"total_time" json:"total_time"`
}
