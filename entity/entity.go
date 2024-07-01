package entity

import "time"

type User struct {
	Id             int       `db:"id" swagg:"-"`
	PassportSerie  string    `db:"passportSerie"  json:"passportSerie"   binding:"required"`
	PassportNumber string    `db:"passportNumber" json:"passportNumber"  binding:"required"`
	CreatedAt      time.Time `db:"createdAt" json:"createdAt" swagg:"-" `
}

type People struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
