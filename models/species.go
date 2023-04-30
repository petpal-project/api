package models

import (
	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	Name         string `json:"name"`
	BinomialName string `json:"binomialName"`
}

func (species Species) GetID() uint { return species.ID }

func RetrieveSpecies(speciesId uint, DB *gorm.DB) (*Species, error) {
	var species *Species
	err := DB.First(&species, "id = ?", speciesId).Error
	return species, err
}
