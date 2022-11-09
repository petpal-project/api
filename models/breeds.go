package models

import "gorm.io/gorm"

type Breed struct {
	gorm.Model
	SpeciesID    uint `json:"speciesId"`
	Species      Species
	Name         string  `json:"name"`
	Size         string  `json:"size"`
	HeightMale   float64 `json:"heightMale"`
	HeightFemale float64 `json:"heightFemale"`
	WeightMale   float64 `json:"weightMale"`
	WeightFemale float64 `json:"weightFemale"`
	Coat         string  `json:"coat"`
	CoatDesc     string  `json:"coatDesc"`
	Colors       string  `json:"colors"`
	ColorsDesc   string  `json:"colorsDesc"`
	Energy       string  `json:"energy"`
	Activities   string  `json:"activities"`
}

func RetrieveBreeds(speciesId uint, DB *gorm.DB) (breeds *[]Breed, err error) {
	err = DB.Find(&breeds, "species_id = ?", speciesId).Error
	return
}

func RetrieveBreed(breedId uint, DB *gorm.DB) (breed *Breed, err error) {
	err = DB.First(&breed, "id = ?", breedId).Error
	return
}
