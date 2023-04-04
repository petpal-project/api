package controllers

import (
	"errors"
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetImagesByPet(c *gin.Context) {
	var images *[]models.Image
	var DB *gorm.DB = config.DB

	userId, userExists := c.Get("user")
	petId, err := strconv.Atoi(c.Param("petId"))
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	images, err = models.RetrieveImagesByPet(uint(userId.(int)), uint(petId), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, images)
}

func GetImagesByUser(c *gin.Context) {
	var images *[]models.Image
	var DB *gorm.DB = config.DB
	var err error

	userId, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	images, err = models.RetrieveImagesByUser(uint(userId.(int)), DB)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, images)
}

func PostImage(c *gin.Context) {
	var DB *gorm.DB = config.DB
	var image *models.Image
	var err error

	userId, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	if err = c.BindJSON(&image); err != nil {
		c.JSON(400, err.Error())
		return
	}

	image.UserID = uint(userId.(int))
	_, err = models.RetrievePet(image.PetID, image.UserID, DB)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, "pet not found")
		return
	}

	err = models.CreateImage(image, DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, image)
}

func DeleteImage(c *gin.Context) {
	var DB *gorm.DB = config.DB

	userId, userExists := c.Get("user")
	imageId, err := strconv.Atoi(c.Param("imageId"))
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err != nil {
		c.JSON(400, "Image ID must be numeric")
		return
	}

	if err = models.DeleteImage(uint(imageId), uint(userId.(int)), DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
