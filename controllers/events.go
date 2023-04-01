package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	DB := config.DB
	userId, userExists := c.Get("user")

	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	events, err := models.RetrieveEvents(uint(userId.(int)), DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, events)
}

func PostEvent(c *gin.Context) {
	DB := config.DB
	var event *models.Event
	userId, userExists := c.Get("user")

	event.UserID = uint(userId.(int))

	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err := c.BindJSON(&event); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.CreateEvent(event, DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &event)
}

func DeleteEvent(c *gin.Context) {
	DB := config.DB
	userId, userExists := c.Get("user")
	eventId, err := strconv.Atoi(c.Param("eventId"))

	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}
	if err = models.DeleteEvent(uint(userId.(uint)), uint(eventId), DB); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.Status(204)

}
