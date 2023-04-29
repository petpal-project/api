package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SpeciesService struct {
	DB	*gorm.DB
}

func (s *SpeciesService) GetSpecies(c *gin.Context) {
	sid, err := strconv.Atoi(c.Param("speciesId"))

	if err != nil {
		c.JSON(400, "Species ID must be numeric")
	} else {
		if species, err := models.RetrieveSpecies(uint(sid), s.DB); err != nil {
			c.JSON(500, "Internal Server Error")
		} else {
			c.JSON(200, species)
		}
	}
}
