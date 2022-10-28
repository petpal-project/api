package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.DB

func GetUserIdFromFirebaseId(tokenUID string) uint {
	var user models.User
	DB.Where("accountID = ?", tokenUID).First(&user)
	return user.ID
}

func PostUser(c *gin.Context) {
	var user models.User
	userId, exists := c.Get("user")
	if exists {
		user.ID = userId.(uint)
		DB.Create(&user)
	}
}

func GetUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if exists {
		var user models.User
		DB.Where("id = ?", userId).First(&user)
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{"error": "User does not exist"})
	}
}

func DeleteUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if exists {
		var user models.User
		DB.Where("id = ?", userId).Delete(&user)
	} else {
		c.JSON(400, gin.H{"error": "User does not exist"})
	}
}
