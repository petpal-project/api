package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var dsn string = os.ExpandEnv("host=$DB_HOST user=$DB_USER password=$DB_PASS dbname=$DB_NAME port=$DB_PORT sslmode=disable TimeZone=America/New_York")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening connection to database %v\n", err)
	}
	DB = db
}
