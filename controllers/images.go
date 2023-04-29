package controllers

import (
	"errors"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageService struct {
	DB	*gorm.DB
}

func (s *ImageService) GetImagesByPet(c *gin.Context) {
	userId, userExists := c.Get("user")
	petId, err := strconv.Atoi(c.Param("petId"))
	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, idMustBeNumeric)
	} else {
		images, err := models.RetrieveImagesByPet(uint(userId.(int)), uint(petId), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, images)
		}
	}
}

func (s *ImageService) GetImagesByUser(c *gin.Context) {
	userId, userExists := c.Get("user")

	if !userExists {
		c.JSON(400, missingUserId)
	} else {
		images, err := models.RetrieveImagesByUser(uint(userId.(int)), s.DB)
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, images)
		}
	}
}

func (s *ImageService) PostImage(c *gin.Context) {
	var image *models.Image

	userId, userExists := c.Get("user")

	if !userExists {
		c.JSON(400, missingUserId)
	} else if err := c.BindJSON(&image); err != nil {
		c.JSON(400, err.Error())
	} else {
		image.UserID = uint(userId.(int))
		_, err = models.RetrievePet(image.PetID, image.UserID, s.DB)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "pet not found")
		} else {
			err = models.CreateImage(image, s.DB)
			if err != nil {
				c.JSON(500, err.Error())
			} else {
				c.JSON(200, image)
			}
		}
	}
}

func (s *ImageService) DeleteImage(c *gin.Context) {
	userId, userExists := c.Get("user")
	imageId, err := strconv.Atoi(c.Param("imageId"))

	if !userExists {
		c.JSON(400, missingUserId)
	} else if err != nil {
		c.JSON(400, "Image ID must be numeric")
	} else if err = models.DeleteImage(uint(imageId), uint(userId.(int)), s.DB); err != nil {
		c.JSON(500, err.Error())
	} else {
		c.Status(204)
	}
}
