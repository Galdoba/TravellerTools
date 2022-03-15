package components

import (
	"fmt"
	"math/bits"
	"sort"
	"strconv"
	"testing"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

func allTech() []int {
	return []int{
		TECH_CONVENTIONAL,
		TECH_GAUSS_TECH,
		TECH_ENERGY,
	}
}

func allFF() []int {
	return []int{
		FEAT_func_ADVANCED_PROJECTILE_WEAPON,
		FEAT_func_ACCUIRED,
		FEAT_func_BULLPUP,
		FEAT_func_COMPACT,
		FEAT_func_COMPACT_VERY,
		FEAT_func_COOLING_SYSTEM_BASIC,
		FEAT_func_COOLING_SYSTEM_ADVANCED,
		FEAT_func_GUIDENCE_SYSTEM,
		FEAT_func_HIGH_CAPACITY,
		FEAT_func_HIGH_QUALITY,
		FEAT_func_INCREASED_RATE_OF_FIRE_1,
		FEAT_func_INCREASED_RATE_OF_FIRE_2,
		FEAT_func_INCREASED_RATE_OF_FIRE_3,
		FEAT_func_INCREASED_RATE_OF_FIRE_4,
		FEAT_func_INCREASED_RATE_OF_FIRE_5,
		FEAT_func_INCREASED_RATE_OF_FIRE_6,
		FEAT_func_LIGHTWEIGHT,
		FEAT_func_LIGHTWEIGHT_EXTREME,
		FEAT_func_LOW_QUALITY_1,
		FEAT_func_LOW_QUALITY_2,
		FEAT_func_LOW_QUALITY_3,
		FEAT_func_LOW_QUALITY_4,
		FEAT_func_LOW_QUALITY_5,
		FEAT_func_QUICKDRAW,
		FEAT_func_RECOIL_COMPENSATION_1,
		FEAT_func_RECOIL_COMPENSATION_2,
		FEAT_func_RUGGED,
	}
}

func allreceivers() []int {
	return []int{
		RCVR_TYPE_HANDGUN,
		RCVR_TYPE_ASSAULT_WEAPON,
		RCVR_TYPE_LONGARM,
		RCVR_TYPE_LIGHT_SUPPORT_WEAPON,
		RCVR_TYPE_HEAVY_WEAPON,
	}
}

func allMechanismas() []int {
	return []int{
		MECHANISM_SINGLE_SHOT,
		MECHANISM_REPEATER,
		MECHANISM_SEMI_AUTOMATIC,
		MECHANISM_BURST_CAPABLE,
		MECHANISM_FULLY_AUTOMATIC,
		//MECHANISM_RAPID_FIRE,
		//MECHANISM_VERY_RAPID_FIRE,
		MECHANISM_UNDERWATER,
	}
}

func allAmmo() []int {
	return []int{
		CALLIBRE_HANDGUN_BlackPowder,
		CALLIBRE_HANDGUN_Light,
		CALLIBRE_HANDGUN_Medium,
		CALLIBRE_HANDGUN_Heavy,
		CALLIBRE_SHOTGUN_Smoothbores_Small,
		CALLIBRE_SHOTGUN_Smoothbores_Light,
		CALLIBRE_SHOTGUN_Smoothbores_Standard,
		CALLIBRE_SHOTGUN_Smoothbores_Heavy,
		CALLIBRE_LONGARM_BlackPowder,
		CALLIBRE_LONGARM_Rifle_Light,
		CALLIBRE_LONGARM_Rifle_Intermediate,
		CALLIBRE_LONGARM_Rifle_Battle,
		CALLIBRE_LONGARM_Rifle_AntiMaterial,
		CALLIBRE_LONGARM_Rifle_AntiMaterialHeavy,
		CALLIBRE_SNUB,
		CALLIBRE_Rocket,
		CALLIBRE_GAUSS_Standard,
		CALLIBRE_GAUSS_Small,
		CALLIBRE_GAUSS_Enchanced,
		CALLIBRE_GAUSS_Shotgun,
	}
}

func TestReceiverManual(t *testing.T) {
	return
	inputSet := [][]int{
		{1, 2, 3, 4},
		{1, 36, 22, 8},
		{1, 54, 32, 8},
	}
	for i, input := range inputSet {
		fmt.Println("Test", i+1, ":")
		r, err := NewReceiver(input...)
		if err != nil {
			t.Errorf("error: %v", err.Error())
			fmt.Println(r.errorDescr)
			continue
		}
		fmt.Println(r)

	}
}

func TestReciver(t *testing.T) {
	input := []int{}
	input = append(input, RCVR_TYPE_HANDGUN) //, RCVR_TYPE_LONGARM)
	//input = append(input, allFF()...)
	//input = append(input, FEAT_cap_STEALTH_EXTREME, FEAT_cap_STEALTH_BASIC)                                                                                                                                                                                 //, FEAT_cap_BULWARKED)
	input = append(input, CALLIBRE_HANDGUN_Medium)
	input = append(input, MECHANISM_SEMI_AUTOMATIC)     //, MECHANISM_FULLY_AUTOMATIC)
	input = append(input, TECH_CONVENTIONAL)            // allTech()...)
	input = append(input, AMMUNITION_CAPACITY_STANDARD) //, AMMUNITION_CAPACITY_STANDARD, AMMUNITION_CAPACITY_50_MORE)
	//input = append(input, WRONG_INSTRUCTION)
	dtStr := []string{}
	for _, s := range input {
		dtStr = append(dtStr, strconv.Itoa(s))
	}
	errors := 0
	fmt.Println("calculating combinations for", dtStr)
	comb := CombinationsTracked(dtStr, 0, false)
	testNum := 0
	errrMap := make(map[string]int)
	for _, strComb := range comb {
		testNum++
		//fmt.Printf("Start test %v (%v) \n", testNum, strComb)
		inp := []int{}
		for _, sInp := range strComb {
			i, _ := strconv.Atoi(sInp)
			inp = append(inp, i)
		}
		_, err := NewReceiver(inp...)
		if err != nil {
			errrMap[err.Error()]++
			errors++
			//t.Errorf("error: %v (%v)", err, inp)
			continue
		}
		fmt.Printf("Test %v (%v) \n", testNum, strComb)

	}
	fmt.Println("Total", testNum, " | errors", errors, "| correct =", testNum-errors)
	errNames := []string{}
	for k, _ := range errrMap {
		if k == "" {
			continue
		}
		errNames = append(errNames, k)
	}
	sort.Strings(errNames)
	for _, name := range errNames {
		switch name {
		default:
			fmt.Println("Error:", name, errrMap[name])
		case "":
			fmt.Println("Correct:", errrMap[name])
		}
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

/*



 */

func CombinationsTracked(set []string, n int, spell bool) (subsets [][]string) {

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
		if spell {
			fmt.Print("  ", subsetBits*100/total, "% of ", total, ": ", len(subsets), " New combination ", subset, " | Total:                  \r")
		}
	}
	if spell {
		fmt.Println("")
	}
	return subsets
}
