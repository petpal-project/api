package config

import (
	"log"
	"os"
	"pet-pal/api/api"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrate() {
	err := DB.AutoMigrate(api.User{}, &api.Breed{}, &api.Food{}, &api.HealthEvent{}, &api.Image{}, &api.Meal{}, &api.Medication{}, &api.Medicine{}, &api.Pet{}, &api.Species{})
	if err != nil {
		log.Fatalf("An error occured while performing auto-migration: %v\n", err)
	}
}

func InitDb() {
	var dsn string = os.ExpandEnv("host=$DB_HOST user=$DB_USER password=$DB_PASS dbname=$DB_NAME port=$DB_PORT sslmode=disable TimeZone=America/New_York")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening connection to database %v\n", err)
	}
	DB = db
	migrate()
}
