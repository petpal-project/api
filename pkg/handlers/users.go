package handlers

import (
	"fmt"
	"pet-pal/api/pkg/models"
	"pet-pal/api/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func (h UserHandler) PostUser(c *gin.Context) {
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := h.UserService.CreateUser(user); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, fmt.Sprintf("User with ID %d Created", user.ID))
}

func (h UserHandler) GetUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		c.JSON(400, missingUserId)
		return
	}

	user, err := h.UserService.GetUserById(uint(userId.(int)))
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, &user)
}

func (h UserHandler) DeleteUser(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		c.JSON(400, missingUserId)
		return
	}

	if err := h.UserService.DeleteUser(uint(userId.(int))); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
