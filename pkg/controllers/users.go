package controllers

import (
	"fmt"
	"pet-pal/api/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) PostUser(c *gin.Context) {
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := models.CreateUser(user, s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, fmt.Sprintf("User with ID %d Created", user.ID))
}

func (s *UserService) GetUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		c.JSON(400, missingUserId)
		return
	}

	user, err := models.RetrieveUser(uint(userId.(int)), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, &user)
}

func (s *UserService) DeleteUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		c.JSON(400, missingUserId)
		return
	}

	if err := models.DeleteUser(uint(userId.(int)), s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
