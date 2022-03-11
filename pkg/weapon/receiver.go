package weapon

import "fmt"

type receiver struct {
	rType               int
	aType               int
	mechanism           int
	tech                int
	features_functional []int
	features_capability []int
}

func newReceiver(instructions ...int) (*receiver, error) {
	r := receiver{}
	for _, inst := range instructions {
		err := r.addType(inst)
		if err != nil {
			return nil, err
		}
	}
	return &r, nil
}

func (r *receiver) addType(inst int) error {
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
			return fmt.Errorf("pwm already assigned")
		}
		r.mechanism = inst
	case isTech(inst):
		if r.tech != _UNDEFINED_ {
			return fmt.Errorf("tech already assigned")
		}
		r.tech = inst
	case isFuncFeat(inst):
		if err := r.addFunctionalFeature(inst); err != nil {
			return fmt.Errorf("functional feature (%v) error: %v", inst, err.Error())
		}
	case isCapFeat(inst):
		if err := r.addCapabilityFeature(inst); err != nil {
			return fmt.Errorf("capability feature (%v) error: %v", inst, err.Error())
		}
	}
	return nil
}

// func (r *receiver) String() string {
// 	str := ""
// 	str += "Receiver Type : " + verbal(r.rType) + "\n"
// 	str += "Ammo Type     : " + verbal(r.aType) + "\n"
// 	str += "Mechanism Type: " + verbal(r.mechanism) + "\n"
// 	str += "FUNCTIONAL FEATURES: \n"
// 	for _, ff := range r.features_functional {
// 		str += verbal(ff) + "\n"
// 	}
// 	str += "CAPABILITY FEATURES: \n"
// 	for _, fc := range r.features_capability {
// 		str += verbal(fc) + "\n"
// 	}
// 	return str
// }

func isRecieverType(i int) bool {
	if i >= receiver_HANDGUN && i <= receiver_HEAVY_WEAPON {
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

func isTech(i int) bool {
	if i >= tech_GAUSS_TECH && i <= tech_CONVENTIONAL {
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

func isCapFeat(i int) bool {
	if i >= feat_cap_ARMORED && i <= feat_cap_VACUUM {
		return true
	}
	return false
}

func (r *receiver) addFunctionalFeature(i int) error {
	if err := funcFeatureError(r.features_functional, i); err != nil {
		return err
	}
	r.features_functional = append(r.features_functional, i)
	return nil
}

func (r *receiver) addCapabilityFeature(i int) error {
	if err := capFeatureError(r.features_capability, i); err != nil {
		return err
	}
	r.features_capability = append(r.features_capability, i)
	return nil
}

func funcFeatureError(accepted []int, new int) error {
	exclusionsList := accepted
	errorMap := make(map[int]error)
	if contains(exclusionsList, new) {
		return fmt.Errorf("instruction duplicated (%v) [%v]", new, accepted)
	}
	for _, v := range exclusionsList {
		switch {
		case v == feat_func_COOLING_SYSTEM_ADVANCED || v == feat_func_COOLING_SYSTEM_BASIC:
			errorMap[feat_func_COOLING_SYSTEM_ADVANCED] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
			errorMap[feat_func_COOLING_SYSTEM_BASIC] = fmt.Errorf("functions %v and %v cannot coexist", v, new)

		case v == feat_func_LIGHTWEIGHT_EXTREAME || v == feat_func_LIGHTWEIGHT:
			errorMap[feat_func_LIGHTWEIGHT_EXTREAME] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
			errorMap[feat_func_LIGHTWEIGHT] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
		case v == feat_func_COMPACT_VERY || v == feat_func_COMPACT:
			errorMap[feat_func_COMPACT_VERY] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
			errorMap[feat_func_COMPACT] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
		case v >= feat_func_INCREASED_RATE_OF_FIRE_1 && v <= feat_func_INCREASED_RATE_OF_FIRE_6:
			errorMap[feat_func_INCREASED_RATE_OF_FIRE_1] = fmt.Errorf("multiple 'INCREASED_RATE_OF_FIRE' features can not coexist (%v) - [%v %v %v %v %v %v]", new, feat_func_INCREASED_RATE_OF_FIRE_1, feat_func_INCREASED_RATE_OF_FIRE_2, feat_func_INCREASED_RATE_OF_FIRE_3, feat_func_INCREASED_RATE_OF_FIRE_4, feat_func_INCREASED_RATE_OF_FIRE_5, feat_func_INCREASED_RATE_OF_FIRE_6)
			errorMap[feat_func_INCREASED_RATE_OF_FIRE_2] = fmt.Errorf("multiple 'INCREASED_RATE_OF_FIRE' features can not coexist (%v) - [%v %v %v %v %v %v]", new, feat_func_INCREASED_RATE_OF_FIRE_1, feat_func_INCREASED_RATE_OF_FIRE_2, feat_func_INCREASED_RATE_OF_FIRE_3, feat_func_INCREASED_RATE_OF_FIRE_4, feat_func_INCREASED_RATE_OF_FIRE_5, feat_func_INCREASED_RATE_OF_FIRE_6)
			errorMap[feat_func_INCREASED_RATE_OF_FIRE_3] = fmt.Errorf("multiple 'INCREASED_RATE_OF_FIRE' features can not coexist (%v) - [%v %v %v %v %v %v]", new, feat_func_INCREASED_RATE_OF_FIRE_1, feat_func_INCREASED_RATE_OF_FIRE_2, feat_func_INCREASED_RATE_OF_FIRE_3, feat_func_INCREASED_RATE_OF_FIRE_4, feat_func_INCREASED_RATE_OF_FIRE_5, feat_func_INCREASED_RATE_OF_FIRE_6)
			errorMap[feat_func_INCREASED_RATE_OF_FIRE_4] = fmt.Errorf("multiple 'INCREASED_RATE_OF_FIRE' features can not coexist (%v) - [%v %v %v %v %v %v]", new, feat_func_INCREASED_RATE_OF_FIRE_1, feat_func_INCREASED_RATE_OF_FIRE_2, feat_func_INCREASED_RATE_OF_FIRE_3, feat_func_INCREASED_RATE_OF_FIRE_4, feat_func_INCREASED_RATE_OF_FIRE_5, feat_func_INCREASED_RATE_OF_FIRE_6)
			errorMap[feat_func_INCREASED_RATE_OF_FIRE_5] = fmt.Errorf("multiple 'INCREASED_RATE_OF_FIRE' features can not coexist (%v) - [%v %v %v %v %v %v]", new, feat_func_INCREASED_RATE_OF_FIRE_1, feat_func_INCREASED_RATE_OF_FIRE_2, feat_func_INCREASED_RATE_OF_FIRE_3, feat_func_INCREASED_RATE_OF_FIRE_4, feat_func_INCREASED_RATE_OF_FIRE_5, feat_func_INCREASED_RATE_OF_FIRE_6)
			errorMap[feat_func_INCREASED_RATE_OF_FIRE_6] = fmt.Errorf("multiple 'INCREASED_RATE_OF_FIRE' features can not coexist (%v) - [%v %v %v %v %v %v]", new, feat_func_INCREASED_RATE_OF_FIRE_1, feat_func_INCREASED_RATE_OF_FIRE_2, feat_func_INCREASED_RATE_OF_FIRE_3, feat_func_INCREASED_RATE_OF_FIRE_4, feat_func_INCREASED_RATE_OF_FIRE_5, feat_func_INCREASED_RATE_OF_FIRE_6)
		case v >= feat_func_LOW_QUALITY_1 && v <= feat_func_LOW_QUALITY_5:
			errorMap[feat_func_LOW_QUALITY_1] = fmt.Errorf("multiple 'LOW_QUALITY' features can not coexist (%v) - [%v]", new, accepted)
			errorMap[feat_func_LOW_QUALITY_2] = fmt.Errorf("multiple 'LOW_QUALITY' features can not coexist (%v) - [%v]", new, accepted)
			errorMap[feat_func_LOW_QUALITY_3] = fmt.Errorf("multiple 'LOW_QUALITY' features can not coexist (%v) - [%v]", new, accepted)
			errorMap[feat_func_LOW_QUALITY_4] = fmt.Errorf("multiple 'LOW_QUALITY' features can not coexist (%v) - [%v]", new, accepted)
			errorMap[feat_func_LOW_QUALITY_5] = fmt.Errorf("multiple 'LOW_QUALITY' features can not coexist (%v) - [%v]", new, accepted)
		}
	}
	for k, v := range errorMap {
		if k == new {
			return v
		}
	}

	return nil
}

func capFeatureError(accepted []int, new int) error {
	exclusionsList := accepted
	errorMap := make(map[int]error)
	if contains(exclusionsList, new) {
		return fmt.Errorf("instruction duplicated (%v) [%v]", new, accepted)
	}
	for _, v := range exclusionsList {
		switch {
		case v == feat_cap_STEALTH_BASIC || v == feat_cap_STEALTH_EXTREME:
			errorMap[feat_cap_STEALTH_BASIC] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
			errorMap[feat_cap_STEALTH_EXTREME] = fmt.Errorf("functions %v and %v cannot coexist", v, new)
		}
	}
	for k, v := range errorMap {
		if k == new {
			return v
		}
	}

	return nil
}
