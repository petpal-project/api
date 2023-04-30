package handlers

import (
	"pet-pal/api/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SpeciesHandler struct {
	SpeciesService services.SpeciesService
}

func (s SpeciesHandler) GetSpecies(c *gin.Context) {
	sid, err := strconv.Atoi(c.Param("speciesId"))
	if err != nil {
		c.JSON(400, "Species ID must be numeric")
		return
	}

	species, err := s.SpeciesService.GetSpeciesById(uint(sid))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, species)
}
