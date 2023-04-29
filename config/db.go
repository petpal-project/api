package config

import (
	"fmt"
	"log"
	"os"
	api "pet-pal/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(api.User{}, &api.Breed{}, &api.Food{}, &api.Event{}, &api.Image{}, &api.Meal{}, &api.Medication{}, &api.Medicine{}, &api.Pet{}, &api.Species{})
	if err != nil {
		log.Fatalf("An error occured while performing auto-migration: %v\n", err)
	} else {
		fmt.Println("Successfully ran automigration")
	}
}

func InitDb() *gorm.DB {
	var dsn string = os.ExpandEnv("host=$DB_HOST user=$DB_USER password=$DB_PASS dbname=$DB_NAME port=$DB_PORT sslmode=disable TimeZone=America/New_York")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening connection to database %v\n", err)
	}
	migrate(db)
	return db
}
