package controllers

import (
	"fmt"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostUser(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
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
}

func GetUser(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
		userId, exists := c.Get("user")
		if !exists {
			c.JSON(400, missingUserId)
			return
		}
	
		user, err := models.RetrieveUser(uint(userId.(int)), DB)
		if err != nil {
			c.JSON(500, err.Error())
		}
	
		c.JSON(200, &user)
	}
}

func DeleteUser(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
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
}
