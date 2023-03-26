package models

import (
	"errors"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	UserID      uint         `json:"userId"`
	Name        string       `json:"name"`
	Breeds      []Breed      `gorm:"many2many:pet_breeds;"`
	SpeciesID   uint         `json:"speciesId"`
	Species     Species      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Age         uint         `json:"age"`
	Images      []Image      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Meals       []Meal       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Medications []Medication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Events      []Event      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (pet *Pet) BeforeUpdate(DB *gorm.DB) (err error) {
	var petInDB *Pet
	uid, userExists := DB.Get("user")
	pid, petExists := DB.Get("pet")
	if !userExists {
		err = errors.New("missing user id")
	} else if !petExists {
		err = errors.New("missing pet id")
	} else if err = DB.Select("user_id").First(&petInDB, "id = ?", pid).Error; err != nil {
		return
	} else if petInDB.UserID != uid {
		err = errors.New("pet does not belong to user")
	}
	return
}

func (pet *Pet) BeforeDelete(DB *gorm.DB) (err error) {
	var petInDB *Pet
	uid, userExists := DB.Get("user")
	pid, petExists := DB.Get("pet")
	if !userExists {
		err = errors.New("missing user id")
	} else if !petExists {
		err = errors.New("missing pet id")
	} else if err = DB.Select("user_id").First(&petInDB, "id = ?", pid).Error; err != nil {
		return
	} else if petInDB.UserID != uid {
		err = errors.New("pet does not belong to user")
	}
	return
}

func CreatePet(pet *Pet, DB *gorm.DB) (err error) {
	err = DB.Create(&pet).Error
	return
}

func RetrievePet(petId uint, userId uint, DB *gorm.DB) (pet *Pet, err error) {
	err = DB.First(&pet, "id = ? AND user_id = ?", petId, userId).Error
	return
}

func RetrievePets(userId uint, DB *gorm.DB) (pets *[]Pet, err error) {
	err = DB.Find(&pets, "user_id = ?", userId).Error
	return
}

func UpdatePet(userId uint, petId uint, pet *Pet, DB *gorm.DB) (updatedPet *Pet, err error) {
	err = DB.Set("user", userId).Set("pet", petId).Model(&pet).Where("id = ?", petId).Updates(&pet).Error
	updatedPet = pet
	return
}

func DeletePet(petId uint, userId uint, DB *gorm.DB) (err error) {
	var pet *Pet
	err = DB.Set("user", userId).Set("pet", petId).Delete(&pet, "id = ? AND user_id = ?", petId, userId).Error
	return
}
