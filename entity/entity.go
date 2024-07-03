package entity

import "time"

// import "time"

type Manager struct {
	Id          int    `db:"id" swagg:"-"`
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
	Status      string `db:"status" json:"status"`
	Description string `db:"description" json:"description"`
}

type Tracker struct {
	TaskID      int       `db:"task_id" json:"task_id"`
	PeopleID    int       `db:"people_id" json:"people_id"`
	Created_at  time.Time `db:"created_at" json:"created_at"`
	Finished_at time.Time `db:"finished_at" json:"finished_at"`
}
