package controllers

import (
	"pet-pal/api/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserIdFromFirebaseId(tokenUID string, DB *gorm.DB) uint {
	var user models.User
	DB.Where("accountID = ?", tokenUID).First(&user)
	return user.ID
}

func PostUser(DB *gorm.DB, fname string, lname string, tokenUID string, userID uint) {
	var model gorm.Model = gorm.Model{}
	model.ID, model.CreatedAt, model.UpdatedAt = userID, time.Now(), time.Now()
	var user models.User = models.User{Model: model, AccountID: tokenUID, FirstName: fname, LastName: lname, Pets: nil}
	DB.Create(&user)
}

func GetUser(DB *gorm.DB, c *gin.Context) {
	userId, err := c.Get("user")
	if !err {
		var user models.User
		DB.Where("id = ?", userId).First(&user)
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{"error": "User does not exist"})
	}
}

func DeleteUser(DB *gorm.DB, c *gin.Context) {
	userId, err := c.Get("user")
	if !err {
		var user models.User
		DB.Where("id = ?", userId).Delete(&user)
	} else {
		c.JSON(400, gin.H{"error": "User does not exist"})
	}
}
