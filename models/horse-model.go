package models

import (
	"fmt"
	"main/constants"
	"main/initializers"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Horse struct {
	gorm.Model
	Owner              uint
	Name               string
	Sex                string
	Coat               string
	Birthdate          time.Time
	Deathdate          *time.Time
	Speed              float32
	Acceleration       float32
	Stamina            float32
	Consistency        float32
	Versatility        float32
	Health             int
	SpeedTalent        string
	AccelerationTalent string
	StaminaTalent      string
	ConsistencyTalent  string
	VersatilityTalent  string
	BreedingPotential  int
	BreedingType       string
	Rating             int
	SourceHorse        bool
	StarterHorse       bool
	SireID             *uint
	Sire               *Horse
	DamID              *uint
	Dam                *Horse
}

func CreateSourceHorses() {
	//I need to create Source Horses in this function. And then save those horses to the Database.
	//I want them to be created with random Stats

	//Init empty slice to store new SourceHorses

	sourceHorses := []Horse{}

	for _, name := range constants.SourceHorseNames {
		sourceHorse := Horse{
			Name:               name,
			Sex:                "stallion",
			Coat:               constants.Coats[rand.Intn(len(constants.Coats))],
			Birthdate:          time.Now(),
			Speed:              rand.Float32() * 21,
			Acceleration:       rand.Float32() * 21,
			Stamina:            rand.Float32() * 21,
			Consistency:        rand.Float32() * 21,
			Versatility:        rand.Float32() * 21,
			Health:             100,
			SpeedTalent:        constants.TalentRarity[rand.Intn(len(constants.TalentRarity))],
			AccelerationTalent: constants.TalentRarity[rand.Intn(len(constants.TalentRarity))],
			StaminaTalent:      constants.TalentRarity[rand.Intn(len(constants.TalentRarity))],
			ConsistencyTalent:  constants.TalentRarity[rand.Intn(len(constants.TalentRarity))],
			VersatilityTalent:  constants.TalentRarity[rand.Intn(len(constants.TalentRarity))],
			BreedingPotential:  rand.Intn(100),
			BreedingType:       constants.BreedingType[rand.Intn(len(constants.BreedingType))],
			Rating:             0,
			SourceHorse:        true,
			StarterHorse:       false,
		}

		sourceHorses = append(sourceHorses, sourceHorse)

	}
	initializers.DB.Create(&sourceHorses)
}

func CreateStarterHorse() {
	selectedHorses, err := getRandomSourceHorses(4)

	if err != nil {
		fmt.Print(err)
	}

	for index, horse := range selectedHorses {
		fmt.Println(index+1, horse.Name)

	}

}

func getRandomSourceHorses(num int) ([]Horse, error) {
	var sourceHorses []Horse

	// Select four source horses at random.
	result := initializers.DB.Where("source_horse = ?", true).Find(&sourceHorses)

	if result.Error != nil {
		fmt.Println("Failed to find Source Horses", result.Error)
		return nil, result.Error
	}

	// Shuffle the slice of source horses
	rand.Shuffle(len(sourceHorses), func(i, j int) {
		sourceHorses[i], sourceHorses[j] = sourceHorses[j], sourceHorses[i]
	})

	selectedHorses := sourceHorses[:num]

	return selectedHorses, nil
}
