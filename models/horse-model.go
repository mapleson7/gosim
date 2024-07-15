package models

import (
	"time"
	"gorm.io/gorm"
	"math/rand"
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
	Health            int
	TalentType        string
	TalentRarity      string
	BreedingPotential int
	BreedingType      string
	Rating            int
	SourceHorse       bool
}

func (h Horse) CreateSourceHorses() {
	//I need to create Source Horses in this function. And then save those horses to the Database.
	//I want them to be created with random Stats and the names listed in a constants file of which I have saved in the github repo of the old sim.
	//I will reuse those names.
	sourceHorseNames := []string{
		"Iroquois",
		"Sysonby",
		"Henry Of Navarre",
		"Domino",
		"Hamburg",
		"Man Of War",
		"Colin",
		"Rock View",
		"Regret",
		"Zev",
		"Exterminator",
		"Sir Barton",
		"Sun Briar",
		"Sweep",
		"Billy Kelly",
		"Gallant Fox",
		"War Admiral",
		"Equipoise",
		"Twenty Grand",
		"Seabiscuit",
		"Whirlaway",
		"Alsab",
		"Citation",
		"Count Fleet",
		"Assault",
		"Stymie",
		"Pensive",
		"Nashua",
		"Swaps",
		"Sword Dancer",
		"Tom Fool",
		"Round Table",
		"Dr Fager",
		"Buckpasser",
		"Damascus",
		"Arts And Letters",
		"Affirmed",
		"Seattle Slew",
		"Spectacular Bid",
		"Alydar",
		"Steve Cauthen",
		"Ferdinand",
		"Alysheba",
		"Sunday Silence",
		"Easy Goer",
		"Cigar",
		"Bargain Blitz",
		"Protester",
		"Pheonix",
		"Cordlace",
		"Bart",
		"Night Sky",
		"Hole In One",
		"Anchor",
		"Red Tape",
		"Alkimos",
		"Gold Digger",
		"Volume One",
		"Harry The Fox",
		"Tambourine",
	}

	//Init empty slice to store new SourceHorses
	sourceHorses := []Horse{}
	
	for _, name := range sourceHorseNames {
		sourceHorse := Horse{
			Name              :name,
	Sex              : "stallion",
	Coat             : Coats[rand.Intn(0, )],
	Birthdate        : time.Now(),
	Speed            : ,
	Acceleration     : float32,
	Stamina          : float32,
	Consistency      : float32,
	Health           : int,
	TalentType       : string,
	TalentRarity     : string,
	BreedingPotential: int,
	BreedingType     : string,
	Rating           : int,
	SourceHorse      : bool,
		}
	}
}
