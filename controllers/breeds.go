package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBreed(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var breed *models.Breed
	bid, err := strconv.Atoi(c.Param("breedId"))

	if err != nil {
		c.JSON(400, "breedId must be numeric")
		return
	}

	breed, err = models.RetrieveBreed(uint(bid), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &breed)
}

func GetBreeds(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var breeds *[]models.Breed
	sid, err := strconv.Atoi(c.Query("speciesId"))

	if err != nil {
		c.JSON(400, "speciesId must be numeric")
		return
	}

	breeds, err = models.RetrieveBreeds(uint(sid), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &breeds)
}
