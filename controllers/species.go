package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSpecies(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var species *models.Species

	sid, err := strconv.Atoi(c.Param("speciesId"))

	if err != nil {
		c.JSON(400, "Species ID must be numeric")
	} else {
		if species, err = models.RetrieveSpecies(uint(sid), DB); err != nil {
			c.JSON(500, "Internal Server Error")
		} else {
			c.JSON(200, species)
		}
	}
}
