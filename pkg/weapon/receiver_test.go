package weapon

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

func testInpute() (input [][]int) {
	for _, ins1 := range allreceivers() {
		for _, ins2 := range allAmmo() {
			for _, ins3 := range allMechanismas() {
				for _, ins4 := range allTech() {
					funcFeat := []int{ins1, ins2, ins3, ins4}
					//input = append(input, []int{ins1, ins2, ins3, ins4})
					for _, v := range randomFromIntSlice(allFF()) {
						funcFeat = append(funcFeat, v)
					}
					input = append(input, funcFeat)
				}
			}
		}
	}
	return input
}

func allTech() []int {
	return []int{
		tech_CONVENTIONAL,
		tech_GAUSS_TECH,
		tech_ENERGY,
	}
}

func allFF() []int {
	return []int{
		feat_func_ADVANCED_PROJECTILE_WEAPON,
		feat_func_ACCUIRED,
		feat_func_BULLPUP,
		feat_func_COMPACT,
		feat_func_COMPACT_VERY,
		feat_func_COOLING_SYSTEM_BASIC,
		feat_func_COOLING_SYSTEM_ADVANCED,
		feat_func_GUIDENCE_SYSTEM,
		feat_func_HIGH_CAPACITY,
		feat_func_HIGH_QUALITY,
		feat_func_INCREASED_RATE_OF_FIRE_1,
		feat_func_INCREASED_RATE_OF_FIRE_2,
		feat_func_INCREASED_RATE_OF_FIRE_3,
		feat_func_INCREASED_RATE_OF_FIRE_4,
		feat_func_INCREASED_RATE_OF_FIRE_5,
		feat_func_INCREASED_RATE_OF_FIRE_6,
		feat_func_LIGHTWEIGHT,
		feat_func_LIGHTWEIGHT_EXTREAME,
		feat_func_LOW_QUALITY_1,
		feat_func_LOW_QUALITY_2,
		feat_func_LOW_QUALITY_3,
		feat_func_LOW_QUALITY_4,
		feat_func_LOW_QUALITY_5,
		feat_func_QUICKDRAW,
		feat_func_RECOIL_COMPENSATION,
		feat_func_RUGGED,
	}
}

func allreceivers() []int {
	return []int{
		receiver_HANDGUN,
		receiver_ASSAULT_WEAPON,
		receiver_LONGARM,
		receiver_LIGHT_SUPPORT_WEAPON,
		receiver_HEAVY_WEAPON,
	}
}

func allMechanismas() []int {
	return []int{
		pwm_SINGLE_SHOT,
		pwm_REPEATER,
		pwm_SEMI_AUTOMATIC,
		pwm_BURST_CAPABLE,
		pwm_FULLY_AUTOMATIC,
		pwm_RAPID_FIRE,
		pwm_VERY_RAPID_FIRE,
		pwm_UNDERWATER,
	}
}

func allAmmo() []int {
	return []int{
		ammo_HANDGUN_BlackPowder,
		ammo_HANDGUN_Light,
		ammo_HANDGUN_Medium,
		ammo_HANDGUN_Heavy,
		ammo_SHOTGUN_Smoothbores_Small,
		ammo_SHOTGUN_Smoothbores_Light,
		ammo_SHOTGUN_Smoothbores_Standard,
		ammo_SHOTGUN_Smoothbores_Heavy,
		ammo_LONGARM_BlackPowder,
		ammo_LONGARM_Rifle_Light,
		ammo_LONGARM_Rifle_Intermediate,
		ammo_LONGARM_Rifle_Battle,
		ammo_LONGARM_Rifle_AntiMaterial,
		ammo_LONGARM_Rifle_AntiMaterialHeavy,
		ammo_SNUB,
		ammo_Rocket,
		ammo_GAUSS_Standard,
		ammo_GAUSS_Small,
		ammo_GAUSS_Enchanced,
		ammo_GAUSS_Shotgun,
	}
}

func TestReciver(t *testing.T) {
	for i, input := range testInpute() {

		fmt.Printf("test %v (%v) ", i+1, input)
		r2, err := newReceiver(input...)
		if err != nil {
			//continue
			t.Errorf("error: %v", err)
			continue
		}
		if r2.tech == _UNDEFINED_ {
			continue
			t.Errorf("tech undefined")
		}
		fmt.Println(r2)
	}

}

func randomFromIntSlice(sl []int) []int {
	l := len(sl)
	dp := dice.New()
	totalLen := dp.Roll("1d6").Sum()
	res := []int{}
	for i := 0; i < totalLen; i++ {
		r := dp.Roll("1d" + strconv.Itoa(l)).Sum()
		res = append(res, sl[r-1])
	}
	return res
}

// func TestReciver(t *testing.T) {
// 	input := testInpute()
// 	totalTests := 0
// 	errorsDetected := 0

// 	for testNum, instruction := range input {
// 		totalTests++

// 		r, err := NewReceiver(instruction[0], instruction[1], instruction[2])
// 		if err != nil {
// 			errorsDetected++
// 			continue
// 			t.Errorf("creation error: %v", err.Error())

// 		}
// 		fmt.Printf("Test %v: (%v + %v + %v)\n", testNum+1, verbal(instruction[0]), verbal(instruction[1]), verbal(instruction[2]))
// 		fmt.Println("untested:", r)
// 		if instruction[2] == pwm_SINGLE_SHOT {
// 			if r.ammoCapacity != 1.0 {
// 				t.Errorf("expect ammo capacity '1', but have '%v'", r.ammoCapacity)
// 			}
// 		}

// 	}
// 	fmt.Println("Total tests:", totalTests, "| errors detected:", errorsDetected, "| valid: ", totalTests-errorsDetected)
// }

/*

-
1
2
3
4
5
12
13
14
15
23
24
25
34
35
123
124
125
134
135
234
235
1234
1235
1345
2345
12345

*/
