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
	var requestBody models.RequestBody

	uid, userExists := c.Get("user")
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	var pid uint = requestBody.PetId

	if userExists {
		pet = models.RetrievePet(pid, uint(uid.(int)), DB)
		c.JSON(200, &pet)
	} else {
		c.JSON(400, "Missing User ID")
	}
}

func PostPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet

	uid, userExists := c.Get("user")

	if err := c.BindJSON(&pet); err != nil {
		return
	}
	if userExists {
		pet.UserID = uint(uid.(int))
		models.CreatePet(pet, DB)
		c.JSON(200, &pet)
	} else {
		c.JSON(400, "Missing User ID in request")
	}
}

func DeletePet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var requestBody models.RequestBody

	uid, userExists := c.Get("user")
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	var pid uint = requestBody.PetId

	if userExists {
		models.DeletePet(pid, uint(uid.(int)), DB)
	} else {
		c.JSON(400, "Missing Pet or User ID")
	}
}
