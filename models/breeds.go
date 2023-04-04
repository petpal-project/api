package models

import "gorm.io/gorm"

type Breed struct {
	gorm.Model
	SpeciesID    uint    `json:"speciesId"`
	Species      Species `json:"-"`
	Name         string  `json:"name"`
	Size         string  `json:"size"`
	HeightMale   string  `json:"heightMale"`
	HeightFemale string  `json:"heightFemale"`
	WeightMale   string  `json:"weightMale"`
	WeightFemale string  `json:"weightFemale"`
	Coat         string  `json:"coat"`
	CoatDesc     string  `json:"coatDesc"`
	Colors       string  `json:"colors"`
	ColorsDesc   string  `json:"colorsDesc"`
	Energy       string  `json:"energy"`
	Activities   string  `json:"activities"`
}

func (breed Breed) GetID() uint { return breed.ID }

func RetrieveBreeds(speciesId uint, DB *gorm.DB) (*[]Breed, error) {
	var breeds *[]Breed
	err := DB.Model(&Breed{}).Find(&breeds, "species_id = ?", speciesId).Error
	return breeds, err
}

func RetrieveBreed(breedId uint, DB *gorm.DB) (*Breed, error) {
	var breed *Breed
	err := DB.Model(&Breed{}).First(&breed, "id = ?", breedId).Error
	return breed, err
}
