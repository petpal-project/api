package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountID string `json:"accountId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Pets      []Pet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func GetUserIdFromFirebaseId(tokenUID string, DB *gorm.DB) uint {
	var user User
	DB.Where("accountID = ?", tokenUID).First(&user)
	return user.ID
}

func CreateUser(user *User, DB *gorm.DB) {
	DB.Create(&user)
}

func RetrieveUser(userId uint, DB *gorm.DB) *User {
	var user *User
	DB.First(&user, "id = ?", userId)

	return user
}

// TODO: make this return a bool
func DeleteUser(userId uint, DB *gorm.DB) {
	var user User
	DB.Unscoped().Where("id = ?", userId).Delete(&user)
}
