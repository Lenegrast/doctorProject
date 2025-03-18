package models

import "time"

type User struct {
	id int
}

type Medicine struct {
	id     int
	Name   string
	Dosage int
	Period int
}

type Schedule struct {
	id         int
	IdUser     int
	IdMedicine int
	DateStart  time.Time
}
