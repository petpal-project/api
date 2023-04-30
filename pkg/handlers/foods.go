package handlers

import (
	"pet-pal/api/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	FoodService services.FoodService
}

func (h FoodHandler) GetFood(c *gin.Context) {
	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		c.JSON(400, "Food ID must be numeric.")
		return
	}

	food, err := h.FoodService.GetFoodById(uint(foodId))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, food)
}

func (h FoodHandler) GetFoods(c *gin.Context) {
	speciesId, err := strconv.Atoi(c.Query("speciesId"))
	if err != nil {
		c.JSON(400, "Species ID must be numeric.")
	}

	foods, err := h.FoodService.GetFoodsBySpeciesId(uint(speciesId))
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, foods)
}
