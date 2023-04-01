package models

import (
	"errors"

	"gorm.io/gorm"
)

type OwnedObject interface {
	Pet | Image | Event | Meal | Medication
	GetUserID() uint
}

func checkOwnership[T OwnedObject](key string, DB *gorm.DB) error {
	var emptyStruct T

	userId, userExists := DB.Get("user")
	id, exists := DB.Get(key)

	if !userExists {
		return errors.New("missing user id")
	}
	if !exists {
		return errors.New("missing event id")
	}

	if err := DB.Where("id = ?", id).Select("user_id").Find(&emptyStruct).Error; err != nil {
		return err
	}
	if emptyStruct.GetUserID() != userId {
		return errors.New("event does not belong to user")
	}
	return nil
}
