package controllers

import (
	"errors"
	"pet-pal/api/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageService struct {
	DB *gorm.DB
}

func (s *ImageService) GetImagesByPet(c *gin.Context) {
	userId, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	petId, err := strconv.Atoi(c.Param("petId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	images, err := models.RetrieveImagesByPet(uint(userId.(int)), uint(petId), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, images)
}

func (s *ImageService) GetImagesByUser(c *gin.Context) {
	userId, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	images, err := models.RetrieveImagesByUser(uint(userId.(int)), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, images)
}

func (s *ImageService) PostImage(c *gin.Context) {
	var image *models.Image

	userId, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	if err := c.BindJSON(&image); err != nil {
		c.JSON(400, err.Error())
		return
	}

	image.UserID = uint(userId.(int))
	_, err := models.RetrievePet(image.PetID, image.UserID, s.DB)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, "pet not found")
		return
	}

	err = models.CreateImage(image, s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, image)
}

func (s *ImageService) DeleteImage(c *gin.Context) {
	userId, userExists := c.Get("user")
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	imageId, err := strconv.Atoi(c.Param("imageId"))
	if err != nil {
		c.JSON(400, "Image ID must be numeric")
		return
	}

	if err = models.DeleteImage(uint(imageId), uint(userId.(int)), s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
