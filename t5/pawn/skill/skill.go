package skill

import "fmt"

const (
	ID_NONE = iota
	ID_Actor
	ID_Admin
	ID_Advocate
	ID_Animals
	ID_Rider
	ID_Teamster
	ID_Trainer
	ID_Artist
	ID_Astrogator
	ID_Athlete
	ID_Author
	ID_Biologics
	ID_Broker
	ID_Bureaucrat
	ID_Chef
	ID_Comms
	ID_Computer
	ID_Counsellor
	ID_Craftsman
	ID_Dancer
	ID_Designer
	ID_Diplomat
	ID_Driver
	ID_ACV
	ID_Automotive
	ID_Grav_d
	ID_Legged
	ID_Mole
	ID_Tracked
	ID_Wheeled
	ID_Electronics
	ID_Engineer
	ID_Jump
	ID_Life_Support
	ID_Maneuver
	ID_Power
	ID_Explosives
	ID_Fighter
	ID_Battle_Dress
	ID_Beams
	ID_Blades
	ID_Exotics
	ID_Slugs
	ID_Sprays
	ID_Unarmed
	ID_Fleet_Tactics
	ID_Fluidics
	ID_Flyer
	ID_Aeronautics
	ID_Flappers
	ID_Grav_f
	ID_LTA
	ID_Rotor
	ID_Winged
	ID_Forensics
	ID_Forward_Observer
	ID_Gambler
	ID_Gravitics
	ID_Gunnery
	ID_Bay_Weapons
	ID_Ortilery
	ID_Screens
	ID_Spines
	ID_Turrets
	ID_Heavy_Weapons
	ID_Artilery
	ID_Launchers
	ID_Ordinance
	ID_WMD
	ID_Language
	ID_High_G
	ID_Hostile_Environ
	ID_JOT
	ID_Leader
	ID_Liaison
	ID_Magnetics
	ID_Mechanic
	ID_Medic
	ID_Musician
	ID_Instrument_Guitar
	ID_Instrument_Banjo
	ID_Instrument_Mandolin
	ID_Instrument_Keyboard
	ID_Instrument_Piano
	ID_Instrument_Voice
	ID_Instrument_Trumpet
	ID_Instrument_Trombone
	ID_Instrument_Tuba
	ID_Instrument_Violin
	ID_Instrument_Viola
	ID_Instrument_Cello
	ID_Naval_Architect
	ID_Navigator
	ID_Photonics
	ID_Pilot
	ID_Small_Craft
	ID_Spacecraft_ABS
	ID_Spacecraft_BCS
	ID_Polymers
	ID_Programmer
	ID_Recon
	ID_Sapper
	ID_Seafarer
	ID_Aquanautics
	ID_Grav_s
	ID_Boat
	ID_Ship
	ID_Sub
	ID_Sensors
	ID_Stealth
	ID_Steward
	ID_Strategy
	ID_Streetwise
	ID_Survey
	ID_Survival
	ID_Tactics
	ID_Teacher
	ID_Trader
	ID_Vacc_Suit
	ID_Zero_G
	ID_Biology
	ID_Chemistry
	ID_Physics
	ID_Planetology
	ID_Robotics
	ID_Archeology
	ID_History
	ID_Linguistics
	ID_Philosophy
	ID_Psionicology
	ID_Psyhohistory
	ID_Psyhology
	ID_Sophontology
	ID_Compute
	ID_Empath
	ID_Hibernate
	ID_Hypno
	ID_Intuition
	ID_Math
	ID_Memaware
	ID_Memorize
	ID_Mempercept
	ID_Memscent
	ID_Memsight
	ID_Memsound
	ID_Morph
	ID_Rage
	ID_Soundmimic
	ID_Kkree
	ID_Anglic
	ID_Battle
	ID_Flash
	ID_Gonk
	ID_Gvegh
	ID_Mariel
	ID_Oynprith
	ID_Sagamaal
	ID_Tezapet
	ID_Trokh
	ID_Vilani
	ID_Zdetl
	ID_END
	One_Trade
	One_Art
	SG_GENERAL      = "General"
	SG_STARSHIP     = "Starship skill"
	SG_TRADE        = "Trade"
	SG_ARTS         = "Art"
	SG_SOLDIER      = "Soldier skill"
	SG_SCIENCE_HARD = "Hard Science"
	SG_SCIENCE_SOFT = "Soft Science"
	SG_SPECIALIZED  = "Specialized"
	SG_PERSONAL     = "Personal"
	TYPE_SKILL      = "Skill"
	TYPE_KNOWLEDGE  = "Knowledge"
	TYPE_TALENT     = "Talent"
)

type Skill struct {
	sklType             string
	Name                string
	id                  int
	ParentSkl           int
	AssociatedKnowledge []int
	related             []int
	group               string
	Default             bool
	KKSrule             bool
	value               int
}

func New(id int) (*Skill, error) {
	skl := Skill{}
	skl.id = id
	skl.Name = NameByID(id)
	switch id {
	default:
		return nil, fmt.Errorf("skill.New(): can not create skill with id '%v'", id)
	case ID_Admin:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Bureaucrat, ID_Leader}
	case ID_Advocate:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Counsellor}
	case ID_Animals:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.AssociatedKnowledge = []int{ID_Rider, ID_Teamster, ID_Trainer}
		skl.KKSrule = true
	case ID_Athlete:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.Default = true
	case ID_Broker:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Trader}
	case ID_Bureaucrat:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Admin, ID_Leader}
	case ID_Comms:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Computer, ID_Programmer}
		skl.Default = true
	case ID_Computer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Comms, ID_Programmer}
		skl.Default = true
	case ID_Counsellor:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Advocate}
	case ID_Designer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Craftsman}
	case ID_Diplomat:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Liaison}
	case ID_Driver:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.AssociatedKnowledge = []int{ID_ACV, ID_Legged, ID_Mole, ID_Tracked, ID_Wheeled, ID_Grav_d}
		skl.Default = true
		skl.KKSrule = true
	case ID_Explosives:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Fleet_Tactics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Strategy, ID_Tactics}
	case ID_Flyer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.AssociatedKnowledge = []int{ID_Flappers, ID_LTA, ID_Rotor, ID_Winged, ID_Grav_f, ID_Aeronautics}
		skl.related = []int{ID_Pilot}
		skl.KKSrule = true
	case ID_Forensics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Gambler:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_High_G:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Hostile_Environ, ID_Zero_G}
	case ID_Hostile_Environ:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_High_G, ID_Zero_G}
	case ID_JOT:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Language:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Linguistics}
		skl.AssociatedKnowledge = []int{ID_Kkree, ID_Anglic, ID_Battle, ID_Flash, ID_Gonk, ID_Gvegh, ID_Mariel, ID_Oynprith, ID_Sagamaal, ID_Tezapet, ID_Trokh, ID_Vilani, ID_Zdetl}
	case ID_Leader:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Admin, ID_Bureaucrat}
	case ID_Liaison:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Diplomat}
	case ID_Naval_Architect:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Seafarer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.AssociatedKnowledge = []int{ID_Aquanautics, ID_Grav_s, ID_Boat, ID_Ship, ID_Sub}
		skl.KKSrule = true
	case ID_Stealth:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Strategy:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Tactics, ID_Fleet_Tactics}
	case ID_Streetwise:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Survey:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Navigator, ID_Astrogator}
	case ID_Survival:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Tactics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Strategy, ID_Fleet_Tactics}
	case ID_Teacher:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Trainer}
	case ID_Trader:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.related = []int{ID_Broker}
	case ID_Vacc_Suit:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
		skl.Default = true
	case ID_Zero_G:
		skl.sklType = TYPE_SKILL
		skl.group = SG_GENERAL
	case ID_Astrogator:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
		skl.related = []int{ID_Navigator, ID_Survey}
	case ID_Engineer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
		skl.AssociatedKnowledge = []int{ID_Jump, ID_Life_Support, ID_Maneuver, ID_Power}
		skl.KKSrule = true
	case ID_Gunnery:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
		skl.AssociatedKnowledge = []int{ID_Bay_Weapons, ID_Ortilery, ID_Screens, ID_Spines, ID_Turrets}
		skl.related = []int{ID_Fighter, ID_Heavy_Weapons}
		skl.KKSrule = true
	case ID_Medic:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
	case ID_Pilot:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
		skl.AssociatedKnowledge = []int{ID_Small_Craft, ID_Spacecraft_ABS, ID_Spacecraft_BCS}
		skl.KKSrule = true
	case ID_Sensors:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
	case ID_Steward:
		skl.sklType = TYPE_SKILL
		skl.group = SG_STARSHIP
		skl.Default = true
	case ID_Biologics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
		skl.related = []int{ID_Biology}
	case ID_Craftsman:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
		skl.related = []int{ID_Designer}
	case ID_Electronics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
	case ID_Fluidics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
	case ID_Gravitics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
	case ID_Magnetics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
	case ID_Mechanic:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
		skl.Default = true
	case ID_Photonics:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
	case ID_Polymers:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
		skl.related = []int{ID_Chemistry}
	case ID_Programmer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_TRADE
		skl.related = []int{ID_Comms, ID_Computer}
	case ID_Actor:
		skl.sklType = TYPE_SKILL
		skl.group = SG_ARTS
		skl.related = []int{ID_Artist, ID_Author, ID_Chef, ID_Dancer, ID_Musician}
		skl.Default = true
	case ID_Artist:
		skl.sklType = TYPE_SKILL
		skl.group = SG_ARTS
		skl.related = []int{ID_Actor, ID_Author, ID_Chef, ID_Dancer, ID_Musician}
		skl.Default = true
	case ID_Author:
		skl.sklType = TYPE_SKILL
		skl.group = SG_ARTS
		skl.related = []int{ID_Actor, ID_Artist, ID_Chef, ID_Dancer, ID_Musician}
		skl.Default = true
	case ID_Chef:
		skl.sklType = TYPE_SKILL
		skl.group = SG_ARTS
		skl.related = []int{ID_Actor, ID_Artist, ID_Author, ID_Dancer, ID_Musician}
	case ID_Dancer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_ARTS
		skl.related = []int{ID_Actor, ID_Artist, ID_Author, ID_Chef, ID_Musician}
	case ID_Musician:
		skl.sklType = TYPE_SKILL
		skl.group = SG_ARTS
		skl.related = []int{ID_Actor, ID_Artist, ID_Author, ID_Chef, ID_Dancer}
		skl.AssociatedKnowledge = []int{ID_Instrument_Guitar, ID_Instrument_Banjo, ID_Instrument_Mandolin, ID_Instrument_Keyboard, ID_Instrument_Piano, ID_Instrument_Voice, ID_Instrument_Trumpet, ID_Instrument_Trombone, ID_Instrument_Tuba, ID_Instrument_Violin, ID_Instrument_Viola, ID_Instrument_Cello}
		skl.KKSrule = true
	case ID_Fighter:
		skl.sklType = TYPE_SKILL
		skl.group = SG_SOLDIER
		skl.AssociatedKnowledge = []int{ID_Battle_Dress, ID_Beams, ID_Blades, ID_Exotics, ID_Slugs, ID_Sprays, ID_Unarmed}
		skl.related = []int{ID_Heavy_Weapons, ID_Gunnery}
		skl.Default = true
		skl.KKSrule = true
	case ID_Forward_Observer:
		skl.sklType = TYPE_SKILL
		skl.group = SG_SOLDIER
	case ID_Heavy_Weapons:
		skl.sklType = TYPE_SKILL
		skl.group = SG_SOLDIER
		skl.AssociatedKnowledge = []int{ID_Artilery, ID_Launchers, ID_Ordinance, ID_WMD}
		skl.related = []int{ID_Fighter, ID_Gunnery}
		skl.KKSrule = true
	case ID_Navigator:
		skl.sklType = TYPE_SKILL
		skl.group = SG_SOLDIER
		skl.related = []int{ID_Astrogator, ID_Survey}
	case ID_Recon:
		skl.sklType = TYPE_SKILL
		skl.group = SG_SOLDIER
	case ID_Sapper:
		skl.sklType = TYPE_SKILL
		skl.group = SG_SOLDIER
	case ID_Bay_Weapons:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Gunnery
	case ID_Ortilery:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Gunnery
	case ID_Screens:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Gunnery
	case ID_Spines:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Gunnery
	case ID_Turrets:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Gunnery
		skl.Default = true
	case ID_Artilery:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Heavy_Weapons
	case ID_Launchers:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Heavy_Weapons
	case ID_Ordinance:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Heavy_Weapons
	case ID_WMD:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Heavy_Weapons
	case ID_Battle_Dress:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Beams:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Blades:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Exotics:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Slugs:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Sprays:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Unarmed:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Fighter
	case ID_Aeronautics:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Flyer
	case ID_Flappers:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Flyer
	case ID_Grav_f:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Flyer
	case ID_LTA:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Flyer
	case ID_Rotor:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Flyer
	case ID_Winged:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Flyer
	case ID_ACV:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Automotive:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Grav_d:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Legged:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Mole:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Tracked:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Wheeled:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Driver
	case ID_Jump:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Engineer
	case ID_Life_Support:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Engineer
	case ID_Maneuver:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Engineer
	case ID_Power:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Engineer
	case ID_Rider:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Animals
	case ID_Teamster:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Animals
	case ID_Trainer:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Animals
		skl.related = []int{ID_Teacher}
	case ID_Aquanautics:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Seafarer
	case ID_Grav_s:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Seafarer
	case ID_Boat:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Seafarer
	case ID_Ship:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Seafarer
	case ID_Sub:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Seafarer
	case ID_Small_Craft:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Pilot
	case ID_Spacecraft_ABS:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Pilot
	case ID_Spacecraft_BCS:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Pilot
	case ID_Instrument_Guitar:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Banjo:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Mandolin:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Keyboard:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Piano:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Voice:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Trumpet:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Trombone:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Tuba:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Violin:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Viola:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Instrument_Cello:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Musician
	case ID_Kkree:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Anglic:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Battle:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Flash:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Gonk:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Gvegh:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Mariel:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Oynprith:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Sagamaal:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Tezapet:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Trokh:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Vilani:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Zdetl:
		skl.sklType = TYPE_KNOWLEDGE
		skl.ParentSkl = ID_Language
	case ID_Biology:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_HARD
		skl.related = []int{ID_Biologics}
	case ID_Chemistry:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_HARD
	case ID_Physics:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_HARD
	case ID_Planetology:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_HARD
	case ID_Robotics:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_HARD
	case ID_Archeology:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_History:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_Linguistics:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
		skl.related = []int{ID_Language}
	case ID_Philosophy:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_Psionicology:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_Psyhohistory:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_Psyhology:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_Sophontology:
		skl.sklType = TYPE_KNOWLEDGE
		skl.group = SG_SCIENCE_SOFT
	case ID_Compute:
		skl.sklType = TYPE_TALENT
	case ID_Empath:
		skl.sklType = TYPE_TALENT
	case ID_Hibernate:
		skl.sklType = TYPE_TALENT
	case ID_Hypno:
		skl.sklType = TYPE_TALENT
	case ID_Intuition:
		skl.sklType = TYPE_TALENT
	case ID_Math:
		skl.sklType = TYPE_TALENT
	case ID_Memaware:
		skl.sklType = TYPE_TALENT
	case ID_Memorize:
		skl.sklType = TYPE_TALENT
	case ID_Mempercept:
		skl.sklType = TYPE_TALENT
	case ID_Memscent:
		skl.sklType = TYPE_TALENT
	case ID_Memsight:
		skl.sklType = TYPE_TALENT
	case ID_Memsound:
		skl.sklType = TYPE_TALENT
	case ID_Morph:
		skl.sklType = TYPE_TALENT
	case ID_Rage:
		skl.sklType = TYPE_TALENT
	case ID_Soundmimic:
		skl.sklType = TYPE_TALENT
	}
	//fmt.Println(skl)
	return &skl, nil
}

func NameByID(id int) string {
	switch id {
	default:
		return "UNDEFINED"
	case ID_NONE:
		return "NONE"
	case ID_Admin:
		return "Admin"
	case ID_Advocate:
		return "Advocate"
	case ID_Animals:
		return "Animals"
	case ID_Athlete:
		return "Athlete"
	case ID_Broker:
		return "Broker"
	case ID_Bureaucrat:
		return "Bureaucrat"
	case ID_Comms:
		return "Comms"
	case ID_Computer:
		return "Computer"
	case ID_Counsellor:
		return "Counsellor"
	case ID_Designer:
		return "Designer"
	case ID_Diplomat:
		return "Diplomat"
	case ID_Driver:
		return "Driver"
	case ID_Explosives:
		return "Explosives"
	case ID_Fleet_Tactics:
		return "Fleet Tactics"
	case ID_Flyer:
		return "Flyer"
	case ID_Forensics:
		return "Forensics"
	case ID_Gambler:
		return "Gambler"
	case ID_High_G:
		return "High-G"
	case ID_Hostile_Environ:
		return "Hostile Environ"
	case ID_JOT:
		return "Jack Of All Trades"
	case ID_Language:
		return "Language"
	case ID_Leader:
		return "Leader"
	case ID_Liaison:
		return "Liaison"
	case ID_Naval_Architect:
		return "Naval Architect"
	case ID_Seafarer:
		return "Seafarer"
	case ID_Stealth:
		return "Stealth"
	case ID_Strategy:
		return "Strategy"
	case ID_Streetwise:
		return "Streetwise"
	case ID_Survey:
		return "Survey"
	case ID_Survival:
		return "Survival"
	case ID_Tactics:
		return "Tactics"
	case ID_Teacher:
		return "Teacher"
	case ID_Trader:
		return "Trader"
	case ID_Vacc_Suit:
		return "Vacc Suit"
	case ID_Zero_G:
		return "Zero-G"
	case ID_Astrogator:
		return "Astrogator"
	case ID_Engineer:
		return "Engineer"
	case ID_Gunnery:
		return "Gunnery"
	case ID_Medic:
		return "Medic"
	case ID_Pilot:
		return "Pilot"
	case ID_Sensors:
		return "Sensors"
	case ID_Steward:
		return "Steward"
	case ID_Biologics:
		return "Biologics"
	case ID_Craftsman:
		return "Craftsman"
	case ID_Electronics:
		return "Electronics"
	case ID_Fluidics:
		return "Fluidics"
	case ID_Gravitics:
		return "Gravitics"
	case ID_Magnetics:
		return "Magnetics"
	case ID_Mechanic:
		return "Mechanic"
	case ID_Photonics:
		return "Photonics"
	case ID_Polymers:
		return "Polymers"
	case ID_Programmer:
		return "Programmer"
	case ID_Actor:
		return "Actor"
	case ID_Artist:
		return "Artist"
	case ID_Author:
		return "Author"
	case ID_Chef:
		return "Chef"
	case ID_Dancer:
		return "Dancer"
	case ID_Musician:
		return "Musician"
	case ID_Fighter:
		return "Fighter"
	case ID_Forward_Observer:
		return "Forward Observer"
	case ID_Heavy_Weapons:
		return "Heavy Weapons"
	case ID_Navigator:
		return "Navigator"
	case ID_Recon:
		return "Recon"
	case ID_Sapper:
		return "Sapper"
	case ID_Bay_Weapons:
		return "Bay Weapons"
	case ID_Ortilery:
		return "Ortilery"
	case ID_Screens:
		return "Screens"
	case ID_Spines:
		return "Spines"
	case ID_Turrets:
		return "Turrets"
	case ID_Artilery:
		return "Artilery"
	case ID_Launchers:
		return "Launchers"
	case ID_Ordinance:
		return "Ordinance"
	case ID_WMD:
		return "WMD"
	case ID_Battle_Dress:
		return "Battle Dress"
	case ID_Beams:
		return "Beams"
	case ID_Blades:
		return "Blades"
	case ID_Exotics:
		return "Exotics"
	case ID_Slugs:
		return "Slugs"
	case ID_Sprays:
		return "Sprays"
	case ID_Unarmed:
		return "Unarmed"
	case ID_Aeronautics:
		return "Aeronautics"
	case ID_Flappers:
		return "Flappers"
	case ID_Grav_f:
		return "Grav (f)"
	case ID_LTA:
		return "LTA"
	case ID_Rotor:
		return "Rotor"
	case ID_Winged:
		return "Winged"
	case ID_ACV:
		return "ACV"
	case ID_Automotive:
		return "Automotive"
	case ID_Grav_d:
		return "Grav (d)"
	case ID_Legged:
		return "Legged"
	case ID_Mole:
		return "Mole"
	case ID_Tracked:
		return "Tracked"
	case ID_Wheeled:
		return "Wheeled"
	case ID_Jump:
		return "Jump"
	case ID_Life_Support:
		return "Life Support"
	case ID_Maneuver:
		return "Maneuver"
	case ID_Power:
		return "Power"
	case ID_Rider:
		return "Rider"
	case ID_Teamster:
		return "Teamster"
	case ID_Trainer:
		return "Trainer"
	case ID_Aquanautics:
		return "Aquanautics"
	case ID_Grav_s:
		return "Grav (s)"
	case ID_Boat:
		return "Boat"
	case ID_Ship:
		return "Ship"
	case ID_Sub:
		return "Sub"
	case ID_Small_Craft:
		return "Small Craft"
	case ID_Spacecraft_ABS:
		return "Spacecraft ABS"
	case ID_Spacecraft_BCS:
		return "Spacecraft BCS"
	case ID_Biology:
		return "Biology"
	case ID_Chemistry:
		return "Chemistry"
	case ID_Physics:
		return "Physics"
	case ID_Planetology:
		return "Planetology"
	case ID_Robotics:
		return "Robotics"
	case ID_Archeology:
		return "Archeology"
	case ID_History:
		return "History"
	case ID_Linguistics:
		return "Linguistics"
	case ID_Philosophy:
		return "Philosophy"
	case ID_Psionicology:
		return "Psionicology"
	case ID_Psyhohistory:
		return "Psyhohistory"
	case ID_Psyhology:
		return "Psyhology"
	case ID_Sophontology:
		return "Sophontology"
	case ID_Instrument_Guitar:
		return "Instrument (Guitar)"
	case ID_Instrument_Banjo:
		return "Instrument (Banjo)"
	case ID_Instrument_Mandolin:
		return "Instrument (Mandolin)"
	case ID_Instrument_Keyboard:
		return "Instrument (Keyboard)"
	case ID_Instrument_Piano:
		return "Instrument (Piano)"
	case ID_Instrument_Voice:
		return "Instrument (Voice)"
	case ID_Instrument_Trombone:
		return "Instrument (Trombone)"
	case ID_Instrument_Trumpet:
		return "Instrument (Trumpet)"
	case ID_Instrument_Tuba:
		return "Instrument (Tuba)"
	case ID_Instrument_Violin:
		return "Instrument (Violin)"
	case ID_Instrument_Viola:
		return "Instrument (Viola)"
	case ID_Instrument_Cello:
		return "Instrument (Cello)"

	case ID_Kkree:
		return "Language: Kkree"
	case ID_Anglic:
		return "Language: Anglic"
	case ID_Battle:
		return "Language: Battle"
	case ID_Flash:
		return "Language: Flash"
	case ID_Gonk:
		return "Language: Gonk"
	case ID_Gvegh:
		return "Language: Gvegh"
	case ID_Mariel:
		return "Language: Mariel"
	case ID_Oynprith:
		return "Language: Oynprith"
	case ID_Sagamaal:
		return "Language: Sagamaal"
	case ID_Tezapet:
		return "Language: Tezapet"
	case ID_Trokh:
		return "Language: Trokh"
	case ID_Vilani:
		return "Language: Vilani"
	case ID_Zdetl:
		return "Language: Zdetl"

	case ID_Compute:
		return "Compute"
	case ID_Empath:
		return "Empath"
	case ID_Hibernate:
		return "Hibernate"
	case ID_Hypno:
		return "Hypno"
	case ID_Intuition:
		return "Intuition"
	case ID_Math:
		return "Math"
	case ID_Memaware:
		return "Memaware"
	case ID_Memorize:
		return "Memorize"
	case ID_Mempercept:
		return "Mempercept"
	case ID_Memscent:
		return "Memscent"
	case ID_Memsight:
		return "Memsight"
	case ID_Memsound:
		return "Memsound"
	case ID_Morph:
		return "Morph"
	case ID_Rage:
		return "Rage"
	case ID_Soundmimic:
		return "Soundmimic"
	}
}

const (
	LongestNameLength = 21
)

func LongestNameLen() int {
	lMax := 0
	for i := ID_NONE; i < ID_END; i++ {
		lCurrent := len(NameByID(i))
		if lCurrent > lMax {
			lMax = lCurrent
		}
	}
	return lMax
}

func (sk *Skill) SType() string {
	return sk.sklType
}

func (sk *Skill) Value() int {
	return sk.value
}

func (sk *Skill) Learn() error {
	switch sk.sklType {
	case TYPE_SKILL, TYPE_TALENT:
		if sk.value >= 15 {
			return fmt.Errorf("cap reached")
		}
	case TYPE_KNOWLEDGE:
		if sk.value >= 6 {
			return fmt.Errorf("cap reached")
		}
	}
	sk.value++
	return nil
}

func TradeCode2SkillID(tc string) []int {
	switch tc {
	default:
		return []int{ID_NONE}
	case "Ab":
		return []int{ID_NONE}
	case "Ag":
		return []int{ID_Animals}
	case "An":
		return []int{ID_NONE}
	case "As":
		return []int{ID_Zero_G}
	case "Ba":
		return []int{ID_NONE}
	case "Bo":
		return []int{ID_Hostile_Environ}
	case "Co":
		return []int{ID_Hostile_Environ}
	case "Cp":
		return []int{ID_Admin}
	case "Cs":
		return []int{ID_Bureaucrat}
	case "Cx":
		return []int{ID_Language}
	case "Da":
		return []int{ID_Fighter}
	case "De":
		return []int{ID_Survival}
	case "Di":
		return []int{ID_NONE}
	case "Ds":
		return []int{ID_Vacc_Suit, ID_Zero_G}
	case "Fa":
		return []int{ID_Animals}
	case "Fl":
		return []int{ID_Hostile_Environ}
	case "Fo":
		return []int{ID_NONE}
	case "Fr":
		return []int{ID_Hostile_Environ}
	case "Ga":
		return []int{ID_Trader}
	case "He":
		return []int{ID_Hostile_Environ}
	case "Hi":
		return []int{ID_Streetwise}
	case "Ho":
		return []int{ID_Hostile_Environ}
	case "Ic":
		return []int{ID_Vacc_Suit}
	case "In":
		return []int{One_Trade}
	case "Lk":
		return []int{ID_NONE}
	case "Lo":
		return []int{ID_Flyer}
	case "Mi":
		return []int{ID_Survey}
	case "Mr":
		return []int{ID_NONE}
	case "Na":
		return []int{ID_Survey}
	case "Ni":
		return []int{ID_Driver}
	case "Oc":
		return []int{ID_High_G}
	case "Pa":
		return []int{ID_Trader}
	case "Ph":
		return []int{ID_NONE}
	case "Pi":
		return []int{ID_JOT}
	case "Po":
		return []int{ID_Steward}
	case "Pr":
		return []int{ID_Craftsman}
	case "Px":
		return []int{ID_NONE}
	case "Pz":
		return []int{ID_NONE}
	case "Re":
		return []int{ID_NONE}
	case "Ri":
		return []int{One_Art}
	case "Tr":
		return []int{ID_Survival}
	case "Tu":
		return []int{ID_Survival}
	case "Tz":
		return []int{ID_Driver}
	case "Va":
		return []int{ID_Vacc_Suit}
	case "Wa":
		return []int{ID_Seafarer}
	}
}
