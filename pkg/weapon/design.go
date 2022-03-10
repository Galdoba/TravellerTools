package weapon

import "fmt"

/*
1 Decide general type (pistol, rifle, ect...)
2 choose ammnition/powertype
3 choose receiver
4 choose receiver's mode of operation (breechloader, semi-automatic, ect...)
5 assign barrel lenght
6 assign furniture
7 choose feed device
8 add accessories that come as standard
9 total cost and weight

*/

const (
	_UNDEFINED_ = iota
	receiver_HANDGUN
	receiver_ASSAULT_WEAPON
	receiver_LONGARM
	receiver_LIGHT_SUPPORT_WEAPON
	receiver_HEAVY_WEAPON
	receiver_GAUSS_TECH
	receiver_CONVENTIONAL
	ammo_HANDGUN_BlackPowder
	ammo_HANDGUN_Light
	ammo_HANDGUN_Medium
	ammo_HANDGUN_Heavy
	ammo_SHOTGUN_Smoothbores_Small
	ammo_SHOTGUN_Smoothbores_Light
	ammo_SHOTGUN_Smoothbores_Standard
	ammo_SHOTGUN_Smoothbores_Heavy
	ammo_LONGARM_BlackPowder
	ammo_LONGARM_Rifle_Light
	ammo_LONGARM_Rifle_Intermediate
	ammo_LONGARM_Rifle_Battle
	ammo_LONGARM_Rifle_AntiMaterial
	ammo_LONGARM_Rifle_AntiMaterialHeavy
	ammo_SNUB
	ammo_Rocket
	ammo_GAUSS_Standard
	ammo_GAUSS_Small
	ammo_GAUSS_Enchanced
	ammo_GAUSS_Shotgun
	pwm_SINGLE_SHOT
	pwm_REPEATER
	pwm_SEMI_AUTOMATIC
	pwm_BURST_CAPABLE
	pwm_FULLY_AUTOMATIC
	pwm_RAPID_FIRE
	pwm_VERY_RAPID_FIRE
	pwm_UNDERWATER
	feat_func_ADVANCED_PROJECTILE_WEAPON
	feat_func_ACCUIRED
	feat_func_BULLPUP
	feat_func_COMPACT
	feat_func_COMPACT_VERY
	feat_func_COOLING_SYSTEM_BASIC
	feat_func_COOLING_SYSTEM_ADVANCED
	feat_func_GUIDENCE_SYSTEM
	feat_func_HIGH_CAPACITY
	feat_func_HIGH_QUALITY
	feat_func_INCREASED_RATE_OF_FIRE_1
	feat_func_INCREASED_RATE_OF_FIRE_2
	feat_func_INCREASED_RATE_OF_FIRE_3
	feat_func_INCREASED_RATE_OF_FIRE_4
	feat_func_INCREASED_RATE_OF_FIRE_5
	feat_func_INCREASED_RATE_OF_FIRE_6
	feat_func_LIGHTWEIGHT
	feat_func_LIGHTWEIGHT_EXTREAME
	feat_func_LOW_QUALITY_1
	feat_func_LOW_QUALITY_2
	feat_func_LOW_QUALITY_3
	feat_func_LOW_QUALITY_4
	feat_func_LOW_QUALITY_5
	feat_func_QUICKDRAW
	feat_func_RECOIL_COMPENSATION
	feat_func_RUGGED
	feat_cap_ARMORED
	feat_cap_BULWARKED
	feat_cap_DISGUISED
	feat_cap_STEALTH_BASIC
	feat_cap_STEALTH_EXTREME
	feat_cap_VACUUM
	WRONG_INSTRUCTION
)

type Weapon struct {
	rcvr   *receiver
	recoil int
	/*
		generaltype	---
		receiver - handgun/assaultweapon/longarm/lightsupportweapon/heavyweapon
		ammunition
		powersource
		barrel
		furniture
		acceesoires
		feed devices

	*/
}

type receiver struct {
	baseCost     float64
	costMod      float64
	cost         float64
	baseWeight   float64
	weightMod    float64
	weight       float64
	baseAC       float64 //Base Ammunition Capacity
	acMod        float64
	ammoCapacity float64
	quickDraw    int
	gaussTech    bool
	pwm          string
	auto         int
	addedTraits  []string
	baseAP       int
}

type receiver2 struct {
	rType               int
	aType               int
	mechanism           int
	features_functional []int
	features_capability []int
	errors              []error
}

func NewReceiver(instructions ...int) *receiver2 {
	r := receiver2{}
	for _, inst := range instructions {
		err := r.addType(inst)
		if err != nil {
			r.errors = append(r.errors, err)
		}
	}
	return &r
}

func (r *receiver2) addType(inst int) error {
	switch {
	default:
		return fmt.Errorf("unknowm instruction '%v'", inst)
	case isRecieverType(inst):
		if r.rType != _UNDEFINED_ {
			return fmt.Errorf("reciver already assigned")
		}
		r.rType = inst
	case isAmmoType(inst):
		if r.aType != _UNDEFINED_ {
			return fmt.Errorf("ammo already assigned")
		}
		r.aType = inst
	case isMechanismType(inst):
		if r.mechanism != _UNDEFINED_ {
			return fmt.Errorf("PWM already assigned")
		}
		r.mechanism = inst
	case isFuncFeat(inst):
		if err := r.addFunctionalFeature(inst); err != nil {
			return fmt.Errorf("r.addFunctionalFeature(inst) '%v' error: %v", inst, err.Error())
		}
	}
	return nil
}

func (r *receiver2) String() string {
	str := ""
	str += "Receiver Type : " + verbal(r.rType) + "\n"
	str += "Ammo Type     : " + verbal(r.aType) + "\n"
	str += "Mechanism Type: " + verbal(r.mechanism) + "\n"
	str += "FUNCTIONAL FEATURES: \n"
	for _, ff := range r.features_functional {
		str += verbal(ff) + "\n"
	}
	str += "CAPABILITY FEATURES: \n"
	for _, fc := range r.features_capability {
		str += verbal(fc) + "\n"
	}
	return str
}

func isRecieverType(i int) bool {
	if i >= receiver_HANDGUN && i <= receiver_CONVENTIONAL {
		return true
	}
	return false
}

func isAmmoType(i int) bool {
	if i >= ammo_HANDGUN_BlackPowder && i <= ammo_GAUSS_Shotgun {
		return true
	}
	return false
}

func isMechanismType(i int) bool {
	if i >= pwm_SINGLE_SHOT && i <= pwm_UNDERWATER {
		return true
	}
	return false
}

func isFuncFeat(i int) bool {
	if i >= feat_func_ADVANCED_PROJECTILE_WEAPON && i <= feat_func_RUGGED {
		return true
	}
	return false
}

func (r *receiver2) addFunctionalFeature(i int) error {
	switch {
	case contains(r.features_functional, i):
		return fmt.Errorf("feature %v cannot be duplicated", i)
	case i == feat_func_COOLING_SYSTEM_BASIC:
		if contains(r.features_functional, feat_func_COOLING_SYSTEM_ADVANCED) {
			return fmt.Errorf("features %v and %v are not compatible", i, feat_func_COOLING_SYSTEM_ADVANCED)
		}
	case i == feat_func_COOLING_SYSTEM_ADVANCED:
		if contains(r.features_functional, feat_func_COOLING_SYSTEM_BASIC) {
			return fmt.Errorf("features %v and %v are not compatible", i, feat_func_COOLING_SYSTEM_BASIC)
		}
	case i == feat_func_LIGHTWEIGHT:
		if contains(r.features_functional, feat_func_LIGHTWEIGHT_EXTREAME) {
			return fmt.Errorf("features %v and %v are not compatible", i, feat_func_LIGHTWEIGHT_EXTREAME)
		}
	case i == feat_func_LIGHTWEIGHT_EXTREAME:
		if contains(r.features_functional, feat_func_LIGHTWEIGHT) {
			return fmt.Errorf("features %v and %v are not compatible", i, feat_func_LIGHTWEIGHT)
		}
	}
	r.features_functional = append(r.features_functional, i)
	return nil
}

func contains(sl []int, e int) bool {
	for _, val := range sl {
		if val == e {
			return true
		}
	}
	return false
}

func verbal(i int) string {
	switch i {
	default:
		return "Unknown"
	case _UNDEFINED_:
		return "_UNDEFINED_"
	case receiver_HANDGUN:
		return "receiver_HANDGUN"
	case receiver_ASSAULT_WEAPON:
		return "receiver_ASSAULT_WEAPON"
	case receiver_LONGARM:
		return "receiver_LONGARM"
	case receiver_LIGHT_SUPPORT_WEAPON:
		return "receiver_LIGHT_SUPPORT_WEAPON"
	case receiver_HEAVY_WEAPON:
		return "receiver_HEAVY_WEAPON"
	case receiver_GAUSS_TECH:
		return "receiver_GAUSS_TECH"
	case receiver_CONVENTIONAL:
		return "receiver_CONVENTIONAL"
	case ammo_HANDGUN_BlackPowder:
		return "ammo_HANDGUN_BlackPowder"
	case ammo_HANDGUN_Light:
		return "ammo_HANDGUN_Light"
	case ammo_HANDGUN_Medium:
		return "ammo_HANDGUN_Medium"
	case ammo_HANDGUN_Heavy:
		return "ammo_HANDGUN_Heavy"
	case ammo_SHOTGUN_Smoothbores_Small:
		return "ammo_SHOTGUN_Smoothbores_Small"
	case ammo_SHOTGUN_Smoothbores_Light:
		return "ammo_SHOTGUN_Smoothbores_Light"
	case ammo_SHOTGUN_Smoothbores_Standard:
		return "ammo_SHOTGUN_Smoothbores_Standard"
	case ammo_SHOTGUN_Smoothbores_Heavy:
		return "ammo_SHOTGUN_Smoothbores_Heavy"
	case ammo_LONGARM_BlackPowder:
		return "ammo_LONGARM_BlackPowder"
	case ammo_LONGARM_Rifle_Light:
		return "ammo_LONGARM_Rifle_Light"
	case ammo_LONGARM_Rifle_Intermediate:
		return "ammo_LONGARM_Rifle_Intermediate"
	case ammo_LONGARM_Rifle_Battle:
		return "ammo_LONGARM_Rifle_Battle"
	case ammo_LONGARM_Rifle_AntiMaterial:
		return "ammo_LONGARM_Rifle_AntiMaterial"
	case ammo_LONGARM_Rifle_AntiMaterialHeavy:
		return "ammo_LONGARM_Rifle_AntiMaterialHeavy"
	case ammo_SNUB:
		return "ammo_SNUB"
	case ammo_Rocket:
		return "ammo_Rocket"
	case ammo_GAUSS_Standard:
		return "ammo_GAUSS_Standard"
	case ammo_GAUSS_Small:
		return "ammo_GAUSS_Small"
	case ammo_GAUSS_Enchanced:
		return "ammo_GAUSS_Enchanced"
	case ammo_GAUSS_Shotgun:
		return "ammo_GAUSS_Shotgun"
	case pwm_SINGLE_SHOT:
		return "pwm_SINGLE_SHOT"
	case pwm_REPEATER:
		return "pwm_REPEATER"
	case pwm_SEMI_AUTOMATIC:
		return "pwm_SEMI_AUTOMATIC"
	case pwm_BURST_CAPABLE:
		return "pwm_BURST_CAPABLE"
	case pwm_FULLY_AUTOMATIC:
		return "pwm_FULLY_AUTOMATIC"
	case pwm_RAPID_FIRE:
		return "pwm_RAPID_FIRE"
	case pwm_VERY_RAPID_FIRE:
		return "pwm_VERY_RAPID_FIRE"
	case pwm_UNDERWATER:
		return "pwm_UNDERWATER"
	case feat_func_ADVANCED_PROJECTILE_WEAPON:
		return "feat_func_ADVANCED_PROJECTILE_WEAPON"
	case feat_func_ACCUIRED:
		return "feat_func_ACCUIRED"
	case feat_func_BULLPUP:
		return "feat_func_BULLPUP"
	case feat_func_COMPACT:
		return "feat_func_COMPACT"
	case feat_func_COOLING_SYSTEM_BASIC:
		return "feat_func_COOLING_SYSTEM_BASIC"
	case feat_func_COOLING_SYSTEM_ADVANCED:
		return "feat_func_COOLING_SYSTEM_ADVANCED"
	case feat_func_GUIDENCE_SYSTEM:
		return "feat_func_GUIDENCE_SYSTEM"
	case feat_func_HIGH_CAPACITY:
		return "feat_func_HIGH_CAPACITY"
	case feat_func_HIGH_QUALITY:
		return "feat_func_HIGH_QUALITY"
	case feat_func_INCREASED_RATE_OF_FIRE_1:
		return "feat_func_INCREASED_RATE_OF_FIRE_1"
	case feat_func_INCREASED_RATE_OF_FIRE_2:
		return "feat_func_INCREASED_RATE_OF_FIRE_2"
	case feat_func_INCREASED_RATE_OF_FIRE_3:
		return "feat_func_INCREASED_RATE_OF_FIRE_3"
	case feat_func_INCREASED_RATE_OF_FIRE_4:
		return "feat_func_INCREASED_RATE_OF_FIRE_4"
	case feat_func_INCREASED_RATE_OF_FIRE_5:
		return "feat_func_INCREASED_RATE_OF_FIRE_5"
	case feat_func_INCREASED_RATE_OF_FIRE_6:
		return "feat_func_INCREASED_RATE_OF_FIRE_6"
	case feat_func_LIGHTWEIGHT:
		return "feat_func_LIGHTWEIGHT"
	case feat_func_LIGHTWEIGHT_EXTREAME:
		return "feat_func_LIGHTWEIGHT_EXTREAME"
	case feat_func_LOW_QUALITY_1:
		return "feat_func_LOW_QUALITY_1"
	case feat_func_LOW_QUALITY_2:
		return "feat_func_LOW_QUALITY_2"
	case feat_func_LOW_QUALITY_3:
		return "feat_func_LOW_QUALITY_3"
	case feat_func_LOW_QUALITY_4:
		return "feat_func_LOW_QUALITY_4"
	case feat_func_LOW_QUALITY_5:
		return "feat_func_LOW_QUALITY_5"
	case feat_func_QUICKDRAW:
		return "feat_func_QUICKDRAW"
	case feat_func_RECOIL_COMPENSATION:
		return "feat_func_RECOIL_COMPENSATION"
	case feat_func_RUGGED:
		return "feat_func_RUGGED"
	case feat_func_COMPACT_VERY:
		return "feat_func_COMPACT_VERY"
	case feat_cap_ARMORED:
		return "feat_cap_ARMORED"
	case feat_cap_BULWARKED:
		return "feat_cap_BULWARKED"
	case feat_cap_DISGUISED:
		return "feat_cap_DISGUISED"
	case feat_cap_STEALTH_BASIC:
		return "feat_cap_STEALTH_BASIC"
	case feat_cap_STEALTH_EXTREME:
		return "feat_cap_STEALTH_EXTREME"
	case feat_cap_VACUUM:
		return "feat_cap_VACUUM"
	case WRONG_INSTRUCTION:
		return "WRONG_INSTRUCTION"
	}
}

// func NewReceiver(rType, tech, pwm uint) (*receiver, error) {
// 	r := &receiver{}
// 	r.costMod = 1
// 	r.weightMod = 1
// 	r.acMod = 1
// 	err := r.setupBaseParameters(rType, tech)
// 	if err != nil {
// 		return r, err
// 	}
// 	err = r.setupProjectileWeaponMechanism(pwm)
// 	if err != nil {
// 		return r, err
// 	}
// 	r.cost = r.baseCost * r.costMod
// 	r.weight = r.baseWeight * r.weightMod
// 	r.ammoCapacity = r.baseAC * r.acMod
// 	return r, nil
// }

func (r *receiver) addFunctionalFeature(feature uint) error {
	return fmt.Errorf("func (r *receiver) addFunctionalFeature(feature uint) error - not implemented")
}

func (r *receiver) setupBaseParameters(rType, tech uint) error {
	switch rType {
	default:
		return fmt.Errorf("Unknown 'rType' (%v) value", rType)
	case receiver_HANDGUN:
		r.baseCost = 175
		r.baseWeight = 0.8
		r.baseAC = 10
		r.quickDraw = 4
	case receiver_ASSAULT_WEAPON:
		r.baseCost = 300
		r.baseWeight = 2
		r.baseAC = 20
		r.quickDraw = 2
	case receiver_LONGARM:
		r.baseCost = 400
		r.baseWeight = 2.5
		r.baseAC = 30
		r.quickDraw = 0
	case receiver_LIGHT_SUPPORT_WEAPON:
		r.baseCost = 1500
		r.baseWeight = 5
		r.baseAC = 50
		r.quickDraw = -4
	case receiver_HEAVY_WEAPON:
		r.baseCost = 3000
		r.baseWeight = 10
		r.baseAC = 50
		r.quickDraw = -8
	}
	switch tech {
	default:
		return fmt.Errorf("Unknown 'tech' (%v) value", tech)
	case receiver_CONVENTIONAL:
	case receiver_GAUSS_TECH:
		r.gaussTech = true
		r.baseCost = r.baseCost * 2
		r.baseWeight = r.baseWeight * 1.25
		r.baseAC = r.baseAC * 3
	}
	return nil
}

func (r *receiver) setupProjectileWeaponMechanism(pwm uint) error {
	switch pwm {
	default:
		return fmt.Errorf("Unknown 'pwm' (%v) value", pwm)
	case pwm_SINGLE_SHOT:
		r.pwm = "pwm_SINGLE_SHOT"
		r.costMod = r.costMod - 0.75
		r.baseAC = 1
	case pwm_REPEATER:
		r.pwm = "pwm_REPEATER"
		r.costMod = r.costMod - 0.5
		r.acMod = r.acMod - 0.5
	case pwm_SEMI_AUTOMATIC:
		r.pwm = "pwm_SEMI_AUTOMATIC"
		//this is base-level PWM
	case pwm_BURST_CAPABLE:
		r.pwm = "pwm_BURST_CAPABLE"
		r.costMod = r.costMod + 0.1
		r.auto = 2
		r.addedTraits = append(r.addedTraits, "Auto")
	case pwm_FULLY_AUTOMATIC:
		r.pwm = "pwm_FULLY_AUTOMATIC"
		r.costMod = r.costMod + 0.2
		r.weightMod = r.weightMod + 0
		r.acMod = r.acMod + 0
		r.addedTraits = append(r.addedTraits, "Auto")
		r.auto = 3
	case pwm_UNDERWATER:
		r.pwm = "pwm_UNDERWATER"
		r.costMod = r.costMod + 1
		r.addedTraits = append(r.addedTraits, "Underwater")
	}
	return nil
}

type projectileMechanism struct {
}
