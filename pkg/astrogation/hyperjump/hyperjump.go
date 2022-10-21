package hyperjump

import (
	"strconv"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/utils"
)

type hyperJump struct {
	effA          int
	effE          int
	timevar       int
	distvar       int
	distance      int
	hours         int
	badJumpA      bool
	badJumpE      bool
	precipitation bool
	misjumpEff    int
	outcome       string
	jumpStatus    string
	misjumpStatus string
	roller        *dice.Dicepool
}

//HyperJump - модуль гиперпрыжка в MGT2 Traveller Companion (p. 141-143)
type HyperJump interface {
	//Stringer
}

// func Test() {
// 	hj := New(TrvCore.Flux(), TrvCore.Flux())
// 	fmt.Println(hj.Report())
// 	fmt.Println(hj.Outcome())
// }

func New(effA, effE int) *hyperJump {
	hj := hyperJump{}
	hj.effA = effA
	hj.effE = effE
	hj.misjumpEff = hj.effA + hj.effE
	hj.jumpStatus = "Normal"
	hj.misjumpStatus = "NO"
	hj.roller = dice.New()
	//fmt.Println(hj.effA, hj.effE)
	hj.misjumpEffects()
	hj.timeVariance()
	hj.distVariance()
	hj.badJumpEffects()

	//fmt.Println(hj.outcome)

	//fmt.Println("---------------")

	return &hj
}

func (hj hyperJump) Report() string {
	rep := ""
	rep += "--------------------------------------------------------------------------------" + "\n"
	rep += "Distance Variance: " + strconv.Itoa(hj.distvar) + "\n"
	rep += "    Time Variance: " + strconv.Itoa(hj.timevar) + "\n"
	rep += "      Jump Status: " + hj.jumpStatus + "\n"
	rep += "          MisJump: " + hj.misjumpStatus + "\n"
	rep += "--------------------------------------------------------------------------------" + "\n"
	if hj.roller.Roll("2d6").Sum() >= 9 {
		rep += "Hyperjump Event happened - check Campaign Guide!" + "\n"
		rep += "--------------------------------------------------------------------------------" + "\n"
	}
	return rep
}

func (hj hyperJump) Outcome() string {
	return hj.outcome
}

func (hj *hyperJump) distVariance() {
	variance := 100
	r := hj.roller.Roll("2d6").DM(hj.effA).Sum()
	r = utils.BoundInt(r, 2, 12)
	hj.distvar = 12 - r
	//fmt.Println("888888", hj.distvar)

	switch r {
	case 2:
		variance += 10 - hj.roller.Roll("3d6").Sum()
		hj.badJumpA = true
	case 3:
		variance += 10 - hj.roller.Roll("2d6").Sum()
		hj.badJumpA = true
	case 4:
		variance += 5 - hj.roller.Roll("1d6").Sum()
		hj.badJumpA = true
	case 5:
		variance += (hj.roller.Roll("2d6").Sum() * 10)
		hj.badJumpA = true
	case 6:
		variance += (hj.roller.Roll("2d6").Sum() * 5)
	case 7:
		variance += hj.roller.Roll("4d6").Sum()
	case 8:
		variance += hj.roller.Roll("3d6").Sum()
	case 9:
		variance += hj.roller.Roll("2d6").Sum()
	case 10:
		variance += hj.roller.Roll("1d6").Sum()
	case 11:
		variance += (hj.roller.Roll("1d3").Sum())
	default:
		variance = 100

	}
	if variance < 100 {
		hj.precipitation = true
	}
	hj.distance += variance
	hj.outcome += " at " + strconv.Itoa(hj.distance) + " diameters from intended planet.\n"
	hj.outcome += "--------------------------------------------------------------------------------" + "\n"

}

func (hj *hyperJump) timeVariance() {
	hj.hours = hj.hours + 168
	variance := 0
	r := hj.roller.Roll("2d6").DM(hj.effE).Sum()
	r = utils.BoundInt(r, 2, 12)
	hj.timevar = 12 - r
	switch r {
	case 2:
		variance = hj.roller.Roll("16d6").Sum()
		hj.badJumpE = true
	case 3:
		variance = hj.roller.Roll("10d6").Sum()
		hj.badJumpE = true
	case 4:
		variance = hj.roller.Roll("8d6").Sum()
		hj.badJumpE = true
	case 5:
		variance = hj.roller.Roll("6d6").Sum()
		hj.badJumpE = true
	case 6:
		variance = hj.roller.Roll("5d6").Sum()
	case 7:
		variance = hj.roller.Roll("4d6").Sum()
	case 8:
		variance = hj.roller.Roll("3d6").Sum()
	case 9:
		variance = hj.roller.Roll("2d6").Sum()
	case 10:
		variance = hj.roller.Roll("1d6").Sum()
	case 11:
		variance = hj.roller.Roll("1d3").Sum()
	case 12:
		variance = 0

	}
	r2 := hj.roller.Roll("1d2").Sum()
	if r2 == 2 {
		variance = variance * -1
	}
	hj.hours = hj.hours + variance
	if hj.hours < 0 {
		realHours := hj.hours * -1
		hj.outcome += "Time paradox occured! Ship spent " + strconv.Itoa(realHours) + " hours (or " + formatTime(realHours) + ") in a jumpbbuble and than emerged to normal space few seconds after the jump"
	}
	if hj.hours == 0 {
		hj.outcome += "Ship emerged to normal space few seconds after the jump"
	}
	if hj.hours > 0 {
		hj.outcome += "After " + strconv.Itoa(hj.hours) + " hours (or " + formatTime(hj.hours) + ") ship emerged from jumpspace"
	}

}

func formatTime(hours int) string {
	hh := hours % 24
	dd := hours / 24
	if hh > 11 {
		dd++
	}
	rep := strconv.Itoa(dd) + " day"
	if dd != 1 {
		rep += "s"
	}
	return rep
	// rep += " and " + strconv.Itoa(hh) + " hour"
	// if hh != 1 {
	// 	rep += "s"
	// }
	// return rep
}

func (hj *hyperJump) badJumpEffects() {
	//fmt.Println(hj.badJumpA, hj.badJumpE)
	if hj.badJumpA || hj.badJumpE || hj.jumpStatus == "Minor" {
		hj.jumpStatus = "Bad Jump"
		hj.outcome += "Everyone aboard the vessel must make an END and INT check. One of these checks is at Routine (6+) difficulty and the other at Difficult (10+) difficulty level. A Traveller can choose which check is taken at each difficulty level.\n"
		hj.outcome += "\nThe END check determines if physical effects are present. These include nausea and possibly vomiting, plus often a blinding headache. If the END check is failed, all checks the Traveller attempts are subject to a DM equal to the Effect of the failure for " + hj.roller.Roll("2d6").SumStr() + " hours after entry into jump, and again after emergence.\n"
		hj.outcome += "An Effect of -6 or worse indicates the Traveller is incapacitated; unconscious or wishing he was, for " + strconv.Itoa(hj.roller.Roll("2d6").Sum()*30) + " minutes after which a DM-6 applies on all checks for the following " + hj.roller.Roll("4d6").SumStr() + " hours.\n"
		hj.outcome += "\nThe INT check determines if psychological effects are present. These typically include unease, irritability and paranoia, but in some cases can lead to a complete breakdown. Anyone who fails the INT check will be irritable, nervous and generally out of sorts for the whole duration of the jump. This manifests as difficulty in concentrating as well as a tendency to be on edge which often makes interactions with other Travellers unpleasant."
		hj.outcome += " DM-2 applies to all checks associated with mental or interpersonal activities. The Traveller will be visibly on edge, and may appear to be behaving suspiciously or in a paranoid fashion. Memory lapses, covering a few minutes to an hour or two, are also possible.\n"
		hj.outcome += "An Effect of -6 or worse indicates the Traveller is suffering from serious mental effects. These manifest as acute paranoia, blackouts and hallucinations. A Traveller in this state might harm themselves or someone else, or take a potentially dangerous action such as locking a hatch that the rest of the crew need to access a critical area.\n"
		hj.outcome += "Psychological effects last throughout the jump and for " + hj.roller.Roll("1d6").SumStr() + " days afterward.\n"
	}
	if (hj.badJumpA && hj.badJumpE) || hj.jumpStatus == "Serious" {
		hj.jumpStatus = "Very Bad Jump"
		hj.outcome += "\nAditional effects: "
		dm := -4
		if hj.precipitation {
			dm = -2
		}
		if hj.misjumpEff <= 0 {
			dm = 0
		}
		r := hj.roller.Roll("2d6").DM(dm).Sum()
		switch r {
		default:
			if r <= 2 {
				hj.outcome += "None"
			}
			if r >= 13 {
				hj.outcome += "Severe jumpspace intrusions occur, jump drive destroyed upon emergence"
				hj.outcome += "A severe intrusion consumes " + hj.roller.Roll("2d6").DM(10).SumStr() + "% of the ship’s Hull per day."
			}
		case 3, 4, 5:
			hj.outcome += "Jump drive requires lengthy recalibration (taking " + hj.roller.Roll("2d6").SumStr() + " days) after emergence"
		case 6, 7:
			hj.outcome += "Jump drive requires minor repairs after emergence"
		case 8, 9:
			hj.outcome += "Jump drive requires major repairs after emergence"
		case 10, 11, 12:
			hj.outcome += "Jumpspace intrusions occur, jump drive destroyed upon emergence. "
			hj.outcome += "Ship suffering jumpspace intrusions suffers damage equal to " + hj.roller.Roll("2d6").DM(-2).SumStr() + "% of its Hull per day"

		}
		hj.outcome += "\n"
	}

}

func (hj *hyperJump) misjumpEffects() {
	if hj.effA+hj.effE > 0 {
		return
	}
	hj.misjumpStatus = "Minor"
	//hj.badJumpE = true
	//hj.badJumpA = false
	if hj.effE < 0 && hj.effA < 0 {
		hj.misjumpStatus = "Serious"
		//hj.badJumpE = true
		//hj.badJumpA = true
	}
	r := hj.roller.Roll("2d6").Sum() + hj.effA + hj.effE
	r = utils.BoundInt(r, 2, 12)
	switch r {
	case 2:
		hj.outcome = "Ship is lost in jumpspace or emerges as fragments and subatomic particles.\n"
		//fmt.Println("Ship is lost in jumpspace or emerges as fragments and subatomic particles.")
		return
	case 3, 4:
		hj.outcome += "Ship misjumps " + strconv.Itoa(hj.roller.Roll("1d6").Sum()*hj.roller.Roll("1d6").Sum()) + " parsecs in a random direction. Its jump drive is completely wrecked upon emergence and the passengers and crew risk lasting psychological effects."
	case 5, 6:
		hj.outcome += "Ship misjumps " + hj.roller.Roll("2d6").SumStr() + " parsecs in a random direction. Its jump drive is severely damaged upon emergence."
	case 7, 8:
		hj.outcome += "Ship misjumps " + hj.roller.Roll("1d6").SumStr() + " parsecs in a random direction"
	case 9, 10:
		days := hj.roller.Roll("2d6").Sum()
		changed := "increased"
		if hj.roller.Roll("1d2").Sum() == 2 {
			changed = "decreased"
			days = days * -1
		}
		num := hj.roller.Roll("1d3").SumStr()
		hj.hours = hj.hours + 24*days
		hj.outcome += "Jump duration " + changed + " by " + strconv.Itoa(days) + " days. Jump drive requires extensive recalibration (" + num + " days of work) but no repairs."
	case 11, 12:
		num := 100 * hj.roller.Roll("2d6").Sum()
		hj.distance += num
		hj.outcome += "Very rough jump, emergence is at " + strconv.Itoa(num) + " diameters farther from the target world."
	}
	hj.outcome += "\n"
}
