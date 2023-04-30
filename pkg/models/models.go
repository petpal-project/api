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

func (meal Meal) GetID() uint { return meal.ID }

type Medication struct {
	gorm.Model
	PetID     uint
	Frequency uint
	StartDate time.Time
	EndDate   time.Time
	Medicine  []Medicine `gorm:"many2many:medication_medicines"`
}

func (medication Medication) GetID() uint { return medication.ID }
