package models

import (
	"errors"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	PetID       uint `json:"petId" binding:"required"`
	UserID      uint
	AssetUrl    string `json:"url" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (image *Image) BeforeUpdate(DB *gorm.DB) (err error) {
	var imageInDB *Image
	userId, userExists := DB.Get("userId")
	imageId, imageExists := DB.Get("imageId")
	if !userExists {
		err = errors.New("missing user id")
	} else if !imageExists {
		err = errors.New("missing image id")
	} else if err = DB.Select("user_id").First(&imageInDB, "id = ? ", imageId).Error; err != nil {
		return
	} else if imageInDB.UserID != userId {
		err = errors.New("image does not belong to user")
	}
	return
}

func CreateImage(image *Image, DB *gorm.DB) (err error) {
	err = DB.Create(&image).Error
	return
}

func RetrieveImagesByUser(userId uint, DB *gorm.DB) (images *[]Image, err error) {
	err = DB.Find(&images, "user_id = ?", userId).Error
	return
}

func RetrieveImagesByPet(userId uint, petId uint, DB *gorm.DB) (images *[]Image, err error) {
	err = DB.Find(&images, "user_id = ? and pet_id = ?", userId, petId).Error
	return
}

func DeleteImage(imageId uint, userId uint, DB *gorm.DB) (err error) {
	err = DB.Set("imageId", imageId).Set("userId", userId).Delete(&Image{}, "id = ?", imageId).Error
	return
}
