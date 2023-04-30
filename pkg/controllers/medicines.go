package controllers

import (
	"pet-pal/api/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MedicineService struct {
	DB *gorm.DB
}

func (s *MedicineService) GetMedicine(c *gin.Context) {
	medId, err := strconv.Atoi(c.Param("medicineId"))
	if err != nil {
		c.JSON(400, "Medicine ID must be numeric.")
		return
	}

	medicine, err := models.GetMedicine(uint(medId), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, medicine)
}

func (s *MedicineService) GetMedicines(c *gin.Context) {
	speciesId, err := strconv.Atoi(c.Query("speciesId"))
	if err != nil {
		c.JSON(400, "Species ID must be numeric.")
		return
	}

	medicines, err := models.GetMedicines(uint(speciesId), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, medicines)
}
