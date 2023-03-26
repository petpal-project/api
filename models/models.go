package models

import (
	"time"

	"gorm.io/gorm"
)

type Meal struct {
	gorm.Model
	PetID     uint
	Frequency uint
	Food      []Food `gorm:"many2many:meal_foods;"`
}

type Medication struct {
	gorm.Model
	PetID     uint
	Frequency uint
	StartDate time.Time
	EndDate   time.Time
	Medicine  []Medicine `gorm:"many2many:medication_medicines"`
}

type Event struct {
	gorm.Model
	PetID       uint
	Name        string
	Category    string
	Description string
	EventDate   time.Time
}
