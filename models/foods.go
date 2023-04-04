package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Category      string
	TargetAge     uint
	Description   string
	TargetSpecies uint
}

func (food Food) GetID() uint { return food.ID }

func RetrieveFood(foodId uint, DB *gorm.DB) (*Food, error) {
	var food *Food
	err := DB.Find(&food, "id = ?", foodId).Error
	return food, err
}

func RetrieveFoods(speciesId uint, DB *gorm.DB) (*[]Food, error) {
	var foods *[]Food
	err := DB.Find(&foods, "target_species = ?", speciesId).Error
	return foods, err
}
