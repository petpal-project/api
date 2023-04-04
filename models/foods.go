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

func RetrieveFood(foodId uint, DB *gorm.DB) (food *Food, err error) {
	err = DB.Find(&food, "id = ?", foodId).Error
	return
}

func RetrieveFoods(speciesId uint, DB *gorm.DB) (foods *[]Food, err error) {
	err = DB.Find(&foods, "target_species = ?", speciesId).Error
	return
}
