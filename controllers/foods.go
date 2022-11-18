package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFood(c *gin.Context) {
	var food *models.Food
	var DB *gorm.DB = config.DB

	foodId, err := strconv.Atoi(c.Param("foodId"))

	if err != nil {
		c.JSON(400, "Food ID must be numeric.")
	} else {
		food, err = models.RetrieveFood(uint(foodId), DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, food)
		}
	}
}

func GetFoods(c *gin.Context) {
	var foods *[]models.Food
	var DB *gorm.DB = config.DB

	speciesId, err := strconv.Atoi(c.Query("speciesId"))

	if err != nil {
		c.JSON(400, "Species ID must be numeric.")
	} else {
		foods, err = models.RetrieveFoods(uint(speciesId), DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, foods)
		}
	}
}
