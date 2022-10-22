package api

import "time"

type User struct {
	ID    uint
	fname string
	lname string
}

type Pet struct {
	ID         uint
	name       string
	breed_ID   uint
	species_ID uint
	user_ID    uint
	age        uint
	image      uint
}

type Breed struct {
	ID          uint
	species_ID  uint
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
	ID        uint
	pet_ID    uint
	frequency []time.Time
	food      uint
}

type Food struct {
	ID          uint
	category    string
	targetAge   uint
	description string
}

type Medication struct {
	ID        uint
	pet_ID    uint
	frequency []time.Time
	medicine  uint
}

type Medicine struct {
	ID            uint
	medication_ID uint
	description   string
}

type HealthEvent struct {
	ID          uint
	pet_ID      uint
	name        string
	category    string
	description string
}

type Image struct {
	ID          uint
	S3url       string
	description string
}
