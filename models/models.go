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
	PetID      uint
	Frequency  uint
	StartDate  time.Time
	EndDate    time.Time
	MedicineID uint
	Medicine   Medicine `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Event struct {
	gorm.Model
	PetID       uint
	Name        string
	Category    string
	Description string
	EventDate   time.Time
}

type Image struct {
	gorm.Model
	PetID       uint
	AssetUrl    string
	Description string
}
