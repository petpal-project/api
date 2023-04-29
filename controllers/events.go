package controllers

import (
	"pet-pal/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventService struct {
	DB *gorm.DB
}

func (s *EventService) GetEvents(c *gin.Context) {
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	events, err := models.RetrieveEvents(uint(userId.(int)), s.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, events)
}

func (s *EventService) PostEvent(c *gin.Context) {
	var event *models.Event = &models.Event{}

	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	event.UserID = uint(userId.(int))
	if err := c.BindJSON(&event); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := models.CreateEvent(event, s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &event)
}

func (s *EventService) PutEvent(c *gin.Context) {
	var event *models.Event = &models.Event{}

	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	eventId, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	event.ID = uint(eventId)
	if err = c.BindJSON(&event); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if event, err = models.UpdateEvent(uint(userId.(int)), event, s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, event)
}

func (s *EventService) DeleteEvent(c *gin.Context) {
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	eventId, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(400, idMustBeNumeric)
		return
	}

	if err = models.DeleteEvent(uint(userId.(int)), uint(eventId), s.DB); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
