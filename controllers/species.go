package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSpecies(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
		sid, err := strconv.Atoi(c.Param("speciesId"))
		if err != nil {
			c.JSON(400, "Species ID must be numeric")
			return
		}
	
		species, err := models.RetrieveSpecies(uint(sid), DB)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.JSON(200, species)
	}
}
