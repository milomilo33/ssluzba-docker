package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Ime     string `json:"ime" gorm:"not null"`
	Prezime string `json:"prezime" gorm:"not null"`
	JMBG    string `json:"jmbg" gorm:"unique;not-null"`
}

type Profesor struct {
	gorm.Model
	Ime     string `json:"ime" gorm:"not null"`
	Prezime string `json:"prezime" gorm:"not null"`
	JMBG    string `json:"jmbg" gorm:"unique;not-null"`
}
