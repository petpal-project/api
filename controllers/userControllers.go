package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserIdFromFirebaseId(tokenUID string) uint {
	var DB *gorm.DB = config.DB
	var user models.User
	DB.Where("accountID = ?", tokenUID).First(&user)
	return user.ID
}

func PostUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user models.User
	userId, exists := c.Get("user")
	if exists {
		if err := c.BindJSON(&user); err != nil {
			return
		}
		user.ID = uint(userId.(int))
		DB.Create(&user)
	} else {
		c.JSON(400, gin.H{"error": "No User ID in request"})
	}
}

func GetUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user models.User

	userId, exists := c.Get("user")

	if exists {
		DB.Where("id = ?", uint(userId.(int))).First(&user)
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{"error": "No User ID in request"})
	}
}

func DeleteUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user models.User

	userId, exists := c.Get("user")
	userId = uint(userId.(int))

	if exists {
		DB.Unscoped().Where("id = ?", userId).Delete(&user)
		c.JSON(200, gin.H{"success": "user deleted"})
	} else {
		c.JSON(400, gin.H{"error": "No User ID in request"})
	}
}
