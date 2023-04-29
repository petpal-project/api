package controllers

import (
	"pet-pal/api/datasources"
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
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	pet, err := datasources.RetrieveSingleRecord[models.Pet](uint(pid), uint(uid.(int)), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &pet)
}

func (s *PetService) GetPets(c *gin.Context) {
	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pets, err := datasources.RetrieveMultipleRecords[models.Pet](uint(uid.(int)), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, pets)
}

func (s *PetService) PostPet(c *gin.Context) {
	var pet *models.Pet

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	if err := c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
		return
	}

	pet.UserID = uint(uid.(int))
	if err := datasources.CreateRecord(pet, s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &pet)
}

func (s *PetService) PutPet(c *gin.Context) {
	var pet *models.Pet

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	if err = c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
		return
	}

	pet.ID = uint(pid)

	pet, err = datasources.UpdateRecord(uint(uid.(int)), *pet, s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, &pet)
}

func (s *PetService) DeletePet(c *gin.Context) {
	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	if err = datasources.DeleteRecord[models.Pet](uint(pid), uint(uid.(int)), s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
