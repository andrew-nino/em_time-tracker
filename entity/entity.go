package entity

// import "time"

type Manager struct {
	Id          int    `db:"id" swagg:"-"`
	Name        string `db:"name" json:"name" binding:"required"`
	Managername string `db:"managername" json:"managername" binding:"required"`
	Password    string `db:"password_hash" json:"password" binding:"required"`
	Role        string `db:"role" json:"role"`
}

type People struct {
	Surname        string    `db:"surname" json:"surname"`
	Name           string    `db:"name" json:"name"`
	Patronymic     string    `db:"patronymic" json:"patronymic"`
	Address        string    `db:"address" json:"address"`
}
