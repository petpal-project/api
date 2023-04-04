package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const GORM_CONTEXT_USER_KEY = "user"
const GORM_CONTEXT_STRUCT_KEY = "struct"
const GORM_CONTEXT_MISSING_USER = "missing user id"
const GORM_CONTEXT_MISSING_OBJECT = "missing object id"

type OwnedObject interface {
	Pet | Image | Event | Meal | Medication
	GetUserID() uint
	GetID() uint
}

func CheckOwnership[T OwnedObject](DB *gorm.DB) error {
	var emptyStruct T

	userId, userExists := DB.Get(GORM_CONTEXT_USER_KEY)
	id, exists := DB.Get(GORM_CONTEXT_STRUCT_KEY)

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
	returnedRecordId := emptyStruct.GetUserID()
	if returnedRecordId != userId {
		return errors.New("requested record does not belong to user")
	}
	return nil
}
