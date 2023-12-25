package wbh

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	orbitns "github.com/Galdoba/TravellerTools/mgt2/wbh/orbits"
	"github.com/Galdoba/TravellerTools/mgt2/wbh/star"
	"github.com/Galdoba/TravellerTools/mgt2/wbh/worlds"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	defauiltVal = iota
	GenerationMethodSpecial
	GenerationMethodUnusual
	GenerationMethodContinuation
	GenerationMethodExpanded
	TypeVariantTraditional
	TypeVariantRealistic

	designationPrimary          = "Aa"
	designationPrimaryCompanion = "Ab"
	designationClose            = "Ba"
	designationCloseCompanion   = "Bb"
	designationNear             = "Ca"
	designationNearCompanion    = "Cb"
	designationFar              = "Da"
	designationFarCompanion     = "Db"
	determinationPrimary        = "Prime"
	determinationRandom         = "Random"
	determinationLesser         = "Lesser"
	determinationSibling        = "Sibling"
	determinationTwin           = "Twin"
	determinationOther          = "Other"
)

type StarSystem struct {
	starGenerationMethod   int                  `json:"Star Generation Method"`
	systemGenerationMethod int                  `json:"System Generation Method"`
	TypeTableVariant       int                  `json:"Type Table Variant"`
	Star                   map[string]star.Star `json:"Stars"`
	Age                    float64              `json:"System Age"`
	WorldType              map[string]int       `json:"World Type map"`
}

func NewStarSystem(dice *dice.Dicepool, starGenerationMethod, tableVariant, sysGenMet int) (*StarSystem, error) {
	fmt.Println(dice, starGenerationMethod, tableVariant, sysGenMet)
	ss := StarSystem{}
	ss.systemGenerationMethod = sysGenMet
	switch ss.systemGenerationMethod {
	case GenerationMethodExpanded, GenerationMethodContinuation:
	default:
		return &ss, fmt.Errorf("systemGenerationMethod unknown (%v)", starGenerationMethod)
	}
	ss.starGenerationMethod = starGenerationMethod
	switch ss.starGenerationMethod {
	case GenerationMethodUnusual, GenerationMethodSpecial:
	default:
		return &ss, fmt.Errorf("starGenerationMethod unknown (%v)", starGenerationMethod)
	}
	ss.Star = make(map[string]star.Star)
	primary, err := star.New(dice, tableVariant, starGenerationMethod, designationPrimary, determinationPrimary)
	if err != nil {
		return &ss, err
	}
	ss.Star[designationPrimary] = primary
	designations := star.DefineStarPresence(ss.Star[designationPrimary], dice)
	for _, desig := range designations {
		if _, ok := ss.Star[desig]; ok {
			continue
		}
		determ, context := star.DefineStarDetermination(primary, desig, dice)
		star, err := star.New(dice, tableVariant, starGenerationMethod, desig, determ, ss.Star[context])
		if err != nil {
			return &ss, fmt.Errorf("secondary star %v creation: %v", desig, err.Error())
		}
		ss.Star[desig] = star
	}
	ss.ageResetIfRequired(dice)

	//TODO: dm for eccentricity
	dm := 0
	for _, desig := range star.DesignationCodes() {
		if _, ok := ss.Star[desig]; !ok {
			continue
		}
		st := ss.Star[desig]
		orbN, err := orbitns.DetermineStarOrbit(dice, desig)
		if err != nil {
			return &ss, fmt.Errorf("orbitns.DetermineStarOrbit: %v", err.Error())
		}
		st.Orbit = orbitns.New(orbN)
		st.Orbit.DetermineEccentrisity(dice, dm)
		//st.normalizeValues()
		ss.Star[desig] = st
	}
	ss.CalculateOP()
	bt, err := json.MarshalIndent(ss, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bt))
	//ss.PlaceWorlds(dice)

	return &ss, nil
}

func systemStars(ss *StarSystem) []star.Star {
	stars := []star.Star{}
	for _, desig := range star.DesignationCodes() {
		if st, ok := ss.Star[desig]; ok {
			stars = append(stars, st)
		}
	}
	return stars
}

func (ss *StarSystem) CalculateOP() {
	for _, code := range []string{"Ba", "Bb", "Ca", "Cb", "Da", "Db"} {
		switch {
		case strings.Contains(code, "Aa"):
			continue
		case strings.Contains(code, "a"):
			far := strings.TrimSuffix(code, "a")
			if _, ok := ss.Star[far+"a"]; ok {
				star1 := ss.Star[far+"a"]
				star2 := ss.Star[far+"b"]
				star3 := ss.Star["Aa"]
				star4 := ss.Star["Ab"]
				au := (star.AUof(star2) + 2*star.AUof(star1)) / 2
				m1 := star1.Mass + star2.Mass
				m2 := star3.Mass + star4.Mass
				orbPer := orbitns.CalculateOrbitalPeriod(au, m1, m2)
				//star1.Orbit.Period = orbPer
				ss.Star[far+"a"].Orbit.Period = orbPer
			}
		case strings.Contains(code, "b"):
			far := strings.TrimSuffix(code, "b")
			if _, ok := ss.Star[far+"b"]; ok {
				star1 := ss.Star[far+"a"]
				star2 := ss.Star[far+"b"]
				au := star.AUof(star2)
				m1 := star1.Mass
				m2 := star2.Mass
				orbPer := orbitns.CalculateOrbitalPeriod(au, m1, m2)
				ss.Star[far+"b"].Orbit.Period = orbPer
			}
		}
	}

}

func (ss *StarSystem) ageResetIfRequired(dice *dice.Dicepool) {
	switch ss.Star["Aa"].Class {
	case star.ClassIa, star.ClassIb, star.ClassII, star.ClassIII, star.ClassIV, star.ClassV, star.ClassVI, star.ClassBD:
		for _, v := range ss.Star {
			switch v.Class {
			case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole:
				primary := ss.Star["Aa"]
				primary.Age = v.Age
				ss.Star["Aa"] = primary

				// primary.Age = starFinalAge(v.Mass, dice)
				// if primary.Age < was {
				// 	fmt.Println("set new age")
				// 	primary.Age = was
				// }
				// if primary.Age > 13.5 {
				// 	fmt.Println("set age border", primary)

				// 	primary.Age = 13.5
				// }
				// ss.Star["Aa"] = primary
			}
		}
	}
	ss.Age = ss.Star["Aa"].Age
}

func (ss *StarSystem) String() string {
	prf := fmt.Sprintf("%v-", len(ss.Star))
	for _, desig := range star.DesignationCodes() {
		if st, ok := ss.Star[desig]; ok {
			switch desig {
			case "Aa":
				prf += star.ShortStarDescription(st) //star.stType + star.subType + " " + star.Class
				prf += "-" + fmt.Sprintf("%v", st.Mass)
				prf += "-" + fmt.Sprintf("%v", st.Diameter)
				prf += "-" + fmt.Sprintf("%v", st.Luminocity)
				prf += "-" + fmt.Sprintf("%v", st.Age)
			default:
				prf += ":" + fmt.Sprintf("%v", desig)
				prf += "-" + fmt.Sprintf("%v", st.Orbit.OrbitNum)
				prf += "-" + fmt.Sprintf("%v", st.Orbit.Eccentricity)
				prf += "-" + star.ShortStarDescription(st)
				prf += "-" + fmt.Sprintf("%v", st.Mass)
				prf += "-" + fmt.Sprintf("%v", st.Diameter)
				prf += "-" + fmt.Sprintf("%v", st.Luminocity)
				//prf += "&" + fmt.Sprintf("%v", st.MAO)
			}
		}
	}
	prf = strings.TrimPrefix(prf, "1-")
	return prf
}

func (ss *StarSystem) PlaceWorlds(dice *dice.Dicepool) {
	ss.WorldType = make(map[string]int)
	ss.WorldType["GG"] = worlds.GasGigantsQuantity(dice, ss.Star)
	ss.WorldType["BELTS"] = worlds.PlanetoidBeltsQuantity(dice, ss.Star, ss.WorldType["GG"])
	ss.WorldType["ROCKY"] = worlds.TerrestialPlanetsQuantity(dice, ss.Star)
	ss.WorldType["TOTAL"] = ss.WorldType["GG"] + ss.WorldType["BELTS"] + ss.WorldType["ROCKY"]
	ss.Star = star.CalculateAllowableOrbits(ss.Star)
	ss.Star = star.DefineHZCO(ss.Star)
	allowed := star.AllocateWorldlimitsByStars(ss.WorldType["TOTAL"], ss.Star)
	hzco := []float64{} //ss.Star["Aa"].HZCO
	for _, stCode := range []string{"Aa", "Ba", "Ca", "Da"} {
		if _, ok := ss.Star[stCode]; !ok {
			hzco = append(hzco, -1)
		} else {
			hzco = append(hzco, ss.Star[stCode].HZCO)
		}
	}
	fmt.Println("hzco", hzco)
	sstmBaseNmbr := systemBaseNumber(ss, dice, ss.WorldType["TOTAL"])
	ss.WorldType["sbn"] = sstmBaseNmbr
	fmt.Println("sstmBaseNmbr", sstmBaseNmbr)
	fmt.Println("allowed", allowed)
	sbnType := systemBaselineNumType(sstmBaseNmbr, ss.WorldType["TOTAL"], ss.systemGenerationMethod)
	fmt.Println("sbnType", sbnType)
	baselineOrbits := []float64{-1.0, -1.0, -1.0, -1.0}
	id2Code := make(map[int]string)
	id2Code[0] = "Aa"
	id2Code[1] = "Ba"
	id2Code[2] = "Ca"
	id2Code[3] = "Da"
	switch sbnType {
	case "a":
		baselineOrbits = baselineOrbit_A(dice, hzco)
	case "b":
		/*first orbit is beyond a point either based on the primary star(s) minimum Orbit#, its
		  HZCO or MAO, whichever is greater.*/
		for i, v := range hzco {
			if v < 0 {
				continue
			}
			if ss.Star[id2Code[i]].MAO > v {
				hzco[i] = ss.Star[id2Code[i]].MAO
			}
		}
		baselineOrbits = baselineOrbit_B(dice, hzco, sstmBaseNmbr, ss.WorldType["TOTAL"])
	case "c":
		baselineOrbits = baselineOrbit_C(dice, hzco, sstmBaseNmbr, ss.WorldType["TOTAL"])
		for i, blo := range baselineOrbits {
			/* if negative treat the baseline Orbit as the HZCO – 0.1 but no
			lower than the primary star’s MAO + the primary star’s total worlds x 0.01.*/
			if blo <= 0 && hzco[i] > 0 {
				newBlo := hzco[i] - 0.1
				minBlo := ss.Star[id2Code[i]].MAO + (float64(allowed[i]) * 0.01)
				if newBlo <= minBlo {
					newBlo = minBlo
				}
				baselineOrbits[i] = newBlo
			}
		}
	case "d":
		panic("continuation Method not implemented")
	}
	fmt.Println("baselineOrbits", baselineOrbits)
	fmt.Println("hzco", hzco)
	fmt.Println(ss.String())
	ss.WorldType["EMPTY"] = emptyOrbits(dice)
	availableOrbits := []float64{}
	maoSl := []float64{}
	for _, st := range []string{"Aa", "Ba", "Ca", "Da"} {
		if _, ok := ss.Star[st]; !ok {
			availableOrbits = append(availableOrbits, -1)
			maoSl = append(maoSl, -1)
			continue
		}
		availableOrbits = append(availableOrbits, ss.Star[st].AvailableOrbits.Sum())
		maoSl = append(maoSl, ss.Star[st].MAO)
	}
	fmt.Println("mao", maoSl)
	spread := determineSystemSpread(baselineOrbits, maoSl, sstmBaseNmbr, allowed, ss.WorldType["TOTAL"], availableOrbits)
	fmt.Println("spread", spread)
	ss.placeOrbitsN(dice, spread, baselineOrbits, allowed, sbnType)
	anomalyMap := generateAnomalyes(dice)
	ss.WorldType["TOTAL"] = ss.WorldType["TOTAL"] + len(anomalyMap)
	fmt.Println(ss.WorldType)
	for _, v := range anomalyMap {
		if err := ss.placeBodyAndConfirm(dice, v); err != nil {
			panic("anomaly " + err.Error())
		}
	}
	for _, body := range []string{"EMPTY", "GG", "BELTS", "ROCKY"} {
		for i := 0; i < ss.WorldType[body]; i++ {
			if err := ss.placeBodyAndConfirm(dice, body); err != nil {
				panic(body + " " + err.Error())
			}
		}
	}
	for code, str := range ss.Star {
		for v, o := range str.ChildOrbit {
			if o.AsignedBody == "" {
				delete(ss.Star[code].ChildOrbit, v)
				continue
			}
			ecDM := 0
			switch o.AsignedBody {
			case "random":
				fmt.Println("TODO: place world to another location")
			case "eccentric":
				ecDM = 5
			case "inclined":
				ecDM = 2
			case "retrograde":
				ecDM = 2
			case "trojan":
				fmt.Println("TODO: trojan world must be dealed before this spot")
			}
			ss.Star[code].ChildOrbit[v].DetermineEccentrisity(dice, ecDM)
			fmt.Println(code, v, ss.Star[code].ChildOrbit[v])
		}
	}

	/*
		//SystemSpread*/

	//PlaceOrbits
	//AddAnomalousPlanets
	//PlaceWorlds
	//DetermineEccentricity
}

func (ss *StarSystem) placeBodyAndConfirm(dice *dice.Dicepool, body string) error {
	if body == "" {
		return fmt.Errorf("body must be named")
	}
	starCode := ""
	freeOrbits := make(map[string][]float64)
	freeInner := make(map[string][]float64)
	freeOuter := make(map[string][]float64)
	codes := []string{"Aa", "Ba", "Ca", "Da"}
	for _, code := range codes {
		if _, ok := ss.Star[code]; !ok {
			continue
		}
		for kStr, v := range ss.Star[code].ChildOrbit {
			k, _ := strconv.ParseFloat(kStr, 64)
			if v.AsignedBody == "" {
				if k <= ss.Star[code].HZCO {
					freeInner[code] = append(freeInner[code], k)
				}
				if k > ss.Star[code].HZCO {
					freeOuter[code] = append(freeOuter[code], k)
				}

				freeOrbits[code] = append(freeOrbits[code], k)
			}
		}
	}
	freeCodes := []string{}
	for i, cd := range freeOrbits {
		if len(cd) > 0 {
			freeCodes = append(freeCodes, i)
		}
	}
	if len(freeCodes) == 0 {
		return fmt.Errorf("must be free orbits to asign body to")
	}
	// starCode = freeCodes[dice.Sroll(fmt.Sprintf("1d%v-1", len(freeCodes)))]

	if ss.WorldType["INNER"] > 0 && len(freeInner) > 0 {
		l := 0
		for l < 1 {
			starCode = freeCodes[dice.Sroll(fmt.Sprintf("1d%v-1", len(freeCodes)))]
			l = len(freeInner[starCode])
			if l < 1 {
				continue

			}
			//r := dice.Sroll(fmt.Sprintf("1d%v-1", l))
			//fmt.Println(r, freeInner[starCode], freeInner[starCode][r], "выбрали")
			//orbt := freeInner[starCode][dice.Sroll(fmt.Sprintf("1d%v-1", len(freeInner[starCode])))]
			orbt := freeInner[starCode][dice.Sroll(fmt.Sprintf("1d%v-1", l))]
			fmt.Println(orbt, "создали")
			ss.Star[starCode].ChildOrbit[fmt.Sprintf("%v", orbt)].AsignedBody = body
			fmt.Println(body, "добавили")
			ss.WorldType["INNER"]--
			fmt.Println(ss.WorldType["INNER"], "уменьшили")
		}

		return nil

	}
	if ss.WorldType["OUTER"] > 0 && len(freeOuter) > 0 {
		orbt := freeOuter[starCode][dice.Sroll(fmt.Sprintf("1d%v-1", len(freeOuter[starCode])))]
		ss.Star[starCode].ChildOrbit[fmt.Sprintf("%v", orbt)].AsignedBody = body
		ss.WorldType["OUTER"]--
		return nil

	}
	if len(freeOrbits) > 0 {
		orbt := freeOrbits[starCode][dice.Sroll(fmt.Sprintf("1d%v-1", len(freeOrbits[starCode])))]
		ss.Star[starCode].ChildOrbit[fmt.Sprintf("%v", orbt)].AsignedBody = body
		return nil

	}
	panic("unexpected out of orbits")
	return nil

}

func countFree(free map[string][]float64) int {
	r := 0
	for _, code := range free {
		r += len(code)
	}
	return r
}

func (ss *StarSystem) placeOrbitsN(dice *dice.Dicepool, spread []float64, bloSl []float64, allowed []int, sbnType string) {
	switch sbnType {
	case "a":
		ss.WorldType["INNER"] = ss.WorldType["TOTAL"] - ss.WorldType["sbn"] - 1
		ss.WorldType["OUTER"] = ss.WorldType["TOTAL"] - ss.WorldType["INNER"] - 1
	case "b":
		ss.WorldType["OUTER"] = ss.WorldType["TOTAL"] - 1
	case "c":
		ss.WorldType["INNER"] = ss.WorldType["TOTAL"] - 1
	}
	for i, st := range []string{"Aa", "Ba", "Ca", "Da"} {
		if _, ok := ss.Star[st]; !ok {
			continue
		}
		starPair := ss.Star[st]
		filled := false
		firstOrbit := normalizeOrbit(ss.Star[st].MAO + (spread[i] * variation(dice, 1)))
		if firstOrbit == 0 {
			firstOrbit += 0.1
		}
		curntOrbit := firstOrbit
		nextOrbit := firstOrbit
		starPair.AddApproved(firstOrbit)
		try := 0
		for !filled {
			try++
			if curntOrbit > 20 {
				filled = true
				break
			}
			if try > 1000 {
				fmt.Println("LOOP")
				break
			}

			nextOrbit = normalizeOrbit(curntOrbit + (spread[i] * variation(dice, try)))

			if nextOrbit <= curntOrbit {
				nextOrbit = curntOrbit + 0.2
			}
			nextOrbit = normalizeOrbit(nextOrbit)
			curntOrbit = nextOrbit
			starPair.AddApproved(nextOrbit)
		}
		ss.Star[st] = starPair
	}

}

func generateAnomalyes(dice *dice.Dicepool) map[int]string {
	r := dice.Sroll("2d6")
	an := r - 9
	anMap := make(map[int]string)
	for i := an; i > 0; i-- {
		switch dice.Sroll("2d6") {
		default:
			anMap[i] = "random"
		case 8:
			anMap[i] = "eccentric"
		case 9:
			anMap[i] = "inclined"
		case 10, 11:
			anMap[i] = "retrograde"
		case 12:
			anMap[i] = "trojan"
		}

	}
	return anMap
}

func variation(dice *dice.Dicepool, try int) float64 {
	vr := 1.0
	for i := 0; i < try; i++ {

		r := dice.Sroll("2d6+3")
		vr = vr * (float64(100+r) / 100)
	}
	return vr
}

func normalizeOrbit(orb float64) float64 {
	or := float64(int(orb*100)) / 100
	if or < 0.1 {
		or = 0.1
	}
	return or
}

/*
easy spread:
10-20%
2d6 + 8
*/

func determineSystemSpread(bloSl []float64, maoSL []float64, bln int, allowedOrbits []int, totalStars int, availableOrbitsSum []float64) []float64 {
	spreads := []float64{}
	if bln < 1 {
		bln = 1
	}
	for i, allocated := range allowedOrbits {
		if allocated == 0 {
			spreads = append(spreads, 0)
			continue
		}
		spr := (bloSl[i] - maoSL[i]) / float64(bln)
		maxSpr := 0.0
		if spr < 0 {
			spr = spr * -1
		}
		fmt.Println(bloSl[i], maoSL[i], float64(bln))
		switch i {
		case 0:
			maxSpr = float64(availableOrbitsSum[i]) / (float64(allowedOrbits[i]) + float64(totalStars))
		case 1, 2, 3:
			maxSpr = availableOrbitsSum[i] / float64(allowedOrbits[i]+1)
		}

		fmt.Println("--", spr, maxSpr)
		if maxSpr <= spr {
			spr = maxSpr
			if allowedOrbits[i] == 0 {
				spr = 0
			}
		}
		spr = float64(int(spr*1000)) / 1000
		spreads = append(spreads, spr)
	}
	return spreads
}

func systemBaselineNumType(sbn, totalWorlds, sysGenMeth int) string {
	if sysGenMeth == GenerationMethodContinuation {
		return "d"
	}
	if sbn >= 1 && sbn <= totalWorlds {
		return "a"
	}
	if sbn < 1 {
		return "b"
	}
	return "c"
}

func baselineOrbit_A(dice *dice.Dicepool, hzcoSl []float64) []float64 {
	bloSl := []float64{}
	for _, hzco := range hzcoSl {
		switch hzco >= 1 {
		case true:
			bloSl = append(bloSl, hzco+(float64(dice.Sroll("2d6")-7)*0.1))
		default:
			bloSl = append(bloSl, hzco+(float64(dice.Sroll("2d6")-7)*0.01))
		}
	}
	return bloSl
}

func baselineOrbit_B(dice *dice.Dicepool, hzcoSl []float64, blNum int, totalWrlds int) []float64 {
	bloSl := []float64{}
	for _, hzco := range hzcoSl {
		switch hzco >= 1 {
		case true:
			bloSl = append(bloSl, hzco-float64(blNum)+float64(totalWrlds)+(float64(dice.Sroll("2d6")-2)*0.1))
		default:
			bloSl = append(bloSl, hzco-(float64(blNum)*0.1)+(float64(dice.Sroll("2d6")-2)*0.01))
		}
	}
	return bloSl
}

func baselineOrbit_C(dice *dice.Dicepool, hzcoSl []float64, blNum int, totalWrlds int) []float64 {
	bloSl := []float64{}
	for _, hzco := range hzcoSl {
		switch hzco >= 1 {
		case true:
			bloSl = append(bloSl, hzco-float64(blNum)+float64(totalWrlds)+(float64(dice.Sroll("2d6")-7)*0.2))
		default:
			bloSl = append(bloSl, hzco-((float64(blNum+totalWrlds)+float64(dice.Sroll("2d6-7"))*0.2)*0.1))
		}
	}
	return bloSl
}

func systemBaseNumber(syst *StarSystem, dice *dice.Dicepool, totalWorlds int) int {
	r := dice.Sroll("2d6")
	if _, ok := syst.Star["Ab"]; !ok {
		r = r - 2
	}
	primary := syst.Star["Aa"]
	switch primary.Class {
	case star.ClassIa, star.ClassIb, star.ClassII:
		r = r + 3
	case star.ClassIII:
		r = r + 2
	case star.ClassIV:
		r = r + 1
	case star.ClassVI:
		r = r - 1
	case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole:
		r = r - 2
	}
	switch totalWorlds {
	default:
		if totalWorlds < 6 {
			r = r - 4
		}
		if totalWorlds > 20 {
			r = r + 2
		}
	case 6, 7, 8, 9:
		r = r - 3
	case 10, 11, 12:
		r = r - 2
	case 13, 14, 15:
		r = r - 1
	case 18, 19, 20:
		r = r + 1
	}
	return r
}

func emptyOrbits(dice *dice.Dicepool) int {
	r := dice.Sroll("2d6")
	eo := r - 9
	if eo <= 0 {
		return 0
	}
	return eo
}
