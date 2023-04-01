package controllers

import (
	"pet-pal/api/config"
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	DB := config.DB
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)

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
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)

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

func PutEvent(c *gin.Context) {
	DB := config.DB
	var event *models.Event
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
	eventId, err := strconv.Atoi(c.Param("eventId"))

	if !userExists {
		c.JSON(400, missingUserId)
		return
	}
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}
	event.ID = uint(eventId)
	if err = c.BindJSON(&event); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if event, err = models.UpdateEvent(uint(userId.(int)), event, DB); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, event)
}

func DeleteEvent(c *gin.Context) {
	DB := config.DB
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
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
