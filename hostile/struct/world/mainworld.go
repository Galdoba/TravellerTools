package world

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	gear_vacc = "Vacc Suit"
	gear_mask = "Survival Mask"
	gear_airs = "Air Supply/Vacc Suit"
	taint_cd  = "Carbone Dioxide"
	taint_m   = "Methane"
	taint_n   = "Nitrogen"
	taint_o   = "Oxygen"
	taint_hg  = "Hydrogen"
	taint_s   = "Sulfur"
	taint_a   = "Ammonia"
	taint_cl  = "Clorine"
	taint_fl  = "Fluorine"
	taint_ht  = "High Temperature"
)

type World struct {
	Name                string `json:"Name"`
	Hex                 string `json:"Hex"`
	UWP                 string `json:"UWP"`
	uwp                 map[string]ehex.Ehex
	Gravity             float64 `json:"Surface Gravity"`
	Radius              float64 `json:"Earth Radii"`
	Mass                float64 `json:"Earth Mass"`
	Density             string  `json:"Atmospheric Density"`
	Pressure            float64 `json:"Atmospheric Pressure"`
	Survival_Gear_Req   string  `json:"Survival Gear Required"`
	Taint               string  `json:"Atmospheric Taint,omitempty"`
	Trade_Codes         string  `json:"Trade Codes,omitempty"`
	Gas_Gigant_Presence bool    `json:"Gas Gigant Present"`
}

func (w *World) Feed(data map[string]ehex.Ehex) {
	w.uwp = data
}

func (w *World) GenerateDetails(dice *dice.Dicepool) error {
	if dice.Sroll("2d6") <= 9 {
		w.Gas_Gigant_Presence = true
	}
	switch w.uwp["size"].Value() {
	case 1:
		w.Gravity = rollFlBetween(dice, 0.05, 0.15)
		w.Radius = rollFlBetween(dice, 0.125, 0.25)
		w.Mass = rollFlBetween(dice, 0.002, 0.016)
	case 2:
		w.Gravity = rollFlBetween(dice, 0.15, 0.25)
		w.Radius = rollFlBetween(dice, 0.25, 0.375)
		w.Mass = rollFlBetween(dice, 0.016, 0.053)
	case 3:
		w.Gravity = rollFlBetween(dice, 0.15, 0.25)
		w.Radius = rollFlBetween(dice, 0.375, 0.5)
		w.Mass = rollFlBetween(dice, 0.053, 0.125)
	case 4:
		w.Gravity = rollFlBetween(dice, 0.35, 0.45)
		w.Radius = rollFlBetween(dice, 0.5, 0.625)
		w.Mass = rollFlBetween(dice, 0.125, 0.244)
	case 5:
		w.Gravity = rollFlBetween(dice, 0.45, 0.7)
		w.Radius = rollFlBetween(dice, 0.625, 0.75)
		w.Mass = rollFlBetween(dice, 0.244, 0.422)
	case 6:
		w.Gravity = rollFlBetween(dice, 0.7, 0.9)
		w.Radius = rollFlBetween(dice, 0.75, 0.875)
		w.Mass = rollFlBetween(dice, 0.422, 0.67)
	case 7:
		w.Gravity = rollFlBetween(dice, 0.9, 1)
		w.Radius = rollFlBetween(dice, 0.875, 1)
		w.Mass = rollFlBetween(dice, 0.67, 1)
	case 8:
		w.Gravity = rollFlBetween(dice, 1, 1.2)
		w.Radius = rollFlBetween(dice, 1, 1.125)
		w.Mass = rollFlBetween(dice, 1, 1.25)
	case 9:
		w.Gravity = rollFlBetween(dice, 1.2, 1.5)
		w.Radius = rollFlBetween(dice, 1.125, 1.25)
		w.Mass = rollFlBetween(dice, 1.25, 2.5)
	case 10:
		w.Gravity = rollFlBetween(dice, 1.5, 1.8)
		w.Radius = rollFlBetween(dice, 1.25, 1.375)
		w.Mass = rollFlBetween(dice, 2.5, 3.75)
	case 11:
		w.Gravity = rollFlBetween(dice, 1.8, 2)
		w.Radius = rollFlBetween(dice, 1.375, 1.5)
		w.Mass = rollFlBetween(dice, 3.75, 5)
	case 12:
		w.Gravity = rollFlBetween(dice, 2, 2.2)
		w.Radius = rollFlBetween(dice, 1.5, 1.625)
		w.Mass = rollFlBetween(dice, 5, 6.25)
	case 13:
		w.Gravity = rollFlBetween(dice, 2.2, 2.3)
		w.Radius = rollFlBetween(dice, 1.625, 1.75)
		w.Mass = rollFlBetween(dice, 6.25, 7.5)
	case 14:
		w.Gravity = rollFlBetween(dice, 2.3, 2.5)
		w.Radius = rollFlBetween(dice, 1.75, 1.875)
		w.Mass = rollFlBetween(dice, 7.5, 8.75)
	case 15:
		w.Gravity = rollFlBetween(dice, 2.5, 2.6)
		w.Radius = rollFlBetween(dice, 1.875, 2)
		w.Mass = rollFlBetween(dice, 8.75, 10)
	case 16:
		w.Gravity = rollFlBetween(dice, 2.6, 3.5)
		w.Radius = rollFlBetween(dice, 2, 6)
		w.Mass = rollFlBetween(dice, 10, 17.2)
	}
	switch w.uwp["size"].Value() {
	case 0:
		w.Survival_Gear_Req = gear_vacc
		w.Density = "None"
	case 1:
		w.Pressure = rollFlBetween(dice, 0.001, 0.009)
		w.Survival_Gear_Req = gear_vacc
		w.Density = "Trace"
	case 2:
		w.Pressure = rollFlBetween(dice, 0.1, 0.42)
		w.Survival_Gear_Req = gear_mask
		w.Taint = taint1(dice)
		w.Density = "Very Thin"
	case 3:
		w.Pressure = rollFlBetween(dice, 0.1, 0.42)
		w.Survival_Gear_Req = gear_mask
		w.Density = "Very Thin"
	case 4:
		w.Pressure = rollFlBetween(dice, 0.43, 0.7)
		w.Survival_Gear_Req = gear_mask
		w.Taint = taint1(dice)
		w.Density = "Thin"
	case 5:
		w.Pressure = rollFlBetween(dice, 0.43, 0.7)
		w.Density = "Thin"
	case 6:
		w.Pressure = rollFlBetween(dice, 0.71, 1.49)
		w.Density = "Standard"
	case 7:
		w.Pressure = rollFlBetween(dice, 0.71, 1.49)
		w.Survival_Gear_Req = gear_mask
		w.Taint = taint1(dice)
		w.Density = "Standard"
	case 8:
		w.Pressure = rollFlBetween(dice, 1.5, 2.49)
		w.Survival_Gear_Req = gear_mask
		w.Density = "Dense"
	case 9:
		w.Pressure = rollFlBetween(dice, 1.5, 2.49)
		w.Survival_Gear_Req = gear_mask
		w.Taint = taint2(dice)
		w.Density = "Dense"
	case 10:
		w.Pressure = rollFlBetween(dice, 0.5, 2.49)
		w.Survival_Gear_Req = gear_airs
		w.Taint = taint3(dice)
		w.Density = "Exotic"
	case 11:
		w.Pressure = rollFlBetween(dice, 0.5, 2.49)
		w.Survival_Gear_Req = gear_vacc
		w.Taint = taint3(dice)
		w.Density = "Corrosive"
	case 12:
		w.Pressure = rollFlBetween(dice, 0.5, 2.49)
		w.Survival_Gear_Req = gear_airs
		w.Taint = taint3(dice)
		w.Density = "Irritant"

	}

	w.Gravity = float64(int(w.Gravity*1000)) / 1000
	w.Radius = float64(int(w.Radius*1000)) / 1000
	w.Mass = float64(int(w.Mass*1000)) / 1000

	return nil
}

func taint1(dice *dice.Dicepool) string {
	tnt := []string{
		taint_cd,
		taint_s,
		taint_fl,
		taint_n,
		taint_cl,
	}
	return tnt[dice.Sroll("1d5-1")]
}
func taint2(dice *dice.Dicepool) string {
	tnt := []string{
		taint_m,
		taint_a,
		taint_hg,
	}
	return tnt[dice.Sroll("1d3-1")]
}

func taint3(dice *dice.Dicepool) string {
	tnt := []string{
		taint_m,
		taint_a,
		taint_hg,
		taint_o,
		taint_ht,
	}
	return tnt[dice.Sroll("1d5-1")]
}

func rollFlBetween(dice *dice.Dicepool, min, max float64) float64 {
	//r := 0.001
	val := 0.0
	maxDelta := max - min
	mds := fmt.Sprintf("%v", int(maxDelta*1000))
	delta := 0.0
	match := false
	for !match {
		for val < min {
			r1 := dice.Sroll("1d" + mds)
			delta = float64(r1) / 1000
			val = val + delta
		}
		if val > max {
			val -= delta
		}
		if val >= min && val <= max {
			match = true
		}
	}
	return val
}

func NewMainWorld() *World {
	w := World{}
	w.uwp = make(map[string]ehex.Ehex)
	return &w
}

func (wrld *World) MarshalJson(path string) error {
	bt, err := json.MarshalIndent(wrld, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		return err
	}
	w, err := f.Write(bt)
	if err != nil {
		return err
	}
	if w < 1 {
		return fmt.Errorf("nothing was written")
	}
	fmt.Println(string(bt))
	return nil
}

func UnmarshalJson(path string) (*World, error) {

	wrld := &World{}
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bt, wrld); err != nil {
		return nil, err
	}

	return wrld, nil
}
