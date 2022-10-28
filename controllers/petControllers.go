package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet

	uid, userExists := c.Get("user")
	pid, petExists := c.Get("pet")

	if petExists && userExists {
		DB.First(&pet, "id = ? AND user_id = ?", uint(pid.(int)), uint(uid.(int)))
		c.JSON(200, &pet)
	} else {
		c.JSON(400, "Missing Pet or User ID")
	}
}
