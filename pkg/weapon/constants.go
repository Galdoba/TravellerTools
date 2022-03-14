package weapon

const (
	_UNDEFINED_ = iota
	receiver_HANDGUN
	receiver_ASSAULT_WEAPON
	receiver_LONGARM
	receiver_LIGHT_SUPPORT_WEAPON
	receiver_HEAVY_WEAPON
	tech_GAUSS_TECH
	tech_ENERGY
	tech_CONVENTIONAL
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
	feat_func_LIGHTWEIGHT_EXTREME
	feat_func_LOW_QUALITY_1
	feat_func_LOW_QUALITY_2
	feat_func_LOW_QUALITY_3
	feat_func_LOW_QUALITY_4
	feat_func_LOW_QUALITY_5
	feat_func_QUICKDRAW
	feat_func_RECOIL_COMPENSATION_1
	feat_func_RECOIL_COMPENSATION_2
	feat_func_RUGGED
	feat_cap_ARMORED
	feat_cap_BULWARKED
	feat_cap_DISGUISED
	feat_cap_STEALTH_BASIC
	feat_cap_STEALTH_EXTREME
	feat_cap_VACUUM
	AMMUNITION_CAPACITY_50_LESS
	AMMUNITION_CAPACITY_40_LESS
	AMMUNITION_CAPACITY_30_LESS
	AMMUNITION_CAPACITY_20_LESS
	AMMUNITION_CAPACITY_10_LESS
	AMMUNITION_CAPACITY_STANDARD
	AMMUNITION_CAPACITY_10_MORE
	AMMUNITION_CAPACITY_20_MORE
	AMMUNITION_CAPACITY_30_MORE
	AMMUNITION_CAPACITY_40_MORE
	AMMUNITION_CAPACITY_50_MORE
	brl_len_MINIMAL
	brl_len_SHORT
	brl_len_HANDGUN
	brl_len_ASSAULT
	brl_len_CARBINE
	brl_len_RIFLE
	brl_len_LONG
	brl_len_VERY_LONG
	brl_weight_HEAVY
	brl_weight_STANDARD
	furniture_STOCKLESS
	furniture_STOCK_FOLDING
	furniture_STOCK_FULL
	furniture_MODULARIZATION
	furniture_BIPOD_ABSENT
	furniture_BIPOD_FIXED
	furniture_BIPOD_DETACHABLE
	furniture_SUPPORT_MOUNT
	accessoire_SUPPRESSOR_ABSENT
	accessoire_SUPPRESSOR_BASIC
	accessoire_SUPPRESSOR_STANDARD
	accessoire_SUPPRESSOR_EXTREME
	accessoire_AFD_MAGAZINE_FIXED
	accessoire_AFD_MAGAZINE_STANDARD
	accessoire_AFD_MAGAZINE_EXTENDED
	accessoire_AFD_MAGAZINE_DRUM
	accessoire_AFD_BELT
	accessoire_AFD_CLIPS
	accessoire_SCOPE_ABSENT
	accessoire_SCOPE_BASIC
	accessoire_SCOPE_LONG_RANGE
	accessoire_SCOPE_LOW_LIGHT
	accessoire_SCOPE_THERMAL
	accessoire_SCOPE_COMBINATION
	accessoire_SCOPE_MULTISPECTRAL
	accessoire_SCOPE_LASER_POINTER
	accessoire_SCOPE_ISS
	accessoire_SCOPE_HOLOGRAPHIC
	accessoire_OTHER_ABSENT
	accessoire_OTHER_BAYONET_LUG
	accessoire_OTHER_BLING
	accessoire_OTHER_FLASHLIGHT
	accessoire_OTHER_GRAVITIC_SUPPORT
	accessoire_OTHER_GUN_CAMERA
	accessoire_OTHER_WEAPON_INTELLIGENT
	accessoire_OTHER_WEAPON_SECURE
	accessoire_OTHER_STABILISATION
	WRONG_INSTRUCTION
)

func verbal(i int) string {
	switch i {
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
	case tech_GAUSS_TECH:
		return "tech_GAUSS_TECH"
	case tech_ENERGY:
		return "tech_ENERGY"
	case tech_CONVENTIONAL:
		return "tech_CONVENTIONAL"
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
	case feat_func_COMPACT_VERY:
		return "feat_func_COMPACT_VERY"
	case feat_func_COOLING_SYSTEM_BASIC:
		return "Cooling System, Basic"
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
	case feat_func_LIGHTWEIGHT_EXTREME:
		return "feat_func_LIGHTWEIGHT_EXTREME"
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
	case feat_func_RECOIL_COMPENSATION_1:
		return "feat_func_RECOIL_COMPENSATION_1"
	case feat_func_RECOIL_COMPENSATION_2:
		return "feat_func_RECOIL_COMPENSATION_2"
	case feat_func_RUGGED:
		return "feat_func_RUGGED"
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
	case brl_len_MINIMAL:
		return "brl_len_MINIMAL"
	case brl_len_SHORT:
		return "brl_len_SHORT"
	case brl_len_HANDGUN:
		return "brl_len_HANDGUN"
	case brl_len_ASSAULT:
		return "brl_len_ASSAULT"
	case brl_len_CARBINE:
		return "brl_len_CARBINE"
	case brl_len_RIFLE:
		return "brl_len_RIFLE"
	case brl_len_LONG:
		return "brl_len_LONG"
	case brl_len_VERY_LONG:
		return "brl_len_VERY_LONG"
	case brl_weight_HEAVY:
		return "brl_weight_HEAVY"
	case brl_weight_STANDARD:
		return "brl_weight_STANDARD"
	case WRONG_INSTRUCTION:
		return "WRONG_INSTRUCTION"
	}
	return "MUST NOT APPEAR"
}
