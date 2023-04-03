package datasources

import (
	"gorm.io/gorm"
)

func RetrieveSingleRecord[T struct{}](structId uint, userId uint, DB *gorm.DB) (*T, error) {
	var record *T
	err := DB.First(&record, "id = ? AND user_id = ?", structId, userId).Error
	return record, err
}

func RetrieveMultipleRecords[T struct{}](userId uint, DB *gorm.DB) (*[]T, error) {
	var records *[]T
	err := DB.Find(&records, "user_id = ?", userId).Error
	return records, err
}
