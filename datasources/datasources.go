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

func CreateRecord[T QueryableStruct](record T, DB *gorm.DB) error {
	return DB.Create(&record).Error
}

func UpdateRecord[T QueryableStruct](userID uint, record T, DB *gorm.DB) (T, error) {
	err := DB.Set("user", userID).Set("struct", record.GetID()).Delete(&record, "id = ?", record.GetID()).Error
	return record, err
}
