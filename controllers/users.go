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
		c.JSON(400, err.Error())
		return
	}

	if err := models.CreateUser(user, DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, fmt.Sprintf("User with ID %d Created", user.ID))
}

func GetUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var user *models.User
	var err error

	userId, exists := c.Get("user")
	if !exists {
		c.JSON(400, missingUserId)
		return
	}

	user, err = models.RetrieveUser(uint(userId.(int)), DB)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var DB *gorm.DB = config.DB
	userId, exists := c.Get("user")

	if !exists {
		c.JSON(400, missingUserId)
		return
	}
	if err := models.DeleteUser(uint(userId.(int)), DB); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.Status(204)
}
