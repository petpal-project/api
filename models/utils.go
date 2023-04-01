package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const GORM_CONTEXT_USER_KEY = "user"
const GORM_CONTEXT_MISSING_USER = "missing user id"
const GORM_CONTEXT_MISSING_OBJECT = "missing object id"

type OwnedObject interface {
	Pet | Image | Event | Meal | Medication
	GetUserID() uint
}

func checkOwnership[T OwnedObject](key string, DB *gorm.DB) error {
	var emptyStruct T

	userId, userExists := DB.Get(GORM_CONTEXT_USER_KEY)
	id, exists := DB.Get(key)

	if !userExists {
		return errors.New(GORM_CONTEXT_MISSING_USER)
	}
	if !exists {
		return errors.New(GORM_CONTEXT_MISSING_OBJECT)
	}
	fmt.Printf("UserID: %d    RecordID: %d \n", userId, id)
	if err := DB.Where("id = ?", id).Select("user_id").Find(&emptyStruct).Error; err != nil {
		return err
	}
	fmt.Printf("returned RecordID: %d \n", emptyStruct.GetUserID())
	returnedRecordId := emptyStruct.GetUserID()
	if returnedRecordId != uint(userId.(int)) {
		return errors.New("requested record does not belong to user")
	}
	return nil
}
