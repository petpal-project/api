package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FoodService struct {
	DB	*gorm.DB
}

func (s *FoodService) GetFood(c *gin.Context) {
	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		c.JSON(400, "Food ID must be numeric.")
		return
	}

	food, err := models.RetrieveFood(uint(foodId), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, food)
}

func (s *FoodService) GetFoods(c *gin.Context) {
	speciesId, err := strconv.Atoi(c.Query("speciesId"))
	if err != nil {
		c.JSON(400, "Species ID must be numeric.")
	}

	foods, err := models.RetrieveFoods(uint(speciesId), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, foods)
}
