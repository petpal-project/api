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
}

type Pet struct {
	gorm.Model
	UserID      uint `json:"userId"`
	User        User
	Name        string `json:"name"`
	Breeds      []Breed
	Species     Species
	Age         uint `json:"age"`
	Images      []Image
	Meals       []Meal
	Medications []Medication
}

type Species struct {
	gorm.Model
	Name         string
	BinomialName string
}

type Breed struct {
	gorm.Model
	SpeciesID    uint
	Species      Species
	Name         string
	Size         string
	HeightMale   float64
	HeightFemale float64
	WeightMale   float64
	WeightFemale float64
	Coat         string
	CoatDesc     string
	Colors       []string
	ColorsDesc   string
	Energy       string
	Activities   []string
}

type Meal struct {
	gorm.Model
	PetID     uint
	Frequency []time.Time
	Food      []Food
}

type Food struct {
	gorm.Model
	Category    string
	TargetAge   uint
	description string
}

type Medication struct {
	gorm.Model
	PetID     uint
	Frequency uint
	StartDate time.Time
	EndDate   time.Time
	Medicine  Medicine
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
	AssetUrl    string
	Description string
}
