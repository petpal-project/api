package models

import (
	"time"

	"gorm.io/gorm"
)

type Breed struct {
	gorm.Model
	SpeciesID    uint `json:"speciesId"`
	Species      Species
	Name         string  `json:"name"` 
	Size         string  `json:"size"`
	HeightMale   string `json:"heightMale"`
	HeightFemale string `json:"heightFemale"`
	WeightMale   string `json:"weightMale"`
	WeightFemale string `json:"weightFemale"`
	Coat         string  `json:"coat"`
	CoatDesc     string  `json:"coatDesc"`
	Colors       string  `json:"colors"`
	ColorsDesc   string  `json:"colorsDesc"`
	Energy       string  `json:"energy"`
	Activities   string  `json:"activities"`
}

type Meal struct {
	gorm.Model
	PetID     uint
	Frequency uint
	Food      []Food `gorm:"many2many:meal_foods;"`
}

type Food struct {
	gorm.Model
	Category    string
	TargetAge   uint
	Description string
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

type Medicine struct {
	gorm.Model
	Description string
}

type HealthEvent struct {
	gorm.Model
	PetID       uint
	Name        string
	Category    string
	Description string
}

type Image struct {
	gorm.Model
	PetID       uint
	AssetUrl    string
	Description string
}
