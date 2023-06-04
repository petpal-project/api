package controllers

import (
	"pet-pal/api/datasources"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPet(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
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
	
		pet, err := datasources.RetrieveSingleRecord[models.Pet](uint(pid), uid.(string), DB)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.JSON(200, &pet)
	}

}

func GetPets(DB *gorm.DB) func(c *gin.Context) {
	return func (c *gin.Context) {
		uid, userExists := c.Get("user")
		if !userExists {
			c.JSON(400, missingUserId)
			return
		}

		pets, err := datasources.RetrieveMultipleRecords[models.Pet](uid.(string), DB)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, pets)
	}
}

func PostPet(DB *gorm.DB) func(c *gin.Context) {
	return func (c *gin.Context) {
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
	
		pet.UserID = uid.(string)
		if err := datasources.CreateRecord(pet, DB); err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.JSON(200, &pet)
	}
}

func PutPet(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
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
	
		pet, err = datasources.UpdateRecord(uid.(string), *pet, DB)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.JSON(201, &pet)
	}
}

func DeletePet(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
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
	
		if err = datasources.DeleteRecord[models.Pet](uint(pid), uid.(string), DB); err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.Status(204)
	}
}
