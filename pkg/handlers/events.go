package handlers

import (
	"pet-pal/api/pkg/models"
	"pet-pal/api/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	EventService services.EventService
}

func (h EventHandler) GetEvents(c *gin.Context) {
	userId, userExists := c.Get(GIN_CONTEXT_USER_KEY)
	if !userExists {
		c.JSON(400, missingUserId)
		return
	}

	events, err := h.EventService.GetEventsByUserId(uint(userId.(int)))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, events)
}

func (h EventHandler) PostEvent(c *gin.Context) {
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

	if err := h.EventService.CreateEvent(event); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, &event)
}

func (h EventHandler) PutEvent(c *gin.Context) {
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

	if event, err = h.EventService.UpdateEvent(uint(userId.(int)), event); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, event)
}

func (h EventHandler) DeleteEvent(c *gin.Context) {
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

	if err = h.EventService.DeleteEvent(uint(userId.(int)), uint(eventId)); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.Status(204)
}
