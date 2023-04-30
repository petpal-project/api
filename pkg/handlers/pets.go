package handlers

import (
	"net/http"
	"pet-pal/api/pkg/models"
	"pet-pal/api/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	PetService services.PetService
}

func (h PetHandler) GetPet(c *gin.Context) {
	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	pet, err := h.PetService.GetPetById(uint(pid))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if pet.UserID != uid {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.JSON(200, &pet)
}

func (h PetHandler) GetPets(c *gin.Context) {
	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pets, err := h.PetService.GetPetsByUserId(uint(uid.(int)))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, pets)
}

func (h PetHandler) PostPet(c *gin.Context) {
	var pet *models.Pet

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	if err := c.BindJSON(&pet); err != nil {
		c.JSON(400, err.Error())
		return
	}

	pet.UserID = uint(uid.(int))
	if err := h.PetService.CreatePet(pet); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &pet)
}

func (h PetHandler) PutPet(c *gin.Context) {
	var petForm models.Pet

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	if err = c.BindJSON(&petForm); err != nil {
		c.JSON(400, err.Error())
		return
	}

	petForm.ID = uint(pid)

	pet, err := h.PetService.UpdatePet(uint(uid.(int)), petForm)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, &pet)
}

func (h PetHandler) DeletePet(c *gin.Context) {
	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	if err = h.PetService.DeletePet(uint(pid), uint(uid.(int))); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
