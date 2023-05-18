package education

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/pawn"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

const (
	NIL = iota
	CHAR_EDU
	CHAR_TRA
	CHAR_INS
)

const (
	CHAR_STRENGHT = iota
	CHAR_DEXTERITY
	CHAR_AGILITY
	CHAR_GRACE
	CHAR_ENDURANCE
	CHAR_STAMINA
	CHAR_VIGOR
	CHAR_INTELLIGENCE
	CHAR_EDUCATION
	CHAR_TRAINING
	CHAR_INSTINCT
	CHAR_SOCIAL
	CHAR_CHARISMA
	CHAR_CASTE
	CHAR_SANITY
	CHAR_PSIONICS
	C1
	C2
	C3
	C4
	C5
	C6
	AUTO
	BasicSchoolED5
	BasicSchoolApprentice
	BasicSchoolTrainingCourse
	BasicSchoolTradeSchool
	Colledge
	University
	LawSchool
	MedicalSchool
	MilitaryAcademy
	MilitaryCommandColledge
	NavalAcademy
	NavalCommandColledge
	ArmySchool
	NavalSchool
	MarineSchool
)

type educationalProcess struct {
	Character *pawn.Pawn
	BA        bool
	MA        bool
}

type institution struct {
	name              string
	preRequsite       string
	applyCheck        string
	duration          int //years
	validPassFailCHAR []string
	howManyRolls      int
	provides          string
	graduation        string
	form              string
	caa               []string
	majMinID          int
	/*
		ED5
		Trade School
			Colledge
				Honors
			University
				Honors
					Medical School
					Law School
				Masters
				Professors
			Service Academy
				ANM Schools
				Command Colledge



	*/
}

func newInstitution(id int) *institution {
	inst := institution{}
	switch id {
	case BasicSchoolED5:
		inst.form = "ED5"
		inst.preRequsite = "Edu 4 -"
		inst.applyCheck = "auto"
		inst.duration = 0
		inst.validPassFailCHAR = []string{"Int"}
		inst.howManyRolls = 1
		inst.provides = ""
		inst.graduation = "Edu=5"
		inst.caa = []string{"", "", ""}
	case BasicSchoolTradeSchool:
		inst.form = "Trade School"
		inst.preRequsite = "Edu 5 +"
		inst.applyCheck = "auto"
		inst.duration = 0
		inst.validPassFailCHAR = []string{"Int"}
		inst.howManyRolls = 1
		inst.provides = ""
		inst.graduation = "Edu=5"
		inst.caa = []string{"", "", ""}
	}
	return &inst
}

type Student interface {
	Profile() profile.Profile
	EducationState() (int, int, string)
	CheckCharacteristic(int, int) bool
}

type Institution interface {
	Form() string
}

func listMajorMinorSkillID(institutionID int) ([]int, error) {
	list := []int{}
	switch institutionID {
	default:
		for i := BasicSchoolED5; i <= MarineSchool; i++ {
			if i == institutionID {
				return []int{}, fmt.Errorf("unimplemented id %v", institutionID)
			}
		}
		return []int{}, fmt.Errorf("unknown id %v", institutionID)
	case BasicSchoolED5:
		list = []int{}
	case BasicSchoolTradeSchool, BasicSchoolApprentice, BasicSchoolTrainingCourse:
		list = []int{
			skill.ID_Admin,
			skill.ID_Comms,
			skill.ID_Computer,
			skill.ID_Explosives,
			skill.ID_High_G,
			skill.ID_Hostile_Environ,
			skill.ID_Language,
			skill.ID_Survey,
			skill.ID_Survival,
			skill.ID_Tactics,
			skill.ID_Trader,
			skill.ID_Vacc_Suit,
			skill.ID_Zero_G,
			skill.ID_Medic,
			skill.ID_Sensors,
			skill.ID_Steward,
			skill.ID_Forward_Observer,
			skill.ID_Navigator,
			skill.ID_Recon,
			skill.ID_Sapper,
			skill.ID_Actor,
			skill.ID_Artist,
			skill.ID_Author,
			skill.ID_Chef,
			skill.ID_Dancer,
			skill.ID_Musician,
			skill.ID_Biologics,
			skill.ID_Craftsman,
			skill.ID_Electronics,
			skill.ID_Fluidics,
			skill.ID_Gravitics,
			skill.ID_Magnetics,
			skill.ID_Mechanic,
			skill.ID_Photonics,
			skill.ID_Polymers,
			skill.ID_Programmer,
			skill.ID_ACV,
			skill.ID_Automotive,
			skill.ID_Grav_d,
			skill.ID_Legged,
			skill.ID_Mole,
			skill.ID_Tracked,
			skill.ID_Wheeled,
			skill.ID_Blades,
			skill.ID_Slugs,
			skill.ID_Unarmed,
			skill.ID_Jump,
			skill.ID_Life_Support,
			skill.ID_Maneuver,
			skill.ID_Power,
			skill.ID_Linguistics,
			skill.ID_Robotics,
			skill.ID_Aeronautics,
			skill.ID_Flappers,
			skill.ID_Grav_f,
			skill.ID_LTA,
			skill.ID_Rotor,
			skill.ID_Winged,
			skill.ID_Small_Craft,
			skill.ID_Rider,
			skill.ID_Teamster,
			skill.ID_Trainer,
			skill.ID_Aquanautics,
			skill.ID_Grav_s,
			skill.ID_Boat,
			skill.ID_Ship,
			skill.ID_Sub,
		}
	case Colledge, University:
		list = []int{
			skill.ID_Athlete,
			skill.ID_Broker,
			skill.ID_Bureaucrat,
			skill.ID_Counsellor,
			skill.ID_Designer,
			skill.ID_Language,
			skill.ID_Teacher,
			skill.ID_Astrogator,
			skill.ID_Actor,
			skill.ID_Artist,
			skill.ID_Author,
			skill.ID_Chef,
			skill.ID_Dancer,
			skill.ID_Musician,
			skill.ID_Biologics,
			skill.ID_Craftsman,
			skill.ID_Electronics,
			skill.ID_Fluidics,
			skill.ID_Gravitics,
			skill.ID_Magnetics,
			skill.ID_Mechanic,
			skill.ID_Photonics,
			skill.ID_Polymers,
			skill.ID_Programmer,
			skill.ID_Archeology,
			skill.ID_Biology,
			skill.ID_Chemistry,
			skill.ID_History,
			skill.ID_Linguistics,
			skill.ID_Philosophy,
			skill.ID_Physics,
			skill.ID_Planetology,
			skill.ID_Psionicology,
			skill.ID_Psyhohistory,
			skill.ID_Psyhology,
			skill.ID_Robotics,
			skill.ID_Sophontology,
			skill.ID_Aeronautics,
			skill.ID_Aquanautics,
		}
	case LawSchool:
		list = []int{
			skill.ID_Advocate,
			skill.ID_Diplomat,
		}
	case MedicalSchool:
		list = []int{
			skill.ID_Forensics,
			skill.ID_Medic,
		}
	case ArmySchool:
		list = []int{
			skill.ID_ACV,
			skill.ID_Automotive,
			skill.ID_Grav_d,
			skill.ID_Legged,
			skill.ID_Mole,
			skill.ID_Tracked,
			skill.ID_Wheeled,
			skill.ID_Battle_Dress,
			skill.ID_Beams,
			skill.ID_Blades,
			skill.ID_Exotics,
			skill.ID_Slugs,
			skill.ID_Sprays,
			skill.ID_Unarmed,
			skill.ID_Life_Support,
			skill.ID_Power,
			skill.ID_Robotics,
			skill.ID_Aeronautics,
			skill.ID_Flappers,
			skill.ID_Grav_f,
			skill.ID_LTA,
			skill.ID_Rotor,
			skill.ID_Winged,
			skill.ID_Screens,
			skill.ID_Small_Craft,
			skill.ID_Rider,
			skill.ID_Teamster,
			skill.ID_Trainer,
			skill.ID_Grav_s,
			skill.ID_Artilery,
			skill.ID_Launchers,
			skill.ID_Ordinance,
			skill.ID_WMD,
		}
	case NavalSchool:
		list = []int{
			skill.ID_Grav_d,
			skill.ID_Wheeled,
			skill.ID_Battle_Dress,
			skill.ID_Slugs,
			skill.ID_Jump,
			skill.ID_Life_Support,
			skill.ID_Maneuver,
			skill.ID_Power,
			skill.ID_Robotics,
			skill.ID_Aeronautics,
			skill.ID_Grav_f,
			skill.ID_Winged,
			skill.ID_Bay_Weapons,
			skill.ID_Ortilery,
			skill.ID_Screens,
			skill.ID_Spines,
			skill.ID_Turrets,
			skill.ID_Spacecraft_ACS,
			skill.ID_Spacecraft_BCS,
			skill.ID_Trainer,
			skill.ID_Grav_s,
			skill.ID_WMD,
		}
	case MarineSchool:
		list = []int{
			skill.ID_Grav_d,
			skill.ID_Tracked,
			skill.ID_Wheeled,
			skill.ID_Battle_Dress,
			skill.ID_Beams,
			skill.ID_Blades,
			skill.ID_Exotics,
			skill.ID_Slugs,
			skill.ID_Sprays,
			skill.ID_Unarmed,
			skill.ID_Power,
			skill.ID_Robotics,
			skill.ID_Grav_f,
			skill.ID_Turrets,
			skill.ID_Small_Craft,
			skill.ID_Grav_s,
			skill.ID_Boat,
			skill.ID_Ship,
			skill.ID_Sub,
			skill.ID_Artilery,
			skill.ID_Launchers,
			skill.ID_Ordinance,
			skill.ID_WMD,
		}
	case MilitaryAcademy, MilitaryCommandColledge:
		list = []int{
			skill.ID_Language,
			skill.ID_Leader,
			skill.ID_Liaison,
			skill.ID_Strategy,
			skill.ID_Tactics,
			skill.ID_Medic,
			skill.ID_Fighter,
			skill.ID_Forward_Observer,
			skill.ID_Heavy_Weapons,
			skill.ID_Navigator,
			skill.ID_Recon,
			skill.ID_Sapper,
		}
	case NavalAcademy, NavalCommandColledge:
		list = []int{
			skill.ID_Fleet_Tactics,
			skill.ID_Language,
			skill.ID_Leader,
			skill.ID_Naval_Architect,
			skill.ID_Strategy,
			skill.ID_Tactics,
			skill.ID_Astrogator,
			skill.ID_Gunnery,
			skill.ID_Medic,
			skill.ID_Sensors,
			skill.ID_Steward,
			skill.ID_Fighter,
			skill.ID_Forward_Observer,
			skill.ID_Heavy_Weapons,
			skill.ID_Navigator,
			skill.ID_Recon,
			skill.ID_Sapper,
		}
	}

	return list, nil
}
