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
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		return
	}
	models.CreateUser(user, DB)
}

func GetUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user *models.User

	userId, exists := c.Get("user")
	if exists {
		user = models.RetrieveUser(uint(userId.(int)), DB)
		c.JSON(200, &user)
	} else {
		c.JSON(400, ": (")
	}
}

func DeleteUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	userId, exists := c.Get("user")

	if exists {
		models.DeleteUser(uint(userId.(int)), DB)
		c.JSON(200, gin.H{"success": "user deleted"})
	} else {
		c.JSON(400, gin.H{"error": "No User ID in request"})
	}
}
