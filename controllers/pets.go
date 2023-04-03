package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	pet, err = models.RetrievePet(uint(pid), uint(uid.(int)), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &pet)
}

func GetPets(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pets *[]models.Pet
	var err error

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pets, err = models.RetrievePets(uint(uid.(int)), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, pets)
}

func PostPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet
	var err error

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	if err = c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
		return
	}

	pet.UserID = uint(uid.(int))
	if err = models.CreatePet(pet, DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &pet)
}

func PutPet(c *gin.Context) {
	var err error
	var pet *models.Pet
	var DB *gorm.DB = config.DB

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	if err = c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if pet, err = models.UpdatePet(uint(uid.(int)), uint(pid), pet, DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, &pet)
}

func DeletePet(c *gin.Context) {
	var DB *gorm.DB = config.DB

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	err = models.DeletePet(uint(pid), uint(uid.(int)), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
