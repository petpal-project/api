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

	if err == nil {
		species = models.RetrieveSpecies(uint(sid), DB)
		c.JSON(200, species)
	} else {
		c.JSON(400, "Species ID must be numeric")
	}

}
