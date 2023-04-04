package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	PetID       uint `json:"petId" binding:"required"`
	UserID      uint
	AssetUrl    string `json:"url" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (image Image) GetUserID() uint { return image.UserID }
func (image Image) GetID() uint     { return image.ID }

func (image *Image) BeforeDelete(DB *gorm.DB) error {
	return CheckOwnership[Image](DB)
}

func CreateImage(image *Image, DB *gorm.DB) error {
	return DB.Create(&image).Error
}

func RetrieveImagesByUser(userId uint, DB *gorm.DB) (*[]Image, error) {
	var images *[]Image
	err := DB.Find(&images, "user_id = ?", userId).Error
	return images, err
}

func RetrieveImagesByPet(userId uint, petId uint, DB *gorm.DB) (*[]Image, error) {
	var images *[]Image
	err := DB.Find(&images, "user_id = ? and pet_id = ?", userId, petId).Error
	return images, err
}

func DeleteImage(imageId uint, userId uint, DB *gorm.DB) error {
	return DB.Set("imageId", imageId).Set("userId", userId).Delete(&Image{}, "id = ?", imageId).Error
}
