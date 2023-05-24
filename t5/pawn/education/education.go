package education

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/genetics"
	"github.com/Galdoba/TravellerTools/t5/pawn"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
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
	OTC
	NOTC
	Masters
	Proffessors
)

type educationalProcess struct {
	Character *pawn.Pawn
	BA        bool
	MA        bool
}

type preRequsite struct {
	baseCharID  int
	baseCharMin int
	baseCharMax int
	degree      string
}

type institution struct {
	ID                int
	name              string
	baseCharID        int
	baseCharMin       int
	baseCharMax       int
	applyCheck        []int
	duration          int //years
	validPassFailCHAR []int
	howManyRolls      int
	provides          map[int][]string
	graduationEdu     int
	graduationDegree  string
	haveHonors        bool
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

func (i *institution) Form() string {
	return i.form
}

func NewInstitution(id int) *institution {
	inst := institution{}
	inst.ID = id
	inst.provides = make(map[int][]string)
	switch id {
	default:
		panic(fmt.Sprintf("not implemented institution %v", id))
	case BasicSchoolED5:
		inst.form = "ED5"
		// inst.baseCharID = characteristic.CHAR_EDUCATION
		// inst.baseCharMin = 0
		// inst.baseCharMax = 4
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE}
		inst.howManyRolls = 1
		inst.graduationEdu = 5
		inst.caa = []string{"", "", ""}
	case BasicSchoolTradeSchool:
		inst.form = "Trade School"
		// inst.baseCharID = characteristic.CHAR_EDUCATION
		// inst.baseCharMin = 5
		// inst.baseCharMax = 999
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE}
		inst.duration = 1
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Mj", "Mj"}
		inst.caa = []string{"", "", ""}
	case Colledge:
		inst.form = "Colledge"
		// inst.baseCharID = characteristic.CHAR_EDUCATION
		// inst.baseCharMin = 5
		// inst.baseCharMax = 999
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationEdu = 8
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.haveHonors = true
		inst.caa = []string{"", "", ""}
	case University:
		inst.form = "University"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.provides[5] = []string{"Mj"}
		inst.provides[6] = []string{"Mj", "Mn"}
		inst.provides[7] = []string{"Mj"}
		inst.provides[8] = []string{"Mj", "Mn"}
		inst.haveHonors = true
		inst.graduationEdu = 9
		inst.caa = []string{"", "", ""}
	case NavalAcademy:
		inst.form = "Naval Academy"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.haveHonors = true
		inst.graduationEdu = 8
		inst.caa = []string{"", "", ""}
	case MilitaryAcademy:
		inst.form = "Military Academy"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.haveHonors = true
		inst.graduationEdu = 8
		inst.caa = []string{"", "", ""}
	case Masters:
		inst.form = "Masters Program"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 2
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 2
		inst.graduationDegree = "MA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.haveHonors = false
		inst.graduationEdu = 9
		inst.caa = []string{"", "", ""}
	case LawSchool:
		inst.form = "Law School"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 2
		inst.graduationDegree = "Attorney"
		inst.provides[1] = []string{"Advocate"}
		inst.provides[2] = []string{"Advocate"}
		inst.haveHonors = false
		inst.graduationEdu = 10
		inst.caa = []string{"", "", ""}
	case MedicalSchool:
		inst.form = "Medical School"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 4
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationDegree = "Doctor"
		inst.provides[1] = []string{"Medic"}
		inst.provides[2] = []string{"Medic"}
		inst.provides[3] = []string{"Medic"}
		inst.provides[4] = []string{"Medic"}
		inst.haveHonors = false
		inst.graduationEdu = 10
		inst.caa = []string{"", "", ""}
	case Proffessors:
		inst.form = "Professors Program"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 2
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 2
		inst.graduationDegree = "Professor"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.haveHonors = false
		inst.graduationEdu = 12
		inst.caa = []string{"", "", ""}

	}
	return &inst
}

type studyOutcome struct {
	gainedMajor     int
	gainedMinor     int
	yearsPassed     int
	waiversUsed     int
	degreeGained    string
	newEducationVal int
	skillsGained    []int
	events          []string
}

func Outcome(out studyOutcome) (int, int, int, int, string, int, []int, []string) {
	return out.gainedMajor, out.gainedMinor, out.yearsPassed, out.waiversUsed, out.degreeGained, out.newEducationVal, out.skillsGained, out.events
}

/*
Циклы образования:
Тренировка: Training, Apprenticeship, Mentor, None
Базовое: ED5, Trade School, None
Высшее: College, University, Service Academy, None
Дополнительное: Аспирантура, Адвокат, Медицина, None
Карьерное: Command Colledge, Flight School

*/

func Attend(student Student, institutionID int) studyOutcome {
	outcome := studyOutcome{}
	outcome.gainedMajor, outcome.gainedMinor, outcome.waiversUsed, _ = student.EducationState()
	institution := NewInstitution(institutionID)
	outcome.events = append(outcome.events, fmt.Sprintf("%v selected for education", institution.form))

	//ADMISSION
	if !applyTo(institution, student, &outcome) {
		return outcome
	}

	//STUDY
	passChar := maxValChar(student.Profile(), institution.validPassFailCHAR) //выбираем большую характеристику для учебы
	for i := 0; i < institution.howManyRolls; i++ {
		if !studyYearSucces(institution, student, &outcome, i, passChar) {
			return outcome
		}
	}
	// if institution.graduationEdu > 0 {
	// 	outcome.newEducationVal = institution.graduationEdu
	// }
	if institution.haveHonors == false {
		outcome.events = append(outcome.events, fmt.Sprintf("Character graduated from %v", institution.form))
		return outcome
	}
	honors := false
	if student.CheckCharacteristic(pawn.CheckAverage, passChar) {
		honors = true
	}
	if honors == false {
		outcome.events = append(outcome.events, fmt.Sprintf("Character fail to get Honors Degree"))
		if student.CheckCharacteristic(pawn.CheckAverage, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
			honors = true
		}
		outcome.waiversUsed++
		outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
	}
	outcome.degreeGained = institution.graduationDegree
	if honors {
		outcome.degreeGained = "Honors " + outcome.degreeGained
		outcome.skillsGained = append(outcome.skillsGained, outcome.gainedMajor)
	}
	if institution.graduationEdu > 0 {
		outcome.newEducationVal = institution.graduationEdu
	}
	///OTC, NOTC
	switch institutionID {
	case Colledge, University:
		volonteerOptions := []int{0, 1, 2, 3, OTC, NOTC}

		volonteer := student.ChooseOne(volonteerOptions)
		switch volonteer {
		case 0, 1:
		case OTC, NOTC:
			outcome.events = append(outcome.events, fmt.Sprintf("Character volonteer to OTC"))
			pass := false
			if student.CheckCharacteristic(pawn.CheckAverage, passChar) {
				outcome.events = append(outcome.events, fmt.Sprintf("Character fail to get to OTC"))
				if student.CheckCharacteristic(pawn.CheckAverage, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
					outcome.events = append(outcome.events, fmt.Sprintf("Waver Used"))
					pass = true
				}
			} else {
				outcome.events = append(outcome.events, fmt.Sprintf("OTC joined"))
				pass = true
			}
			if pass {
				switch volonteer {
				case OTC:
					outcome.skillsGained = append(outcome.skillsGained, skill.SolderSkill)
					outcome.degreeGained = "Army Officer1 " + outcome.degreeGained
					outcome.events = append(outcome.events, fmt.Sprintf("Army commision resiived"))
				case NOTC:
					outcome.skillsGained = append(outcome.skillsGained, skill.ShipSkill)
					switch student.ChooseOne([]int{0, 1}) {
					case 0:
						outcome.degreeGained = "Navy Officer1 " + outcome.degreeGained
						outcome.events = append(outcome.events, fmt.Sprintf("Navy commision resiived"))
					case 1:
						outcome.degreeGained = "Marine Officer1 " + outcome.degreeGained
						outcome.events = append(outcome.events, fmt.Sprintf("Marine commision resiived"))
					}

				}

			}

		}
		higherEdu := true
		for higherEdu {
			higherEducationOptions := []int{0, 1, 2}
			if strings.Contains(outcome.degreeGained, "BA") && institutionID == University {
				higherEducationOptions = append(higherEducationOptions, Masters, Masters)
			}
			if strings.Contains(outcome.degreeGained, "Honors BA") && institutionID == University {
				higherEducationOptions = append(higherEducationOptions, LawSchool, MedicalSchool)
			}
			if strings.Contains(outcome.degreeGained, "MA") {
				higherEducationOptions = append(higherEducationOptions, Proffessors, Proffessors)
			}
			higherEducationSelected := student.ChooseOne(higherEducationOptions)
			switch higherEducationSelected {
			default:
				higherEdu = false
			case Masters, LawSchool, MedicalSchool, Proffessors:
				higherEducationProgram := NewInstitution(higherEducationSelected)
				if applyTo(higherEducationProgram, student, &outcome) {
					for i := 0; i < higherEducationProgram.howManyRolls; i++ {
						fmt.Println("YEAR", i)
						if !studyYearSucces(higherEducationProgram, student, &outcome, i, passChar) {
							return outcome
						}
					}
					outcome.degreeGained = higherEducationProgram.graduationDegree
					outcome.newEducationVal = higherEducationProgram.graduationEdu
					edu := student.Profile().Data("C5")
					if edu.Value() < outcome.newEducationVal {
						student.Profile().Inject("C5", outcome.newEducationVal)
					} else {
						student.Profile().Inject("C5", edu.Value()+1)
					}
				}

			}
		}
		outcome.events = append(outcome.events, "End Education")
	case MilitaryAcademy:
		outcome.degreeGained = "Army Officer1 " + outcome.degreeGained
	case NavalAcademy:
		switch student.ChooseOne([]int{0, 1}) {
		case 0:
			outcome.degreeGained = "Navy Officer1 " + outcome.degreeGained
		case 1:
			outcome.degreeGained = "Marine Officer1 " + outcome.degreeGained
		}
	}
	flBranch := false
	if strings.Contains(outcome.degreeGained, "Officer1 Honors BA") {
		switch student.ChooseOne([]int{0, 1}) {
		case 0:
		case 1:
			flBranch = true
		}
	}
	if strings.Contains(outcome.degreeGained, "Officer1 BA") {
		switch student.CheckCharacteristic(pawn.CheckAverage, CHAR_SOCIAL, -outcome.waiversUsed) {
		case true:
			flBranch = true
		}
	}
	if flBranch {
		c2 := student.Profile().Data(genetics.KEY_GENE_PRF_2)
		c2CHAR := 0
		switch c2.Value() {
		case CHAR_DEXTERITY:
			c2CHAR = CHAR_DEXTERITY
		case CHAR_AGILITY:
			c2CHAR = CHAR_AGILITY
		case CHAR_GRACE:
			c2CHAR = CHAR_GRACE
		}
		switch student.CheckCharacteristic(pawn.CheckAverage, c2CHAR) {
		case true:
			outcome.skillsGained = append(outcome.skillsGained, skill.ID_Pilot, skill.ID_Pilot, skill.ID_Pilot)
			outcome.degreeGained = "Flight Branch " + outcome.degreeGained
			outcome.events = append(outcome.events, "Flight Branch joined")
		case false:
			outcome.events = append(outcome.events, "Flight Branch was not attended")
		}
	}

	outcome.events = append(outcome.events, fmt.Sprintf("Character graduated from %v with %v degree", institution.form, outcome.degreeGained))
	return outcome
}

func printSl(sl []string) {
	for _, s := range sl {
		fmt.Println(s)
	}
}

func (o *studyOutcome) String() string {
	str := o.degreeGained
	printSl(o.events)
	return str
}

func maxValChar(prf profile.Profile, chars []int) int {
	bestCode := 0
	bestVal := 0
	for _, chr := range chars {
		char := characteristic.FromProfile(prf, chr)
		if bestVal < char.Actual() {
			bestCode = chr
		}
	}
	return bestCode
}

type Student interface {
	Profile() profile.Profile
	EducationState() (int, int, int, string)
	CheckCharacteristic(int, int, ...int) bool
	SetMajorMinorWaiver(int, int, int)
	ChooseOne([]int) int
}

type Institution interface {
	Form() string
}

func listMajorMinorSkillID(institutionID int) []int {
	list := []int{}
	switch institutionID {
	default:
		panic(fmt.Sprintf("no list for %v", institutionID))
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
	case Colledge, University, Masters, Proffessors:
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
	if len(list) == 0 {
		list = append(list, skill.ID_NONE)
	}
	return list
}

func CurrentOptions(student Student) []int {
	prof := student.Profile()
	_, _, _, degree := student.EducationState()
	baseCharHex := prof.Data(genetics.KEY_GENE_PRF_5)
	if baseCharHex == nil {
		return []int{}
	}
	options := []int{}
	baseChar := characteristic.FromProfile(prof, baseCharHex.Value())
	fmt.Println(baseChar.Abb(), baseChar.Actual())
	if degree == "" {
		if baseChar.Abb() == "Edu" && baseChar.Actual() <= 4 {
			options = append(options, BasicSchoolED5)
		}
		if baseChar.Abb() == "Edu" && baseChar.Actual() >= 5 {
			options = append(options, BasicSchoolTradeSchool)
		}
		if baseChar.Abb() == "Edu" && baseChar.Actual() >= 5 {
			options = append(options, Colledge)
		}
		if baseChar.Abb() == "Edu" && baseChar.Actual() >= 6 {
			options = append(options, MilitaryAcademy)
		}
		if baseChar.Abb() == "Edu" && baseChar.Actual() >= 6 {
			options = append(options, NavalAcademy)
		}
		if baseChar.Abb() == "Edu" && baseChar.Actual() >= 7 {
			options = append(options, University)
		}
	}
	return options
}

func selectMjMn(institutionID int, student Student) (int, int) {
	gainedMajor := skill.ID_NONE
	gainedMinor := skill.ID_NONE
	switch institutionID {
	case BasicSchoolTradeSchool:
		gainedMajor = student.ChooseOne(listMajorMinorSkillID(institutionID))
	case Colledge, University, MilitaryAcademy, NavalAcademy, Masters, LawSchool, MedicalSchool, Proffessors:
		gainedMajor = student.ChooseOne(listMajorMinorSkillID(institutionID))
		gainedMinor = student.ChooseOne(listMajorMinorSkillID(institutionID))
		for gainedMajor == gainedMinor {
			gainedMajor = student.ChooseOne(listMajorMinorSkillID(institutionID))
			gainedMinor = student.ChooseOne(listMajorMinorSkillID(institutionID))
		}
	}
	return gainedMajor, gainedMinor
}

func applyTo(institution *institution, student Student, outcome *studyOutcome) bool {
	autoApply := false
	if len(institution.applyCheck) == 0 {
		autoApply = true
	}
	applyChar := maxValChar(student.Profile(), institution.applyCheck) //выбираем большую характеристику для поступления
	switch {
	case autoApply:
	default:
		switch student.CheckCharacteristic(pawn.CheckAverage, applyChar) {
		case true:
			outcome.events = append(outcome.events, fmt.Sprintf("Apply test Success"))
		case false:
			if !student.CheckCharacteristic(pawn.CheckAverage, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
				outcome.waiversUsed++ //если вейфер не прошел - тратим год и уходим
				outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
				outcome.yearsPassed = 1
				return false
			}
			outcome.waiversUsed++
			outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
		}
	}
	outcome.events = append(outcome.events, fmt.Sprintf("Character joined %v", institution.form))
	if outcome.gainedMajor == outcome.gainedMinor || student.ChooseOne([]int{0, 1}) == 1 {
		outcome.gainedMajor, outcome.gainedMinor = selectMjMn(institution.ID, student)
	}
	return true
}

func studyYearSucces(institution *institution, student Student, outcome *studyOutcome, year, passChar int) bool {
	outcome.yearsPassed++
	if !student.CheckCharacteristic(pawn.CheckAverage, passChar) {
		outcome.events = append(outcome.events, fmt.Sprintf("Year %v exams fail", year+1))
		if !student.CheckCharacteristic(pawn.CheckAverage, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
			outcome.waiversUsed++
			outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
			outcome.events = append(outcome.events, fmt.Sprintf("On year %v character was expelled from %v", year+1, institution.form))
			return false
		}
		outcome.waiversUsed++
		outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
	}
	outcome.events = append(outcome.events, fmt.Sprintf("Year %v exams pass", year+1))
	//gain Skill
	for _, gain := range institution.provides[year+1] {
		if gain == "" {
			panic("no skill gain provided")
		}
		switch gain {
		default:
			fmt.Println(outcome)
			panic("skill gain unknown")
		case "Mj":
			outcome.skillsGained = append(outcome.skillsGained, outcome.gainedMajor)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn %v", skill.NameByID(outcome.gainedMajor)))
		case "Mn":
			outcome.skillsGained = append(outcome.skillsGained, outcome.gainedMinor)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn %v", skill.NameByID(outcome.gainedMinor)))
		case "Advocate":
			outcome.skillsGained = append(outcome.skillsGained, skill.ID_Advocate)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn %v", skill.NameByID(skill.ID_Advocate)))
		case "Medic":
			outcome.skillsGained = append(outcome.skillsGained, skill.ID_Medic)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn %v", skill.NameByID(skill.ID_Medic)))
		}
	}
	return true
}
