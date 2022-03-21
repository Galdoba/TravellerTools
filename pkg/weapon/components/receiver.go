package components

import "fmt"

type Receiver struct {
	rType               int
	aType               int
	mechanism           int
	gauss               bool
	ammoCapacityMod     int
	features_functional []int
	features_capability []int
	errorDescr          string
}

func NewReceiver(instructions ...int) (*Receiver, error) {
	r := Receiver{}
	if err := r.analize(instructions); err != nil {
		return &r, err
	}
	if r.errorDescr != "" {
		return &r, inErr()
	}
	for _, inst := range instructions {
		err := r.addType(inst)
		if err != nil {
			return &r, err
		}
	}
	return &r, nil
}

func (r *Receiver) analize(instructions []int) error {
	if timesCrossed(instructions, []int{FEAT_func_COMPACT, FEAT_func_COMPACT_VERY, FEAT_func_HIGH_CAPACITY}) > 1 {
		return fmt.Errorf("Functional Features: Compact, Very Compact and High Capacity - can't coexist")
	}
	if timesCrossed(instructions, []int{FEAT_func_LIGHTWEIGHT_EXTREME, FEAT_func_LIGHTWEIGHT}) > 1 {
		return fmt.Errorf("Functional Features: Lightweight and Lightweight EXTREME - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_COOLING_SYSTEM_ADVANCED, FEAT_func_COOLING_SYSTEM_BASIC}) > 1 {
		return fmt.Errorf("Functional Features: 'Cooling System, Basic' and 'Cooling System, Advanced' - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_INCREASED_RATE_OF_FIRE_1, FEAT_func_INCREASED_RATE_OF_FIRE_2, FEAT_func_INCREASED_RATE_OF_FIRE_3, FEAT_func_INCREASED_RATE_OF_FIRE_4, FEAT_func_INCREASED_RATE_OF_FIRE_5, FEAT_func_INCREASED_RATE_OF_FIRE_6}) > 1 {
		return fmt.Errorf("Functional Features: multiple 'Increased Rate of Fire' features - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_LOW_QUALITY_1, FEAT_func_LOW_QUALITY_2, FEAT_func_LOW_QUALITY_3, FEAT_func_LOW_QUALITY_4, FEAT_func_LOW_QUALITY_5}) > 1 {
		return fmt.Errorf("Functional Features: multiple 'Low Quality' features - can't coexist")

	}
	ttMet := timesCrossed(instructions, allreceiverInstructions("Tech Type"))
	if ttMet > 1 {
		return fmt.Errorf(fmt.Sprintf("Tech type met %v times in instructions set", ttMet))

	}
	rtMet := timesCrossed(instructions, allreceiverInstructions("Receiver Type"))
	if rtMet < 1 {
		return fmt.Errorf("Receiver type instruction is missing")

	}
	if rtMet > 1 {
		return fmt.Errorf(fmt.Sprintf("Receiver type met %v times in instructions set", rtMet))

	}
	mtMet := timesCrossed(instructions, allreceiverInstructions("PWM Type"))
	if mtMet < 1 {
		return fmt.Errorf("PWM type instruction is missing")

	}
	if mtMet > 1 {
		return fmt.Errorf(fmt.Sprintf("PWM type met %v times in instructions set", mtMet))

	}
	atMet := timesCrossed(instructions, allreceiverInstructions("Ammo Type"))
	if atMet < 1 {
		return fmt.Errorf("Ammo type instruction is missing")

	}
	if atMet != 1 {
		return fmt.Errorf(fmt.Sprintf("Ammo type met %v times in instructions set", atMet))
	}
	if ttMet == 1 && timesCrossed(instructions, []int{CALLIBRE_GAUSS_Enchanced, CALLIBRE_GAUSS_Shotgun, CALLIBRE_GAUSS_Small, CALLIBRE_GAUSS_Standard}) < 1 {
		return fmt.Errorf(fmt.Sprintf("ammunition: Gauss technology MUST correlate with Gauss Ammo"))
	}
	acMet := timesCrossed(instructions, allreceiverInstructions("Ammunition Capacity"))
	if acMet < 1 {
		return fmt.Errorf("Ammunition Capacity type instruction is missing")

	}
	if acMet > 1 {
		return fmt.Errorf(fmt.Sprintf("Ammunition Capacity type met %v times in instructions set", acMet))

	}
	//functional Features Exclusions:
	if timesCrossed(instructions, []int{FEAT_func_COMPACT, FEAT_func_COMPACT_VERY, FEAT_func_HIGH_CAPACITY}) > 1 {
		return fmt.Errorf("Functional Features: Compact, Very Compact and High Capacity - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_LIGHTWEIGHT_EXTREME, FEAT_func_LIGHTWEIGHT}) > 1 {
		return fmt.Errorf("Functional Features: Lightweight and Lightweight EXTREME - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_COOLING_SYSTEM_ADVANCED, FEAT_func_COOLING_SYSTEM_BASIC}) > 1 {
		return fmt.Errorf("Functional Features: 'Cooling System, Basic' and 'Cooling System, Advanced' - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_INCREASED_RATE_OF_FIRE_1, FEAT_func_INCREASED_RATE_OF_FIRE_2, FEAT_func_INCREASED_RATE_OF_FIRE_3, FEAT_func_INCREASED_RATE_OF_FIRE_4, FEAT_func_INCREASED_RATE_OF_FIRE_5, FEAT_func_INCREASED_RATE_OF_FIRE_6}) > 1 {
		return fmt.Errorf("Functional Features: multiple 'Increased Rate of Fire' features - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_LOW_QUALITY_1, FEAT_func_LOW_QUALITY_2, FEAT_func_LOW_QUALITY_3, FEAT_func_LOW_QUALITY_4, FEAT_func_LOW_QUALITY_5}) > 1 {
		return fmt.Errorf("Functional Features: multiple 'Low Quality' features - can't coexist")

	}
	//capability Features Exclusions:
	if timesCrossed(instructions, []int{FEAT_cap_STEALTH_BASIC, FEAT_cap_STEALTH_EXTREME}) > 1 {
		return fmt.Errorf("Capability Features: 'Stealth, Basic' and 'Stealth, Extreme' - can't coexist")

	}
	if timesCrossed(instructions, []int{FEAT_func_RECOIL_COMPENSATION_1, FEAT_func_RECOIL_COMPENSATION_2}) > 1 {
		return fmt.Errorf("Capability Features: multiple 'Recoil Compensation' features - can't coexist")

	}
	return nil

}

func (r *Receiver) addType(inst int) error {
	switch {
	default:
		//return fmt.Errorf("unknowm instruction '%v'", inst)
	case isRecieverType(inst):
		r.rType = inst
	case isAmmoType(inst):
		r.aType = inst
	case isMechanismType(inst):
		r.mechanism = inst
	case isTech(inst):
		r.gauss = true
	case isAmmunitionCapacity(inst):
		r.ammoCapacityMod = inst
	case isFuncFeat(inst):
		r.features_functional = append(r.features_functional, inst)
	case isCapFeat(inst):
		r.features_capability = append(r.features_capability, inst)
	}
	return nil
}

func (r *Receiver) String() string {
	if r.errorDescr != "" {
		return "INTERNAL ERROR: " + r.errorDescr
	}
	str := ""
	str += "Receiver Type : " + Verbal(r.rType) + "\n"
	str += "Ammo Type     : " + Verbal(r.aType) + "\n"
	str += "Mechanism Type: " + Verbal(r.mechanism) + "\n"

	str += "FUNCTIONAL FEATURES: \n"
	for _, ff := range r.features_functional {
		str += Verbal(ff) + "\n"
	}
	str += "CAPABILITY FEATURES: \n"
	for _, fc := range r.features_capability {
		str += Verbal(fc) + "\n"
	}
	return str
}

func isRecieverType(i int) bool {
	if i >= RCVR_TYPE_HANDGUN && i <= RCVR_TYPE_HEAVY_WEAPON {
		return true
	}
	return false
}

func allreceiverInstructions(types string) []int {
	ri := []int{}
	switch types {
	case "Receiver Type":
		for i := RCVR_TYPE_HANDGUN; i <= RCVR_TYPE_HEAVY_WEAPON; i++ {
			ri = append(ri, i)
		}
	case "Ammo Type":
		for i := CALLIBRE_HANDGUN_BlackPowder; i <= CALLIBRE_GAUSS_Shotgun; i++ {
			ri = append(ri, i)
		}
	case "Tech Type":
		for i := TECH_GAUSS_TECH; i <= TECH_GAUSS_TECH; i++ {
			ri = append(ri, i)
		}
	case "PWM Type":
		for i := MECHANISM_SINGLE_SHOT; i <= MECHANISM_UNDERWATER; i++ {
			ri = append(ri, i)
		}
	case "func_feat":
		for i := FEAT_func_ADVANCED_PROJECTILE_WEAPON; i <= FEAT_func_RUGGED; i++ {
			ri = append(ri, i)
		}
	case "cap_feat":
		for i := FEAT_cap_ARMORED; i <= FEAT_cap_VACUUM; i++ {
			ri = append(ri, i)
		}
	case "Ammunition Capacity":
		for i := AMMUNITION_CAPACITY_50PCT_LESS; i <= AMMUNITION_CAPACITY_50PCT_MORE; i++ {
			ri = append(ri, i)
		}
	}
	return ri
}

func isAmmoType(i int) bool {
	if i >= CALLIBRE_HANDGUN_BlackPowder && i <= CALLIBRE_GAUSS_Shotgun {
		return true
	}
	return false
}

func isMechanismType(i int) bool {
	if i >= MECHANISM_SINGLE_SHOT && i <= MECHANISM_UNDERWATER {
		return true
	}
	return false
}

func isTech(i int) bool {
	if i >= TECH_GAUSS_TECH && i <= TECH_CONVENTIONAL {
		return true
	}
	return false
}

func isAmmunitionCapacity(i int) bool {
	if i >= AMMUNITION_CAPACITY_50PCT_LESS && i <= AMMUNITION_CAPACITY_50PCT_MORE {
		return true
	}
	return false
}

func isFuncFeat(i int) bool {
	if i >= FEAT_func_ADVANCED_PROJECTILE_WEAPON && i <= FEAT_func_RUGGED {
		return true
	}
	return false
}

func isCapFeat(i int) bool {
	if i >= FEAT_cap_ARMORED && i <= FEAT_cap_VACUUM {
		return true
	}
	return false
}

func (r *Receiver) addFunctionalFeature(i int) {
	r.features_functional = append(r.features_functional, i)
}

func (r *Receiver) addCapabilityFeature(i int) {
	r.features_capability = append(r.features_capability, i)
}

func capFeatureError(accepted []int, new int) error {
	exclusionsList := accepted
	errorMap := make(map[int]error)
	if contains(exclusionsList, new) {
		return fmt.Errorf("instruction duplicated (%v) [%v]", new, accepted)
	}
	for _, v := range exclusionsList {
		switch {
		case v == FEAT_cap_STEALTH_BASIC || v == FEAT_cap_STEALTH_EXTREME:
			errorMap[FEAT_cap_STEALTH_BASIC] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
			errorMap[FEAT_cap_STEALTH_EXTREME] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
		}
	}
	for k, v := range errorMap {
		if k == new {
			return v
		}
	}

	return nil
}

//ReceiverType - returns Receiver Type
func (r *Receiver) ReceiverType() int {
	return r.rType
}

//Mechanism - returns Mechanism
func (r *Receiver) Mechanism() int {
	return r.mechanism
}

func (r *Receiver) AmmunitionType() int {
	return r.aType
}

func (r *Receiver) AmmunitionCapacity() int {
	return r.ammoCapacityMod
}

func (r *Receiver) IsGauss() bool {
	return r.gauss
}