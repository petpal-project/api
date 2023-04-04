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

func (user User) GetID() uint { return user.ID }

func GetUserIdFromFirebaseId(tokenUID string, DB *gorm.DB) (uint, error) {
	var user User
	var err error = DB.Where("accountID = ?", tokenUID).First(&user).Error
	return user.ID, err
}

func CreateUser(user *User, DB *gorm.DB) error {
	var err error = DB.Create(&user).Error
	return err
}

func RetrieveUser(userId uint, DB *gorm.DB) (*User, error) {
	var user *User
	var err error = DB.First(&user, "id = ?", userId).Error
	return user, err
}

// TODO: make this return a bool
func DeleteUser(userId uint, DB *gorm.DB) error {
	var user User
	var err error = DB.Unscoped().Where("id = ?", userId).Delete(&user).Error
	return err
}
