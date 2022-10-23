package api

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint
	fname string
	lname string
}

type Pet struct {
	gorm.Model
	ID          uint
	user_ID     uint
	user        User
	name        string
	breeds      []Breed
	species     Species
	age         uint
	images      []Image
	meals       []Meal
	medications []Medication
}

type Species struct {
	gorm.Model
	ID            uint
	name          string
	binomial_name string
}

type Breed struct {
	gorm.Model
	ID          uint
	species_ID  uint
	species     Species
	name        string
	size        string
	height_m    float64
	height_f    float64
	weight_m    float64
	weight_f    float64
	coat        string
	coat_desc   string
	colors      []string
	colors_desc string
	energy      string
	activities  []string
}

type Meal struct {
	gorm.Model
	ID        uint
	pet_ID    uint
	frequency []time.Time
	food      []Food
}

type Food struct {
	gorm.Model
	ID          uint
	category    string
	targetAge   uint
	description string
}

type Medication struct {
	gorm.Model
	ID        uint
	pet_ID    uint
	frequency uint
	startDate time.Time
	endDate   time.Time
	medicine  Medicine
}

type Medicine struct {
	gorm.Model
	ID            uint
	medication_ID uint
	description   string
}

type HealthEvent struct {
	gorm.Model
	ID          uint
	pet_ID      uint
	name        string
	category    string
	description string
}

type Image struct {
	gorm.Model
	ID          uint
	S3url       string
	description string
}
