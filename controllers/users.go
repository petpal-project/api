package controllers

import (
	"fmt"
	"pet-pal/api/config"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserIdFromFirebaseId(tokenUID string) uint {
	var DB *gorm.DB = config.DB
	var user models.User
	DB.Where("account_id = ?", tokenUID).First(&user)
	return user.ID
}

func PostUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		return
	}
	err := models.CreateUser(user, DB)
	if err != nil {
		c.JSON(500, "Internal Server Error")
	} else {
		c.JSON(200, fmt.Sprintf("User with ID %d Created", user.ID))
	}
}

func GetUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user *models.User
	var err error

	userId, exists := c.Get("user")
	if exists {
		user, err = models.RetrieveUser(uint(userId.(int)), DB)
		if err != nil {
			c.JSON(500, "Internal Server Error")
		} else {
			c.JSON(200, &user)
		}
	} else {
		c.JSON(400, "Missing User ID in Authorization Header")
	}
}

func DeleteUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	userId, exists := c.Get("user")

	if exists {
		err := models.DeleteUser(uint(userId.(int)), DB)
		if err != nil {
			c.JSON(500, "Internal Server Error")
		} else {
			c.JSON(200, gin.H{"success": "user deleted"})
		}
	} else {
		c.JSON(400, gin.H{"error": "No User ID in request"})
	}
}
