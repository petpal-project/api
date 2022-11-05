package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSpecies(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var species *models.Species
	var requestBody models.RequestBody

	if err := c.BindHeader(&requestBody); err != nil {
		return
	}
	var sid uint = requestBody.SpeciesId
	species = models.RetrieveSpecies(sid, DB)
	c.JSON(200, species)
}
