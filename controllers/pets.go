package controllers

import (
	"fmt"
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postPetRequestBody struct {
	Name           string `json:"name"`
	BreedIDs       []uint `json:"breedIDs"`
	SpeciesID      uint   `json:"speciesId"`
	Age            uint   `json:"age"`
	Images         []uint `json:"imageIDs"`
	MealIDs        []uint `json:"mealIDs"`
	MedicationIDs  []uint `json:"medicationIDs"`
	HealthEventIDs []uint `json:"healthEventIDs"`
}

type putPetRequestBody struct {
	Name           string `json:"name"`
	BreedIDs       []uint `json:"breedIDs"`
	SpeciesID      uint   `json:"speciesID"`
	Age            uint   `json:"age"`
	Images         []uint `json:"imageIDs"`
	MealIDs        []uint `json:"mealIDs"`
	MedicationIDs  []uint `json:"medicationIDs"`
	HealthEventIDs []uint `json:"healthEventIDs"`
}

func (reqBody *postPetRequestBody) bindToPet(pet *models.Pet, DB *gorm.DB) {
	species, _ := models.RetrieveSpecies(reqBody.SpeciesID, DB)
	pet.Name = reqBody.Name
	pet.Age = reqBody.Age
	pet.SpeciesID = reqBody.SpeciesID
	pet.Species = *species
}

func (reqBody *putPetRequestBody) bindToPet(pet *models.Pet, DB *gorm.DB) {
	if reqBody.Name != "" {
		pet.Name = reqBody.Name
	}
	if reqBody.Age != 0 {
		pet.Age = reqBody.Age
	}
	if reqBody.SpeciesID != 0 {
		pet.SpeciesID = reqBody.SpeciesID
		species, _ := models.RetrieveSpecies(pet.SpeciesID, DB)
		fmt.Print(species)
		pet.Species = *species
	}
}

const idMustBeNumeric = "Pet Id must be numeric"
const missingUserId = "Missing User ID"

func GetPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))

	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else {
		pet, err = models.RetrievePet(uint(pid), uint(uid.(int)), DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, &pet)
		}
	}
}

func GetPets(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pets *[]models.Pet
	var err error

	uid, userExists := c.Get("user")
	if userExists {
		pets, err = models.RetrievePets(uint(uid.(int)), DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, pets)
		}
	}
}

func PostPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var requestBody *postPetRequestBody
	var pet *models.Pet = &models.Pet{}
	var err error

	uid, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
	} else if err = c.BindJSON(&requestBody); err != nil {
		c.JSON(400, err.Error())
	} else {
		requestBody.bindToPet(pet, DB)
		pet.UserID = uint(uid.(int))
		if len(requestBody.BreedIDs) > 0 {
			var breed *models.Breed
			for _, breedId := range requestBody.BreedIDs {
				breed, _ = models.RetrieveBreed(breedId, DB)
				pet.Breeds = append(pet.Breeds, *breed)
			}
		}
		if err = models.CreatePet(pet, DB); err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, &pet)
		}
	}
}

func PutPet(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var pet *models.Pet = &models.Pet{}
	var requestBody *putPetRequestBody
	var err error

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else if err = c.BindJSON(&requestBody); err != nil {
		c.JSON(400, err.Error())
	} else {
		requestBody.bindToPet(pet, DB)
		if pet, err = models.UpdatePet(uint(uid.(int)), uint(pid), pet, DB); err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(201, &pet)
		}
	}
}

func DeletePet(c *gin.Context) {
	var DB *gorm.DB = config.DB

	uid, userExists := c.Get("user")
	pid, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
	}
	if userExists {
		err = models.DeletePet(uint(pid), uint(uid.(int)), DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.Status(204)
		}
	} else {
		c.JSON(400, missingUserId)
	}
}
