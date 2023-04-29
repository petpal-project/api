package datasources

import (
	"pet-pal/api/models"

	"gorm.io/gorm"
)

type QueryableStruct interface {
	models.User | models.Pet | models.Breed | models.Species | models.Event | models.Medication | models.Medicine | models.Meal | models.Food | models.Image
	GetID() uint
}

func RetrieveSingleRecord[T QueryableStruct](structId uint, userId uint, DB *gorm.DB) (*T, error) {
	var record *T
	err := DB.First(&record, "id = ? AND user_id = ?", structId, userId).Error
	return record, err
}

func RetrieveMultipleRecords[T QueryableStruct](userId uint, DB *gorm.DB) (*[]T, error) {
	var records *[]T
	err := DB.Find(&records, "user_id = ?", userId).Error
	return records, err
}

func CreateRecord[T QueryableStruct](record *T, DB *gorm.DB) error {
	return DB.Create(&record).Error
}

func UpdateRecord[T models.OwnedObject](userID uint, record T, DB *gorm.DB) (*T, error) {
	// I do not like how this doesn't take a pointer where most other fns do but its fine
	err := DB.Set("user", userID).Set("struct", record.GetID()).Model(&record).Where("id = ?", record.GetID()).Updates(&record).Error
	return &record, err
}

func DeleteRecord[T models.OwnedObject](structId uint, userId uint, DB *gorm.DB) error {
	var record *T
	return DB.Set("user", userId).Set("struct", structId).Delete(&record, "id = ?", structId).Error
}
