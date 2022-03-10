package weapon

import (
	"fmt"
	"testing"
)

func testInpute() (input [][]int) {
	for _, ins1 := range allreceivers() {
		for _, ins2 := range allAmmo() {
			for _, ins3 := range allMechanismas() {
				input = append(input, []int{ins1, ins2, ins3})
			}
		}
	}
	return input
}

func allreceivers() []int {
	return []int{
		receiver_HANDGUN,
		receiver_ASSAULT_WEAPON,
		receiver_LONGARM,
		receiver_LIGHT_SUPPORT_WEAPON,
		receiver_HEAVY_WEAPON,
		receiver_GAUSS_TECH,
		receiver_CONVENTIONAL,
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

func TestReciver2(t *testing.T) {
	for i, input := range testInpute() {
		fmt.Println("test", i+1)
		r2 := NewReceiver(input...)
		for _, err := range r2.errors {
			t.Errorf("error: %v", err)
		}
		//fmt.Println(r2)
	}

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
