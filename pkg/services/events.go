package services

import (
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type EventService interface {
	GetEventsByUserId(userId uint) (*[]models.Event, error)
	CreateEvent(event *models.Event) error
	UpdateEvent(userId uint, event *models.Event) (*models.Event, error)
	DeleteEvent(userId uint, eventId uint) error
}

type GormEventService struct {
	DB	*gorm.DB
}

func (s GormEventService) GetEventsByUserId(userId uint) (*[]models.Event, error) {
	return models.RetrieveEvents(userId, s.DB)
}

func (s GormEventService) CreateEvent(event *models.Event) error {
	return models.CreateEvent(event, s.DB)
}

func (s GormEventService) UpdateEvent(userId uint, event *models.Event) (*models.Event, error) {
	return models.UpdateEvent(userId, event, s.DB)
}

func (s GormEventService) DeleteEvent(userId uint, eventId uint) error {
	return models.DeleteEvent(userId, eventId, s.DB)
}


