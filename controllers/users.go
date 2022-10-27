package controllers

import (
	"pet-pal/api/models"

	"gorm.io/gorm"
)

func GetUserIdFromFirebaseId(tokenUID string, DB *gorm.DB) uint {
	var user models.User
	DB.Where("account_id = ?", tokenUID).First(&user)
	return user.ID
}
