package models

import (
	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	Name         string `json:"name"`
	BinomialName string `json:"binomialName"`
}

func RetrieveSpecies(speciesId uint, DB *gorm.DB) (*Species, error) {
	var species *Species
	var err error
	if err = DB.First(&species, "id = ?", speciesId).Error; err != nil {
		return species, err
	} else {
		return species, err
	}

}
