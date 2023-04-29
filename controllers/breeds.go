package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BreedService struct {
	DB *gorm.DB
}

func (s *BreedService) GetBreed(c *gin.Context) {
	bid, err := strconv.Atoi(c.Param("breedId"))

	if err != nil {
		c.JSON(400, "breedId must be numeric")
		c.Abort()
		return
	} else {
		breed, err := models.RetrieveBreed(uint(bid), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, &breed)
		}
	}
}

func (s *BreedService) GetBreeds(c *gin.Context) {
	sid, err := strconv.Atoi(c.Query("speciesId"))

	if err != nil {
		c.JSON(400, "speciesId must be numeric")
	} else {
		breeds, err := models.RetrieveBreeds(uint(sid), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, &breeds)
		}
	}
}
