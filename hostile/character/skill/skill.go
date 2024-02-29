package skill

import "fmt"

const (
	Undefined = iota
	Administration
	Agriculture
	Aircraft
	Brawling
	Blade_Combat
	Bribery
	Broker
	Carousing
	Comms
	Computer
	Demolitions
	Electronics
	Engineering
	Forgery
	Gambling
	Ground_Vechicle
	Gun_Combat
	Gunnery
	Heavy_Weapons
	Investigate
	Jack_of_All_Trades
	Leader
	Liason
	Loader
	Mechanical
	Medical
	Mining
	Navigation
	Pilot
	Recon
	Security
	Steward
	Streetwise
	Survival
	Tactics
	Vacc_Suit
	Watercraft
	Vechicle
)

type Skill struct {
	ID          int
	Name        string
	Description string
}

type SkillSet struct {
	skillVals map[int]int
}

func NewSkillSet() *SkillSet {
	sklst := SkillSet{}
	sklst.skillVals = make(map[int]int)
	return &sklst
}

func (ss *SkillSet) String() string {
	str := ""
	for i := Administration; i <= Watercraft; i++ {
		if v, ok := ss.skillVals[i]; ok {
			str += fmt.Sprintf("%v", v)
		} else {
			str += "-"
		}
	}
	return str
}

func (ss *SkillSet) AddBackGroundSkill(id int) error {
	skl := newSkill(id)
	if skl.Name == "error" {
		return fmt.Errorf("can't add background skill: %v", skl.Description)
	}
	if _, ok := ss.skillVals[id]; ok {
		return fmt.Errorf("can't add background skill: skill %v already exists in skillset", skl.Name)
	}
	ss.skillVals[id] = 0
	return nil
}

func (ss *SkillSet) Increase(id int) error {
	skl := newSkill(id)
	if skl.Name == "error" {
		return fmt.Errorf("can't add background skill: %v", skl.Description)
	}
	if val, ok := ss.skillVals[id]; ok {
		if val >= 5 {
			return fmt.Errorf("can't add background skill: skill %v already at level %v", skl.Name, val)
		}
		ss.skillVals[id] = val + 1
	} else {
		ss.skillVals[id] = 1
	}
	return nil
}

func (ss *SkillSet) SkillVal(id int) int {
	skl := newSkill(id)
	if skl.Name == "error" {
		return -99
	}
	if val, ok := ss.skillVals[id]; ok {
		return val
	}
	return -3
}

func newSkill(id int) Skill {
	skl := Skill{}
	skl.ID = id
	switch id {
	default:
		skl.Name = fmt.Sprintf("error")
		skl.Description = fmt.Sprintf("no skill with code %v", skl.Name)
	case Administration:
		skl.Name = "Administration"
		skl.Description = "Admin, paper-work, dealing with rules, regulations and officials, as well as law and legal personnel."
	case Agriculture:
		skl.Name = "Agriculture"
		skl.Description = "Growing and harvesting crops and raising animals, including hydroponics and aquaculture (on an ocean world)."
	case Aircraft:
		skl.Name = "Aircraft"
		skl.Description = "Operation and control of atmospheric craft, including transport planes, fast jets, jump jets, hovercars and tilt-rotors."
	case Brawling:
		skl.Name = "Brawling"
		skl.Description = "Fighting unarmed with fists, feet or using wrestling or grappling moves. It includes the use of blunt weapons such as clubs."
	case Blade_Combat:
		skl.Name = "Blade_Combat"
		skl.Description = "Fighting with knives or other edged weapons."
	case Bribery:
		skl.Name = "Bribery"
		skl.Description = "Offering bribes to circumvent local law, or to influence someone’s decision. The cash bribe must be appropriate to the situation."
	case Broker:
		skl.Name = "Broker"
		skl.Description = "Locating suppliers and buyers, and facilitating the purchase and resale of commercial goods, haggling, bartering and so forth."
	case Carousing:
		skl.Name = "Carousing"
		skl.Description = "Social skills, including picking up gossip or rumour, making friends and reading people’s body language."
	case Comms:
		skl.Name = "Comms"
		skl.Description = "Operating drones, sensors and radio equipment. Skilled characters can boost an incoming or outgoing signal, create or break a secure channel, detect signals and anomalies, hide or piggyback on another signal, jam local communications, locate and assess potential threats, and analyse complex sensor data."
	case Computer:
		skl.Name = "Computer"
		skl.Description = "Operation and programming of computers, including creating or breaking data encryption; mining data effectively; creating or breaking data and network security protocols; overriding a computer protocol, as well as other general programming tasks."
	case Demolitions:
		skl.Name = "Demolitions"
		skl.Description = "Using demolition charges and other explosive devices, including assembling or disarming bombs."
	case Electronics:
		skl.Name = "Electronics"
		skl.Description = "Operating and repairing complex electronic devices."
	case Engineering:
		skl.Name = "Engineering"
		skl.Description = "Use and maintenance of powerplants, reactors and starship drives."
	case Forgery:
		skl.Name = "Forgery"
		skl.Description = "Faking documents, currencies, and identification badges in order to deceive officials, government agents and security forces."
	case Gambling:
		skl.Name = "Gambling"
		skl.Description = "Running games, winning games and making money! If he has the opportunity, the character can organise games for others to play which will net him a nice cash win whilst still letting other gamblers also take away money."
	case Ground_Vechicle:
		skl.Name = "Ground Vechicle"
		skl.Description = "Operation and control of wheeled and tracked vehicles, including ATVs and Armoured Fighting Vehicles."
	case Gun_Combat:
		skl.Name = "Gun Combat"
		skl.Description = "Using and maintaining small arms including pistols and rifles, SMGs, lasers, machineguns and shotguns."
	case Gunnery:
		skl.Name = "Gunnery"
		skl.Description = "Using starship-mounted weaponry."
	case Heavy_Weapons:
		skl.Name = "Heavy Weapons"
		skl.Description = "Using military support weapons including rocket launchers, heavy artillery guns, tactical missile launchers, grenade launchers and high-velocity tank guns."
	case Investigate:
		skl.Name = "Investigate"
		skl.Description = "Scientific analysis and the use of complex and accurate scientific tools and equipment to gather clues at a crime scene or scientific location."
	case Jack_of_All_Trades:
		skl.Name = "Jack_of_All_Trades"
		skl.Description = "The character is a quick learner, able to pick up skills by watching and learning. He can help any skilled character, making his JoT roll to provide a +1 bonus to the skilled character’s attempt even though he doesn’t possess the required skill. The benefit of JoT is that the character with the skill can assist anyone, carrying out almost any careful activity, as long as he can be directed and instructed. JoT does not incur the standard -3 penalty for attempting a task while unskilled."
	case Leader:
		skl.Name = "Leader"
		skl.Description = "Motivating others in times of crisis or stress, particularly Non-Player Characters who might be reluctant to carry out an action."
	case Liason:
		skl.Name = "Liason"
		skl.Description = "The art of negotiation and diplomacy. Opposed checks can resolve cases when two diplomats are engaged in negotiations."
	case Loader:
		skl.Name = "Loader"
		skl.Description = "Operation and control of all cargo loading devices, everything from fork-lift trucks to cranes, augmented exoframe loaders and all forms of starship loading equipment."
	case Mechanical:
		skl.Name = "Mechanical"
		skl.Description = "Operating and repairing mechanical devices, from truck engines to airlock doors, hydraulic lifters to life support machinery."
	case Medical:
		skl.Name = "Medical"
		skl.Description = "Training and skill in the medical arts and sciences, from diagnosis and triage to surgery and other corrective treatments. This skill represents a character's ability to provide emergency care, short term care, long-term care, and specialized treatment for diseases, poisons and debilitating injuries."
	case Mining:
		skl.Name = "Mining"
		skl.Description = "The character has experience and training in prospecting and mining, both on a world surface and in a zero-G environment. It includes the use of mining equipment and machinery."
	case Navigation:
		skl.Name = "Navigation"
		skl.Description = "Plotting courses on planets, and in space and using instruments, maps and beacons to determine exact location."
	case Pilot:
		skl.Name = "Pilot"
		skl.Description = "Operation and control of interplanetary and interstellar spacecraft."
	case Recon:
		skl.Name = "Recon"
		skl.Description = "Scouting out dangers and spotting threats, ambushes, booby traps, unusual objects or out of place people. Characters are also adept at silent movement and in camouflage techniques."
	case Security:
		skl.Name = "Security"
		skl.Description = "Installing and also bypassing or dismantling security measures, from mechanical locks to swipe-card locks, keypad locks, surveillance cameras and various types of alarms and their triggers."
	case Steward:
		skl.Name = "Steward"
		skl.Description = "The care and serving of passengers and other guests and including customer service."
	case Streetwise:
		skl.Name = "Streetwise"
		skl.Description = "Familiarity with underworld society, its rules, personalities, groups and customs."
	case Survival:
		skl.Name = "Survival"
		skl.Description = "Staying alive in the wild, including hunting or trapping animals, avoiding exposure, locating sources of food and fresh water, producing fires, finding shelter, avoiding dangerous flora and fauna and dealing with the dangers of hazardous climates (arctic, desert, etc.)."
	case Tactics:
		skl.Name = "Tactics"
		skl.Description = "This skill covers tactical military planning and decision making, as well as the calling up of fire support from friendly artillery, aircraft and even ships in orbit."
	case Vacc_Suit:
		skl.Name = "Vacc_Suit"
		skl.Description = "Use and training in all types of vacuum suits, hostile environment suits and combat armour, as well as operating in a zero-G environment."
	case Watercraft:
		skl.Name = "Watercraft"
		skl.Description = "Operation and control of ocean or water-going craft including motorboats, hydrofoils, hovercraft, submarines and large commercial ships."
	case Vechicle:
		skl.Name = "Vechicle (Cascade Skill)"
		skl.Description = "The various specialties of this skill cover different types of planetary transportation. When this skill is received, the character must immediately select one of the following: Aircraft, Ground Vehicle or Watercraft."
	}
	return skl
}
