package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Description   string
	TargetSpecies uint
}

func GetMedicine(medId uint, DB *gorm.DB) (med *Medicine, err error) {
	err = DB.First(&med, "id =  ?", medId).Error
	return
}

func GetMedicines(speciesId uint, DB *gorm.DB) (meds *[]Medicine, err error) {
	err = DB.Find(&meds, "target_species = ?", speciesId).Error
	return
}
