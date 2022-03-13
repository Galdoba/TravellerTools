package weapon

import (
	"fmt"
	"math/bits"
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
					funcFeat = append(funcFeat, randomFromIntSlice(allFF())...)
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

func TestReceiverManual(t *testing.T) {
	inputSet := [][]int{
		{1, 2, 3, 4},
		{1, 36, 22, 8},
		{1, 54, 32, 8},
	}
	for i, input := range inputSet {
		fmt.Println("Test", i+1, ":")
		r, err := newReceiver(input...)
		if err != nil {
			t.Errorf("error: %v", err.Error())
			fmt.Println(r.errorDescr)
			continue
		}
		fmt.Println(r)

	}
}

func TestReciver(t *testing.T) {
	input := append([]int{}, receiver_HANDGUN, receiver_LONGARM)
	input = append(input, feat_func_COMPACT_VERY, feat_func_COMPACT, feat_func_ADVANCED_PROJECTILE_WEAPON)
	input = append(input, feat_cap_STEALTH_EXTREME, feat_cap_DISGUISED)
	input = append(input, ammo_HANDGUN_BlackPowder, ammo_LONGARM_Rifle_Battle) // allAmmo()...)
	input = append(input, pwm_SEMI_AUTOMATIC, pwm_FULLY_AUTOMATIC)
	input = append(input, allTech()...)
	input = append(input, WRONG_INSTRUCTION)
	dtStr := []string{}
	for _, s := range input {
		dtStr = append(dtStr, strconv.Itoa(s))
	}
	errors := 0
	fmt.Println("calculating combinations for [", dtStr, "] 4...")
	comb := CombinationsTracked(dtStr, 0)
	testNum := 0
	for _, strComb := range comb {
		testNum++
		//fmt.Printf("Start test %v (%v) \n", testNum, strComb)
		inp := []int{}
		for _, sInp := range strComb {
			i, _ := strconv.Atoi(sInp)
			inp = append(inp, i)
		}
		r2, err := newReceiver(inp...)
		if err != nil {
			errors++
			if err.Error() == "Input is incorrect" {
				continue
			}
			if err.Error() == "unknowm instruction '79'" {
				continue
			}
			t.Errorf("error: -%v-", err)

			//fmt.Println(r2.errorDescr)
			continue
		}
		fmt.Printf("Test %v (%v) \n", testNum, strComb)
		if r2.tech == _UNDEFINED_ {
			t.Errorf("tech undefined")
			errors++
		}
		if r2.rType == _UNDEFINED_ {
			t.Errorf("reciver type undefined")
			errors++
		}
		if r2.aType == _UNDEFINED_ {
			t.Errorf("callibre undefined")
			errors++
		}
		if r2.mechanism == _UNDEFINED_ {
			t.Errorf("mechanism undefined")
			errors++
		}
		fmt.Println(r2)
	}
	fmt.Println("--------")
	fmt.Println("--------")
	fmt.Println("Total", testNum, " | errors", errors, "| correct =", testNum-errors)
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

/*



 */

func CombinationsTracked(set []string, n int) (subsets [][]string) {

	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	total := (1 << length)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
		fmt.Print("  ", subsetBits*100/total, "% of ", total, ": ", len(subsets), " New combination ", subset, " | Total:                  \r")
	}
	return subsets
}
