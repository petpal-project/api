package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMedicine(c *gin.Context) {
	var medicine *models.Medicine
	var DB *gorm.DB = config.DB

	medId, err := strconv.Atoi(c.Param("medicineId"))

	if err != nil {
		c.JSON(400, "Medicine ID must be numeric.")
		return
	}

	medicine, err = models.GetMedicine(uint(medId), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, medicine)
}

func GetMedicines(c *gin.Context) {
	var medicines *[]models.Medicine
	var DB *gorm.DB = config.DB

	speciesId, err := strconv.Atoi(c.Query("speciesId"))
	if err != nil {
		c.JSON(400, "Species ID must be numeric.")
		return
	}

	medicines, err = models.GetMedicines(uint(speciesId), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, medicines)
}
