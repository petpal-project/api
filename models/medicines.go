package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Description   string
	TargetSpecies uint
}

func GetMedicine(medId uint, DB *gorm.DB) (*Medicine, error) {
	var medicine *Medicine
	err := DB.First(&medicine, "id =  ?", medId).Error
	return medicine, err
}

func GetMedicines(speciesId uint, DB *gorm.DB) (*[]Medicine, error) {
	var medicines *[]Medicine
	err := DB.Find(&medicines, "target_species = ?", speciesId).Error
	return medicines, err
}
