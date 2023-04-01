package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	PetID       uint
	UserID      uint
	Name        string
	Category    string
	Description string
	EventDate   time.Time
}

func (event Event) GetUserID() uint { return event.UserID }

func (event *Event) BeforeUpdate(DB *gorm.DB) error {
	return checkOwnership[Event]("event", DB)
}

func (event *Event) BeforeDelete(DB *gorm.DB) error {
	return checkOwnership[Event]("event", DB)
}

func CreateEvent(event *Event, DB *gorm.DB) error {
	err := DB.Create(&event).Error
	return err
}

func RetrieveEvents(userId uint, DB *gorm.DB) (*[]Event, error) {
	events := &[]Event{}
	err := DB.Find(&events, "user_id = ?", userId).Error
	return events, err
}

func UpdateEvent(userId uint, event *Event, DB *gorm.DB) (*Event, error) {
	err := DB.Set("user", userId).Set("event", event.ID).Model(&event).Where("id = ?", event.ID).Updates(&event).Error
	return event, err
}

func DeleteEvent(userId uint, eventId uint, DB *gorm.DB) error {
	var event *Event
	err := DB.Set("user", userId).Set("event", eventId).Delete(&event, "id = ? AND user_id = ?", eventId, userId).Error
	return err
}
