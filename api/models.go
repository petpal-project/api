package api

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountID string `json:"accountId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Pets      []Pet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Pet struct {
	gorm.Model
	UserID       uint          `json:"userId"`
	Name         string        `json:"name"`
	Breeds       []Breed       `gorm:"many2many:pet_breeds;"`
	SpeciesID    uint          `json:"speciesId"`
	Species      Species       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Age          uint          `json:"age"`
	Images       []Image       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Meals        []Meal        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Medications  []Medication  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HealthEvents []HealthEvent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Species struct {
	gorm.Model
	Name         string `json:"name"`
	BinomialName string `json:"binomialName"`
}

type Breed struct {
	gorm.Model
	SpeciesID    uint `json:"speciesId"`
	Species      Species
	Name         string   `json:"name"`
	Size         string   `json:"size"`
	HeightMale   float64  `json:"heightMale"`
	HeightFemale float64  `json:"heightFemale"`
	WeightMale   float64  `json:"weightMale"`
	WeightFemale float64  `json:"weightFemale"`
	Coat         string   `json:"coat"`
	CoatDesc     string   `json:"coatDesc"`
	Colors       []string `json:"colors"`
	ColorsDesc   string   `json:"colorsDesc"`
	Energy       string   `json:"energy"`
	Activities   []string `json:"activities"`
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
	PetID     uint
	Frequency uint
	StartDate time.Time
	EndDate   time.Time
	Medicine  Medicine `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Medicine struct {
	gorm.Model
	MedicationID uint
	Description  string
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
