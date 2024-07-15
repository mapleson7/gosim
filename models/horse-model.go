package models

import (
	"time"

	"gorm.io/gorm"
)

type Horse struct {
	
	gorm.Model
	Name              string
	Sex               string
	Coat              string
	Birthdate         time.Time
	Deathdate         *time.Time
	Speed             float32
	Acceleration      float32
	Stamina           float32
	Consistency       float32
	Health           int
	TalentType        string
	TalentRarity      string
	BreedingPotential int
	BreedingType      string
	Rating            int
}
