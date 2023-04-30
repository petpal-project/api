package handlers

import (
	"pet-pal/api/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	MedicineService	services.MedicineService
}

func (h MedicineHandler) GetMedicine(c *gin.Context) {
	medId, err := strconv.Atoi(c.Param("medicineId"))
	if err != nil {
		c.JSON(400, "Medicine ID must be numeric.")
		return
	}

	medicine, err := h.MedicineService.GetMedicineById(uint(medId))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, medicine)
}

func (h MedicineHandler) GetMedicines(c *gin.Context) {
	speciesId, err := strconv.Atoi(c.Query("speciesId"))
	if err != nil {
		c.JSON(400, "Species ID must be numeric.")
		return
	}

	medicines, err := h.MedicineService.GetMedicinesBySpeciesId(uint(speciesId))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, medicines)
}
