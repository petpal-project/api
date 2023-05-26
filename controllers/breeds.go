package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBreed(DB *gorm.DB) func(c *gin.Context) {
	return func (c *gin.Context) {
		bid, err := strconv.Atoi(c.Param("breedId"))

		if err != nil {
			c.JSON(400, "breedId must be numeric")
			return
		}
	
		breed, err := models.RetrieveBreed(uint(bid), DB)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.JSON(200, &breed)
	}
}

func GetBreeds(DB *gorm.DB) func (c *gin.Context) {
	return func (c *gin.Context) {
		sid, err := strconv.Atoi(c.Query("speciesId"))

		if err != nil {
			c.JSON(400, "speciesId must be numeric")
			return
		}
	
		breeds, err := models.RetrieveBreeds(uint(sid), DB)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
	
		c.JSON(200, &breeds)
	}
}
