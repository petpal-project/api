package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const idMustBeNumeric = "Pet Id must be numeric"
const ISE = "Internal Server Error"
const missingUserId = "Missing User ID"
const badRequestBody = "Request Body Invalid"

func GetPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))

	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else {
		pet, err = models.RetrievePet(uint(pid), uint(uid.(int)), DB)
		if err != nil {
			c.JSON(500, ISE)
		} else {
			c.JSON(200, &pet)
		}
	}

}

func GetPets(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pets *[]models.Pet
	var err error

	uid, userExists := c.Get("user")
	if userExists {
		pets, err = models.RetrievePets(uint(uid.(int)), DB)
		if err != nil {
			c.JSON(500, ISE)
		} else {
			c.JSON(200, pets)
		}
	}
}

func PostPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet
	var err error
	uid, userExists := c.Get("user")

	if !userExists {
		c.JSON(400, missingUserId)
	} else {
		if err = c.BindJSON(&pet); err != nil {
			c.JSON(400, badRequestBody)
		} else {
			pet.UserID = uint(uid.(int))
			if err = models.CreatePet(pet, DB); err != nil {
				c.JSON(500, ISE)
			} else {
				c.JSON(200, "Pet Posted")
			}
		}
	}
}

func PutPet(c *gin.Context) {
	var err error
	var pet *models.Pet
	var DB *gorm.DB = config.DB

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))

	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else if err = c.BindJSON(&pet); err != nil {
		c.JSON(400, badRequestBody)
	} else if err = models.UpdatePet(uint(uid.(int)), uint(pid), pet, DB); err != nil {
		c.JSON(500, ISE)
	} else {
		c.JSON(200, "Pet Updated")
	}
}

func DeletePet(c *gin.Context) {
	var DB *gorm.DB = config.DB

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
	}
	if userExists {
		err = models.DeletePet(uint(pid), uint(uid.(int)), DB)
		if err != nil {
			c.JSON(500, ISE)
		} else {
			c.JSON(200, "Pet Deleted")
		}
	} else {
		c.JSON(400, missingUserId)
	}
}
