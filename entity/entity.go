package entity

import "time"

type Manager struct {
	Id          int    `db:"id" swagg:"-"`
	Name        string `db:"name" json:"name" binding:"required"`
	Managername string `db:"managername" json:"managername" binding:"required"`
	Password    string `db:"password_hash" json:"password" binding:"required"`
	Role        string `db:"role" json:"role"`
}

type People struct {
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Patronymic     string    `json:"patronymic"`
	Address        string    `json:"address"`
	PassportSerie  string    `db:"passportSerie"  json:"passportSerie"   binding:"required"`
	PassportNumber string    `db:"passportNumber" json:"passportNumber"  binding:"required"`
	CreatedAt      time.Time `db:"createdAt" json:"createdAt" swagg:"-" `
}
