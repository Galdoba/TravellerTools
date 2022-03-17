package weapon

import (
	"fmt"
	"math"

	. "github.com/Galdoba/TravellerTools/pkg/weapon/components"
)

/*
1 Decide general type (pistol, rifle, ect...)
2 choose ammnition/powertype
3 choose receiver
4 choose receiver's mode of operation (breechloader, semi-automatic, ect...)
5 assign barrel lenght
6 assign furniture
7 choose feed device
8 add accessories that come as standard
9 total cost
10 total weight

*/

type WpnSheet struct {
	rcvr            *Receiver
	brl             *Barrel
	frn             *Furniture
	acc             *Accessoire
	penetration     int
	sigPhysical     int
	sigEmmision     int
	mishapThreshold int
	///////////////////
	name           string
	tl             int
	effectiveRange int
	damageDice     int
	damageMod      int
	weight         float64
	cost           int
	magazine       int
	magazineCost   int
	quickdraw      int
	autoValue      int
	recoil         int
	traits         []string
}

func New(instr ...int) (*WpnSheet, error) {
	w := WpnSheet{}
	err := fmt.Errorf("error was not adressed")
	if w.rcvr, err = NewReceiver(instr...); err != nil {
		return nil, err
	}
	if w.brl, err = NewBarrel(instr...); err != nil {
		return nil, err
	}
	if w.frn, err = NewFurniture(instr...); err != nil {
		return nil, err
	}
	if w.acc, err = NewAccessoires(instr...); err != nil {
		return nil, err
	}
	//calculations
	if err := w.calculateStats(); err != nil {
		return &w, err
	}
	fmt.Println("Stop")
	return &w, nil
}

func (w *WpnSheet) calculateStats(instr ...int) error {
	rcvrTotals := 1.0
	totalWeight := 0.0
	ammoCap := 0.0
	effectiveRange := 1.0
	auto := 0
	ammoCostPer100Rounds := 0.0
	addTrait := []string{}
	switch w.rcvr.ReceiverType() {
	default:
		return fmt.Errorf("calculateCost: stats were not assigned by ReceiverType")
	case RCVR_TYPE_HANDGUN:
		rcvrTotals = rcvrTotals * 175
		totalWeight += 0.8
		ammoCap = 10
		w.quickdraw = 4
	case RCVR_TYPE_ASSAULT_WEAPON:
		rcvrTotals = rcvrTotals * 300
		totalWeight += 2
		ammoCap = 20
		w.quickdraw = 2
	case RCVR_TYPE_LONGARM:
		rcvrTotals = rcvrTotals * 400
		totalWeight += 2.5
		ammoCap = 30
	case RCVR_TYPE_LIGHT_SUPPORT_WEAPON:
		rcvrTotals = rcvrTotals * 1500
		totalWeight += 5
		ammoCap = 40
		w.quickdraw = -4
	case RCVR_TYPE_HEAVY_WEAPON:
		rcvrTotals = rcvrTotals * 3000
		totalWeight += 10
		ammoCap = 50
		w.quickdraw = -8
	}
	if w.rcvr.IsGauss() {
		rcvrTotals = rcvrTotals * 2
		totalWeight = totalWeight * 1.25
		ammoCap = ammoCap * 3
	}

	switch w.rcvr.Mechanism() {
	default:
		return fmt.Errorf("calculateCost: cost was not modified by Mechanism")
	case MECHANISM_SINGLE_SHOT:
		rcvrTotals = rcvrTotals * 0.25
		ammoCap = 1
	case MECHANISM_REPEATER:
		rcvrTotals = rcvrTotals * 0.5
		ammoCap = ammoCap * 0.5
	case MECHANISM_SEMI_AUTOMATIC:
		//This is default mechanism
	case MECHANISM_BURST_CAPABLE:
		rcvrTotals = rcvrTotals * 1.1
		auto = 2
	case MECHANISM_FULLY_AUTOMATIC:
		rcvrTotals = rcvrTotals * 1.2
		auto = 3
	case MECHANISM_UNDERWATER:
		rcvrTotals = rcvrTotals * 2
		addTrait = append(addTrait, "Underwater")
	}
	/////////Calibre
	switch w.rcvr.AmmunitionType() {
	default:
		return fmt.Errorf("calculateCost: stats was not modified by calibre val %v", w.rcvr.AmmunitionType())
		////////////////////////////////////////////////////////////////////////////////////////
	case CALLIBRE_HANDGUN_Light:
		w.damageDice = 2
		ammoCostPer100Rounds = 60
		ammoCap = ammoCap * 1.2
		rcvrTotals = rcvrTotals * 0.8
		totalWeight = totalWeight * 0.75
	case CALLIBRE_HANDGUN_Medium:
		w.damageDice = 3
		w.damageMod = -3
		ammoCostPer100Rounds = 75
	case CALLIBRE_HANDGUN_Heavy:
		w.damageDice = 3
		w.damageMod = -1
		ammoCap = ammoCap * 0.8
		ammoCostPer100Rounds = 100
		rcvrTotals = rcvrTotals * 1.2
		totalWeight = totalWeight * 1.15
		addTrait = append(addTrait, "Bulky")
		////////////////////////////////////////////////////////////////////////////////////////
	case CALLIBRE_SMOOTHBORES_Small:
		rcvrTotals = rcvrTotals * 0.75
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_HANDGUN:
			addTrait = append(addTrait, "Bulky")
		case RCVR_TYPE_ASSAULT_WEAPON:
		case RCVR_TYPE_LONGARM:
		}
		w.damageDice = 3
		w.damageMod = -2
		ammoCostPer100Rounds = 100
		ammoCap = ammoCap * 1.4
		totalWeight = totalWeight * 0.6
		w.penetration = -1
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_HANDGUN:
			addTrait = append(addTrait, "Bulky")
		case RCVR_TYPE_ASSAULT_WEAPON:
		case RCVR_TYPE_LONGARM:
		}
	case CALLIBRE_SMOOTHBORES_Light:
		rcvrTotals = rcvrTotals * 0.75
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_HANDGUN:
			addTrait = append(addTrait, "Very Bulky")
		case RCVR_TYPE_ASSAULT_WEAPON:
			addTrait = append(addTrait, "Bulky")
		case RCVR_TYPE_LONGARM:
		}
		w.penetration = -1
		w.damageDice = 4
		w.damageMod = -4
		ammoCostPer100Rounds = 125
		ammoCap = ammoCap * 1.2
		totalWeight = totalWeight * 0.8
	case CALLIBRE_SMOOTHBORES_Standard:
		rcvrTotals = rcvrTotals * 0.75
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_ASSAULT_WEAPON:
			addTrait = append(addTrait, "Very Bulky")
		case RCVR_TYPE_LONGARM:
			addTrait = append(addTrait, "Bulky")
		}
		w.penetration = -1
		w.damageDice = 4
		ammoCostPer100Rounds = 150
	case CALLIBRE_SMOOTHBORES_Heavy:
		rcvrTotals = rcvrTotals * 0.75
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_LONGARM:
			addTrait = append(addTrait, "Very Bulky")
		}
		w.penetration = -1
		w.damageDice = 4
		w.damageMod = 4
		ammoCostPer100Rounds = 175
		ammoCap = ammoCap * 0.8
		totalWeight = totalWeight * 1.2
		////////////////////////////////////////////////////////////////////////////////////////
	case CALLIBRE_LONGARM_Rifle_Light:
		w.damageDice = 2
		ammoCostPer100Rounds = 40
		ammoCap = ammoCap * 1.2
		totalWeight = totalWeight * 0.6
	case CALLIBRE_LONGARM_Rifle_Intermediate:
		w.damageDice = 3
		ammoCostPer100Rounds = 50
		totalWeight = totalWeight * 0.8
	case CALLIBRE_LONGARM_Rifle_Battle:
		w.damageDice = 3
		w.damageMod = 3
		ammoCostPer100Rounds = 100
		ammoCap = ammoCap * 0.8
	case CALLIBRE_LONGARM_Rifle_Heavy:
		w.damageDice = 4
		ammoCostPer100Rounds = 250
		rcvrTotals = rcvrTotals * 1.25
		totalWeight = totalWeight * 1.1
		ammoCap = ammoCap * 0.6
	case CALLIBRE_LONGARM_Rifle_AntiMaterial:
		w.damageDice = 5
		ammoCostPer100Rounds = 1500
		rcvrTotals = rcvrTotals * 2.5
		ammoCap = ammoCap * 0.4
		totalWeight = totalWeight * 1.5
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_LIGHT_SUPPORT_WEAPON, RCVR_TYPE_HEAVY_WEAPON:
		}
		addTrait = append(addTrait, "Bulky")
	case CALLIBRE_LONGARM_Rifle_AntiMaterialHeavy:
		w.damageDice = 6
		ammoCostPer100Rounds = 3000
		rcvrTotals = rcvrTotals * 3.5
		ammoCap = ammoCap * 0.2
		totalWeight = totalWeight * 2
		addTrait = append(addTrait, "Very Bulky")
		switch w.rcvr.ReceiverType() {
		default:
			return fmt.Errorf("Not possiple to have %v with %v", Verbal(w.rcvr.ReceiverType()), Verbal(w.rcvr.AmmunitionType()))
		case RCVR_TYPE_LIGHT_SUPPORT_WEAPON, RCVR_TYPE_HEAVY_WEAPON:
		}
	case CALLIBRE_HANDGUN_BlackPowder:
		w.damageDice = 2
		w.damageMod = -3
		ammoCostPer100Rounds = 10
		w.penetration = -2
		addTrait = append(addTrait, "Slow Loader 8")
	case CALLIBRE_LONGARM_BlackPowder:
		w.damageDice = 3
		w.damageMod = -3
		ammoCostPer100Rounds = 25
		w.penetration = -2
		addTrait = append(addTrait, "Slow Loader 12")
	case CALLIBRE_SNUB:
		w.damageDice = 3
		w.damageMod = -3
		ammoCostPer100Rounds = 200
		w.penetration = -1
		addTrait = append(addTrait, "Zero-G")

		//if totalWeight > 2 = remove Bulky
	}

	switch w.rcvr.AmmunitionCapacity() {
	default:
		return fmt.Errorf("calculateCost: cost was not modified by Ammo Capacity")
	case AMMUNITION_CAPACITY_50PCT_LESS:
		rcvrTotals = rcvrTotals * 0.75
		ammoCap = ammoCap * 0.5
		totalWeight = totalWeight * 0.75
	case AMMUNITION_CAPACITY_40PCT_LESS:
		rcvrTotals = rcvrTotals * 0.8
		ammoCap = ammoCap * 0.6
		totalWeight = totalWeight * 0.8
	case AMMUNITION_CAPACITY_30PCT_LESS:
		rcvrTotals = rcvrTotals * 0.85
		ammoCap = ammoCap * 0.7
		totalWeight = totalWeight * 0.85
	case AMMUNITION_CAPACITY_20PCT_LESS:
		rcvrTotals = rcvrTotals * 0.90
		ammoCap = ammoCap * 0.8
		totalWeight = totalWeight * 0.90
	case AMMUNITION_CAPACITY_10PCT_LESS:
		rcvrTotals = rcvrTotals * 0.95
		ammoCap = ammoCap * 0.9
		totalWeight = totalWeight * 0.95
	case AMMUNITION_CAPACITY_STANDARD:
		//this is default option
	case AMMUNITION_CAPACITY_10PCT_MORE:
		rcvrTotals = rcvrTotals * 1.1
		ammoCap = ammoCap * 1.1
		totalWeight = totalWeight * 1.05
	case AMMUNITION_CAPACITY_20PCT_MORE:
		rcvrTotals = rcvrTotals * 1.2
		ammoCap = ammoCap * 1.2
		totalWeight = totalWeight * 1.1
	case AMMUNITION_CAPACITY_30PCT_MORE:
		rcvrTotals = rcvrTotals * 1.3
		ammoCap = ammoCap * 1.3
		totalWeight = totalWeight * 1.15
	case AMMUNITION_CAPACITY_40PCT_MORE:
		rcvrTotals = rcvrTotals * 1.4
		ammoCap = ammoCap * 1.4
		totalWeight = totalWeight * 1.2
	case AMMUNITION_CAPACITY_50PCT_MORE:
		rcvrTotals = rcvrTotals * 1.5
		ammoCap = ammoCap * 1.5
		totalWeight = totalWeight * 1.25
	}
	fmt.Println("RECIEVER TOTALS:", rcvrTotals)
	barelCost := 1.0
	switch w.brl.Length() {
	default:
		return fmt.Errorf("calculateCost: cost was not modified by Barrel Length")
	case BRL_len_HANDGUN:
		barelCost = rcvrTotals * 0.15
	}
	fmt.Println("BARREL TOTALS:", barelCost)
	costInt := int(math.Round(rcvrTotals + barelCost))
	fmt.Println("----------")
	w.cost = costInt
	w.weight = math.Round(totalWeight*100) / 100
	w.magazine = int(math.Ceil(ammoCap))
	w.magazineCost = (int(math.Ceil(ammoCap)*math.Ceil(ammoCostPer100Rounds)) / 100) + w.cost/100
	w.autoValue = auto
	w.effectiveRange = int(effectiveRange)
	w.traits = addTrait
	fmt.Println("Weapon TOTALS:", w.cost)
	return nil
}

//DesignWorksheet - будет возвращать в виде таблицы покомпонентный расклад оружия
func (wp *WpnSheet) DesignWorksheet() string {
	fmt.Println("[Weapon   TL   Range   Damage   Kg   Cost   Magazine   Magazine Cost   Traits]")
	return ""
}

//Summary - будет возвращать в виде таблицы выжимку по оружию
func (wp *WpnSheet) Summary() string {
	return fmt.Sprintf("TL: %v Range: %v Damage: %vD+(%v) Kg: %v Cost: Cr%v Magazine: %v Magazine Cost: Cr%v Traits: %v", wp.tl, wp.effectiveRange,
		wp.damageDice, wp.damageMod, wp.weight, wp.cost, wp.magazine, wp.magazineCost, wp.traits)
}
