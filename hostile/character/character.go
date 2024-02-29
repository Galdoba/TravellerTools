package character

import (
	"github.com/Galdoba/TravellerTools/hostile/character/characteristic"
	"github.com/Galdoba/TravellerTools/hostile/character/skill"
)

type Character struct {
	Name           string
	PC             bool
	Homeworld      string
	Age            int
	Career         career.CareerPath
	Characteristic characteristic.CharSet
	Skill          skill.SkillSet
}
