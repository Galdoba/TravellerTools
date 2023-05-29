package education

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/profile"
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
	NavalAcademy
	ArmySchool
	MilitaryCommandColledge
	NavalSchool
	NavalCommandColledge
	MarineSchool
	MarineCommandColledge
	OTC
	NOTC
	Masters
	Proffessors
	Mentor
	FlightSchool
)

// type educationalProcess struct {
// 	Character *pawn.Pawn
// 	BA        bool
// 	MA        bool
// }

// type preRequsite struct {
// 	baseCharID  int
// 	baseCharMin int
// 	baseCharMax int
// 	degree      string
// }

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
	commision         string
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
	case BasicSchoolTradeSchool:
		inst.form = "Trade School"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE}
		inst.duration = 1
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Mj", "Mj"}
	case BasicSchoolApprentice:
		inst.form = "Aprenticeship"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_TRAINING}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Mj", "Mj", "Mj", "Mj"}
		inst.graduationDegree = "Aprentice"
	case Colledge:
		inst.form = "Colledge"
		// inst.baseCharID = characteristic.CHAR_EDUCATION
		// inst.baseCharMin = 5
		// inst.baseCharMax = 999
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 4
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationEdu = 8
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.haveHonors = true
	case University:
		inst.form = "University"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 4
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
	case NavalAcademy:
		inst.form = "Naval Academy"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 4
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.haveHonors = true
		inst.graduationEdu = 8
	case MilitaryAcademy:
		inst.form = "Military Academy"
		inst.applyCheck = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.duration = 4
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 4
		inst.graduationDegree = "BA"
		inst.provides[1] = []string{"Mj"}
		inst.provides[2] = []string{"Mj", "Mn"}
		inst.provides[3] = []string{"Mj"}
		inst.provides[4] = []string{"Mj", "Mn"}
		inst.haveHonors = true
		inst.graduationEdu = 8
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
	case OTC:
		inst.form = "Officer Training Corps"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 1
		inst.graduationDegree = "Army Officer1"
		inst.provides[1] = []string{"Solder Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case NOTC:
		inst.form = "Naval Officer Training Corps"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION}
		inst.howManyRolls = 1
		inst.graduationDegree = "Navy Officer1"
		inst.provides[1] = []string{"Ship Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case MilitaryCommandColledge:
		inst.form = "Military Command Colledge"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION, characteristic.CHAR_INSTINCT, characteristic.CHAR_TRAINING}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Army Skill", "Army Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case ArmySchool:
		inst.form = "Army School"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_DEXTERITY, characteristic.CHAR_AGILITY, characteristic.CHAR_GRACE, characteristic.CHAR_ENDURANCE, characteristic.CHAR_STAMINA, characteristic.CHAR_VIGOR}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Army Skill", "Army Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case NavalCommandColledge:
		inst.form = "Naval Command Colledge"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION, characteristic.CHAR_INSTINCT, characteristic.CHAR_TRAINING}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Naval Skill", "Naval Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case NavalSchool:
		inst.form = "Naval School"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_DEXTERITY, characteristic.CHAR_AGILITY, characteristic.CHAR_GRACE, characteristic.CHAR_ENDURANCE, characteristic.CHAR_STAMINA, characteristic.CHAR_VIGOR}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Naval Skill", "Naval Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case MarineCommandColledge:
		inst.form = "Marine Command Colledge"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_INTELLIGENCE, characteristic.CHAR_EDUCATION, characteristic.CHAR_INSTINCT, characteristic.CHAR_TRAINING}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Marine Skill", "Marine Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case MarineSchool:
		inst.form = "Marine School"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_DEXTERITY, characteristic.CHAR_AGILITY, characteristic.CHAR_GRACE, characteristic.CHAR_ENDURANCE, characteristic.CHAR_STAMINA, characteristic.CHAR_VIGOR}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Marine Skill", "Marine Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case BasicSchoolTrainingCourse:
		inst.form = "Training Course"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_TRAINING}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"School Skill"}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case Mentor:
		inst.form = "Mentor"
		inst.applyCheck = []int{}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_TRAINING, characteristic.CHAR_INTELLIGENCE}
		inst.howManyRolls = 1
		inst.provides[1] = []string{}
		inst.haveHonors = false
		inst.graduationEdu = 0
	case FlightSchool:
		inst.form = "Flight School"
		inst.applyCheck = []int{characteristic.CHAR_SOCIAL}
		inst.duration = 0
		inst.validPassFailCHAR = []int{characteristic.CHAR_DEXTERITY, characteristic.CHAR_AGILITY, characteristic.CHAR_GRACE}
		inst.howManyRolls = 1
		inst.provides[1] = []string{"Pilot", "Pilot", "Pilot"}
		inst.haveHonors = false
		inst.graduationEdu = 0

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

func Outcome(out studyOutcome) (int, int, int, int, string, []int, []string) {
	return out.gainedMajor, out.gainedMinor, out.yearsPassed, out.waiversUsed, out.degreeGained, out.skillsGained, out.events
}

/*
Циклы образования:
Тренировка: Training, Apprenticeship, Mentor, None
Базовое: ED5, Trade School, None
Высшее: College, University, Service Academy, None
Дополнительное: Аспирантура, Адвокат, Медицина, None
Карьерное: Command Colledge, Flight School

*/

func (outcome *studyOutcome) honors(institution *institution, student Student) {
	passChar := maxValChar(student.Profile(), institution.validPassFailCHAR) //выбираем большую характеристику для учебы
	if institution.haveHonors == false {
		return
	}
	honors := false
	if student.CheckCharacteristic(0, passChar) {
		honors = true
	}
	if honors == false {
		outcome.events = append(outcome.events, fmt.Sprintf("Character fail to get Honors Degree"))
		if student.CheckCharacteristic(0, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
			honors = true
		}
		outcome.waiversUsed++
		outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
	}
	//outcome.degreeGained = institution.graduationDegree
	if honors {
		outcome.degreeGained = strings.ReplaceAll(outcome.degreeGained, "BA", "Honors BA")
		outcome.skillsGained = append(outcome.skillsGained, outcome.gainedMajor)
	}
	return
}

func (outcome *studyOutcome) processEducationEvents(institution *institution, student Student) bool {
	passChar := maxValChar(student.Profile(), institution.validPassFailCHAR) //выбираем большую характеристику для учебы

	for i := 0; i < institution.howManyRolls; i++ {
		if !studyYearSucces(institution, student, outcome, i, passChar) {
			return false
		}
	}
	switch institution.ID {
	case Colledge, University:
		outcome.degreeGained = "BA"
	case MilitaryAcademy, OTC:
		outcome.degreeGained = "Army Officer1 BA"
	case NavalAcademy, NOTC:
		switch student.ChooseOne([]int{0, 1}) {
		case 0:
			outcome.degreeGained = "Navy Officer1 BA"
		case 1:
			outcome.degreeGained = "Marine Officer1 BA"
		}
	}
	if institution.graduationEdu > 0 {
		outcome.newEducationVal = institution.graduationEdu
	}
	outcome.honors(institution, student)
	return true
}

func Attend(student Student, institutionID int) studyOutcome {
	outcome := studyOutcome{}
	outcome.gainedMajor, outcome.gainedMinor, outcome.waiversUsed, _ = student.EducationState()
	institution := NewInstitution(institutionID)
	outcome.events = append(outcome.events, fmt.Sprintf("%v selected for education", institution.form))

	//ADMISSION
	if !addmissionSuccess(institution, student, &outcome) {
		return outcome
	}
	if !outcome.processEducationEvents(institution, student) {
		return outcome
	}

	///OTC, NOTC
	switch institutionID {
	case Colledge, University:
		volonteerOptions := []int{0, 1, 2, 3, OTC, NOTC}

		volonteer := student.ChooseOne(volonteerOptions)
		switch volonteer {
		case 0, 1:
		case OTC, NOTC:
			trainingCorps := NewInstitution(volonteer)
			switch addmissionSuccess(trainingCorps, student, &outcome) {
			case false:
			case true:
				outcome.processEducationEvents(trainingCorps, student)

			}
		}
	}
	//FLIGHT BRANCH IN CAREER
	// flBranch := false
	// if strings.Contains(outcome.degreeGained, "Officer1 Honors BA") {
	// 	switch student.ChooseOne([]int{0, 1}) {
	// 	case 0:
	// 	case 1:
	// 		flBranch = true
	// 	}
	// }
	// if strings.Contains(outcome.degreeGained, "Officer1 BA") {
	// 	switch student.CheckCharacteristic(0, CHAR_SOCIAL, -outcome.waiversUsed) {
	// 	case true:
	// 		flBranch = true
	// 	}
	// }
	// if flBranch {
	// 	c2 := student.Profile().Data(genetics.KEY_GENE_PRF_2)
	// 	c2CHAR := 0
	// 	switch c2.Value() {
	// 	case CHAR_DEXTERITY:
	// 		c2CHAR = CHAR_DEXTERITY
	// 	case CHAR_AGILITY:
	// 		c2CHAR = CHAR_AGILITY
	// 	case CHAR_GRACE:
	// 		c2CHAR = CHAR_GRACE
	// 	}
	// 	switch student.CheckCharacteristic(0, c2CHAR) {
	// 	case true:
	// 		outcome.skillsGained = append(outcome.skillsGained, skill.ID_Pilot, skill.ID_Pilot, skill.ID_Pilot)
	// 		outcome.degreeGained = "Flight Branch " + outcome.degreeGained
	// 		outcome.events = append(outcome.events, "Flight Branch joined")
	// 	case false:
	// 		outcome.events = append(outcome.events, "Flight Branch was not attended")
	// 	}
	// }

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
			bestVal = char.Actual()
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
	HasRequsite(string) bool
}

type Institution interface {
	Form() string
}

func listMajorMinorSkillID(institutionID int) []int {
	list := []int{}
	switch institutionID {
	default:
		panic(fmt.Sprintf("no list for %v", institutionID))
	case BasicSchoolED5, Mentor, FlightSchool:
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
	case ArmySchool, OTC:
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
	case NavalSchool, NOTC:
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
	case MarineCommandColledge, MarineSchool:
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
	return list
}

func CurrentOptions(student Student) []int {
	options := []int{}
	if student.HasRequsite("noED5") {
		if student.HasRequsite("Edu 5+") {
			options = append(options, BasicSchoolTradeSchool, Colledge)
		}
		if student.HasRequsite("Edu 6+") {
			options = append(options, MilitaryAcademy, NavalAcademy)
		}
		if student.HasRequsite("Edu 7+") {
			options = append(options, University)
		}
		if student.HasRequsite("C5=Tra") {
			options = append(options, Mentor)
		}
	}
	if student.HasRequsite("BA") {
		options = append(options, Masters)
	}
	if student.HasRequsite("Honors BA") {
		options = append(options, LawSchool, MedicalSchool)
	}
	if student.HasRequsite("MA") {
		options = append(options, Proffessors)
	}
	return options
}

func selectMjMn(institutionID int, student Student) (int, int) {
	gainedMajor := skill.ID_NONE
	gainedMinor := skill.ID_NONE
	for {
		gainedMajor = student.ChooseOne(listMajorMinorSkillID(institutionID))
		gainedMinor = student.ChooseOne(listMajorMinorSkillID(institutionID))
		if gainedMajor == gainedMinor || gainedMajor == skill.ID_NONE || gainedMinor == skill.ID_NONE {
			fmt.Println(gainedMajor, gainedMinor, institutionID)
			continue
		}
		break
	}
	return gainedMajor, gainedMinor
}

func addmissionSuccess(institution *institution, student Student, outcome *studyOutcome) bool {
	autoApply := false
	if len(institution.applyCheck) == 0 {
		autoApply = true
	}
	switch {
	case institution.ID == FlightSchool:
		_, _, _, degree := student.EducationState()
		if strings.Contains(degree, "Officer1 Honors BA") {
			autoApply = true
		} else {
			outcome.events = append(outcome.events, fmt.Sprintf("Addmission to %v rejected", institution.form))
			return false
		}
	}

	applyChar := maxValChar(student.Profile(), institution.applyCheck) //выбираем большую характеристику для поступления

	switch {
	case autoApply:
		outcome.events = append(outcome.events, fmt.Sprintf("Admission is automatic to %v", institution.form))
	default:
		switch student.CheckCharacteristic(0, applyChar) {
		case true:
			outcome.events = append(outcome.events, fmt.Sprintf("Apply test Success"))
		case false:
			if !student.CheckCharacteristic(0, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
				outcome.waiversUsed++ //если вейфер не прошел - тратим год и уходим
				outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
				switch institution.ID {
				case BasicSchoolTradeSchool, BasicSchoolTrainingCourse, Colledge, University, NavalAcademy, MilitaryAcademy,
					Masters, Proffessors, MedicalSchool, LawSchool, MilitaryCommandColledge, NavalCommandColledge:
					outcome.yearsPassed++
				}
				outcome.events = append(outcome.events, fmt.Sprintf("Addmission to %v rejected", institution.form))
				return false
			}
			outcome.waiversUsed++
			outcome.events = append(outcome.events, fmt.Sprintf("Waiver used %v times", outcome.waiversUsed))
		}
	}
	outcome.events = append(outcome.events, fmt.Sprintf("Character joined %v", institution.form))
	if len(listMajorMinorSkillID(institution.ID)) == 0 {
		return true
	}
	if outcome.gainedMajor == skill.ID_NONE || outcome.gainedMinor == skill.ID_NONE || student.ChooseOne([]int{0, 1}) == 1 {
		outcome.gainedMajor, outcome.gainedMinor = selectMjMn(institution.ID, student)
	}
	return true
}

func studyYearSucces(institution *institution, student Student, outcome *studyOutcome, year, passChar int) bool {
	if institution.duration > 0 {
		outcome.yearsPassed++
	}
	if !student.CheckCharacteristic(0, passChar) {
		outcome.events = append(outcome.events, fmt.Sprintf("Year %v exams fail", year+1))
		if !student.CheckCharacteristic(0, characteristic.CHAR_SOCIAL, -outcome.waiversUsed) {
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
		case "Solder Skill":
			outcome.skillsGained = append(outcome.skillsGained, skill.SolderSkill)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn Solder Skill %v", skill.NameByID(skill.SolderSkill)))
		case "Ship Skill":
			outcome.skillsGained = append(outcome.skillsGained, skill.ShipSkill)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn Ship Skill %v", skill.NameByID(skill.ShipSkill)))
		case "Army Skill":
			choise := student.ChooseOne(listMajorMinorSkillID(ArmySchool))
			outcome.skillsGained = append(outcome.skillsGained, choise)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn Army Skill %v", skill.NameByID(choise)))
		case "Naval Skill":
			choise := student.ChooseOne(listMajorMinorSkillID(NavalSchool))
			outcome.skillsGained = append(outcome.skillsGained, choise)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn Naval Skill %v", skill.NameByID(choise)))
		case "School Skill":
			choise := student.ChooseOne(listMajorMinorSkillID(BasicSchoolTradeSchool))
			outcome.skillsGained = append(outcome.skillsGained, choise)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn School Skill %v", skill.NameByID(choise)))
		case "Pilot":
			outcome.skillsGained = append(outcome.skillsGained, skill.ID_Pilot)
			outcome.events = append(outcome.events, fmt.Sprintf("Learn %v", skill.NameByID(skill.ID_Pilot)))
		}
	}
	return true
}
