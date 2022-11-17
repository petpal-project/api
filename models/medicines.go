package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Description string
}

func GetMedicine(medId uint, DB *gorm.DB) (med *Medicine, err error) {
	err = DB.First(&Medicine{}, "id =  ?", medId).Error
	return
}

func GetMedicines(DB *gorm.DB) (meds *[]Medicine, err error) {
	err = DB.Find(&meds).Error
	return
}
