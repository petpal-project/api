package controllers

import (
	"fmt"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) PostUser(c *gin.Context) {
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		return
	}
	err := models.CreateUser(user, s.DB)
	if err != nil {
		c.JSON(500, "Internal Server Error")
	} else {
		c.JSON(200, fmt.Sprintf("User with ID %d Created", user.ID))
	}
}

func (s *UserService) GetUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if exists {
		user, err := models.RetrieveUser(uint(userId.(int)), s.DB)
		if err != nil {
			c.JSON(500, "Internal Server Error")
		} else {
			c.JSON(200, &user)
		}
	} else {
		c.JSON(400, "Missing User ID in Authorization Header")
	}
}

func (s *UserService) DeleteUser(c *gin.Context) {
	userId, exists := c.Get("user")

	if exists {
		err := models.DeleteUser(uint(userId.(int)), s.DB)
		if err != nil {
			c.JSON(500, "Internal Server Error")
		} else {
			c.JSON(200, gin.H{"success": "user deleted"})
		}
	} else {
		c.JSON(400, gin.H{"error": "No User ID in request"})
	}
}
