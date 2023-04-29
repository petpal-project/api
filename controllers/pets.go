package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PetService struct {
	DB *gorm.DB
}

func (s *PetService) GetPet(c *gin.Context) {
	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))

	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else {
		pet, err := models.RetrievePet(uint(pid), uint(uid.(int)), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, &pet)
		}
	}
}

func (s *PetService) GetPets(c *gin.Context) {
	uid, userExists := c.Get("user")
	if userExists {
		pets, err := models.RetrievePets(uint(uid.(int)), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, pets)
		}
	}
}

func (s *PetService) PostPet(c *gin.Context) {
	var pet *models.Pet

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
	} else if err := c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
	} else {
		pet.UserID = uint(uid.(int))
		if err = models.CreatePet(pet, s.DB); err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, &pet)
		}
	}
}

func (s *PetService) PutPet(c *gin.Context) {
	var pet *models.Pet

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else if err = c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
	} else if pet, err = models.UpdatePet(uint(uid.(int)), uint(pid), pet, s.DB); err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(201, &pet)
	}
}

func (s *PetService) DeletePet(c *gin.Context) {
	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
	}
	if userExists {
		err = models.DeletePet(uint(pid), uint(uid.(int)), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.Status(204)
		}
	} else {
		c.JSON(400, missingUserId)
	}
}
