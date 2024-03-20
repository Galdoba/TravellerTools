package career

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/Galdoba/TravellerTools/hostile/character/characteristic"
	"github.com/Galdoba/TravellerTools/hostile/character/check"
	"github.com/Galdoba/TravellerTools/pkg/decidion"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	colonist          = "colonist.json"
	commercial_spacer = "commercial_spacer.json"
	corporate_agent   = "corporate_agent.json"
	corporate_exec    = "corporate_exec.json"
	marine            = "marine.json"
	marshal           = "marshal.json"
	military_spacer   = "military_spacer.json"
	physician         = "physician.json"
	ranger            = "ranger.json"
	rogue             = "rogue.json"
	roughneck         = "roughneck.json"
	scientist         = "scientist.json"
	survey_scout      = "survey_scout.json"
	technician        = "technician.json"
)

type CareerState interface {
	Qualify(*dice.Dicepool, *characteristic.CharSet) bool
	Survived(*dice.Dicepool, *characteristic.CharSet) bool
	CommisionReceived(*dice.Dicepool, *characteristic.CharSet) bool
	AdvancementReceived(*dice.Dicepool, *characteristic.CharSet, bool) bool
	Name() string
	Train(*dice.Dicepool, bool) string
	CanAdvance(bool) bool
	ReEnlisted(*dice.Dicepool, bool) bool
	MusterOut(*dice.Dicepool, bool, bool) []string
	RankBonus() string
}

type careerState struct {
	careerStats     *CareerStats
	commisionPassed bool
	totalTerms      int
	activeRank      int
	nextTermEnlist  int
}

func StartCareer(careerName string, dice *dice.Dicepool, charSet *characteristic.CharSet, byDraft bool) (*careerState, error) {
	cr := careerState{}
	cs, err := LoadCareerStats(careerName)
	if err != nil {
		return nil, fmt.Errorf("can't start career: %v", err.Error())
	}
	cr.careerStats = cs
	cr.activeRank = cr.careerStats.Ranks[0].Value
	if byDraft {
		return &cr, nil
	}
	if cr.Qualify(dice, charSet) {
		return &cr, nil
	}
	if _, _, err := check.ParseCode(cs.Qualification); err != nil {
		_, _, err = check.ParseCode(cs.QualificationReq)
		if err != nil {
			return nil, fmt.Errorf("can't start career: %v", err.Error())
		}
	}
	// if !cr.Qualify(dice, charSet) {
	// 	return nil, fmt.Errorf("can't start career: failed to qualify")
	// }

	return &cr, nil
}

func (cs *careerState) Qualify(dice *dice.Dicepool, charSet *characteristic.CharSet) bool {
	code := cs.careerStats.Qualification
	if code == "" && cs.careerStats.QualificationReq != "" {
		chrCode, tn, err := check.ParseCode(cs.careerStats.QualificationReq)
		if err != nil {
			panic(err.Error())
		}
		if charSet.Chars[chrCode].Maximum.Value() >= tn {
			return true
		}
		return false
	}
	chrCode, tn, _ := check.ParseCode(code)
	dm := charSet.Chars[chrCode].Mod()
	r := dice.Sroll("2d6") + dm
	if r >= tn {
		return true
	}
	return false
}

func (cs *careerState) Train(dice *dice.Dicepool, pc bool) string {
	keys := keysFrom(cs.careerStats.SkillTable)
	key := ""
	switch pc {
	case false:
		key = decidion.Random_One(dice, keys...)
	case true:
		panic(1)
	}
	// fmt.Println("table", key, "selected")
	bonus := decidion.Random_One(dice, cs.careerStats.SkillTable[key]...)
	// fmt.Println(bonus, "== received")
	return bonus
}

func keysFrom(smap map[string][]string) []string {
	keys := []string{}
	for k := range smap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (cs *careerState) Survived(dice *dice.Dicepool, charSet *characteristic.CharSet) bool {
	chrCode, tn, err := check.ParseCode(cs.careerStats.Survival)
	if err != nil {
		panic(err.Error())
	}
	dm := 0
	switch chrCode {
	case characteristic.BASIC:
	default:
		dm = charSet.Chars[chrCode].Mod()
	}
	r := dice.Sroll("2d6") + dm
	if r >= tn {
		cs.totalTerms++
		return true
	}
	// check.Char(cs.careerStats.Survival, dice, charSet)

	return false
}

func (cs *careerState) CommisionReceived(dice *dice.Dicepool, charSet *characteristic.CharSet) bool {
	if cs.commisionPassed == true {
		return false
	}
	chrCode, tn, err := check.ParseCode(cs.careerStats.Commision)
	if err != nil {
		panic(err.Error())
	}
	dm := charSet.Chars[chrCode].Mod() + charSet.Chars[characteristic.SOC].Mod()
	r := dice.Sroll("2d6") + dm
	if r >= tn {
		cs.commisionPassed = true
		cs.activeRank = 1
		return true
	}
	return false
}

func (cs *careerState) AdvancementReceived(dice *dice.Dicepool, charSet *characteristic.CharSet, nco bool) bool {
	switch nco {
	case true:
		if cs.CanAdvance(true) {
			cs.activeRank++
			return true
		}
	case false:
		if cs.commisionPassed == false {
			return false
		}
		if !cs.CanAdvance(false) {
			return false
		}
		chrCode, tn, err := check.ParseCode(cs.careerStats.Advance)
		if err != nil {
			panic(err.Error())
		}
		dm := charSet.Chars[chrCode].Mod() + charSet.Chars[characteristic.SOC].Mod()
		r := dice.Sroll("2d6") + dm
		if r >= tn {
			cs.activeRank++
			return true
		}
	}
	return false
}

func LoadCareerStats(name string) (*CareerStats, error) {
	cs := &CareerStats{}
	sep := string(filepath.Separator)

	usrHome, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("userHomeread: %v", err)
	}
	name = name2file(name)
	bt, err := os.ReadFile(usrHome + sep + `TabletopGames` + sep + `HOSTILE` + sep + `careers` + sep + name)
	if err != nil {
		return nil, fmt.Errorf("readFile: %v", err)
	}
	err = json.Unmarshal(bt, cs)

	if err != nil {
		return nil, fmt.Errorf("Unmarshal: %v", err)
	}
	return cs, nil
}

func LoadAllDEBUG() []*CareerStats {
	list := []*CareerStats{}
	for _, name := range []string{
		Colonist,
		CorporateAgent,
		CorporateExec,
		CommersialSpacer,
		Marine,
		Marshal,
		MilitarySpacer,
		Physician,
		Ranger,
		Rogue,
		Roughneck,
		Scientist,
		SurveyScout,
		Technician,
	} {
		cs, _ := LoadCareerStats(name)
		list = append(list, cs)
	}
	return list
}

func name2file(name string) string {
	nmap := make(map[string]string)
	nmap[Colonist] = colonist
	nmap[CommersialSpacer] = commercial_spacer
	nmap[CorporateAgent] = corporate_agent
	nmap[CorporateExec] = corporate_exec
	nmap[Marine] = marine
	nmap[Marshal] = marshal
	nmap[MilitarySpacer] = military_spacer
	nmap[Physician] = physician
	nmap[Ranger] = ranger
	nmap[Rogue] = rogue
	nmap[Roughneck] = roughneck
	nmap[Scientist] = scientist
	nmap[SurveyScout] = survey_scout
	nmap[Technician] = technician

	return nmap[name]
}

func (cs *careerState) Name() string {
	return cs.careerStats.Name
}

func (cs *careerState) CanAdvance(nco bool) bool {
	rnk, err := cs.careerStats.RankCurrent(cs.activeRank, cs.commisionPassed)
	if err != nil {
		fmt.Println(cs.activeRank, cs.commisionPassed, cs.careerStats.Name)
		panic("GET RANK " + err.Error())
	}
	switch nco {
	case true:
		if cs.commisionPassed == true {
			return false
		}
		for _, rank := range cs.careerStats.Ranks {
			if rank.Value == (rnk.Value+1) && rnk.CommisionRequired == !nco {
				return true
			}
		}
	case false:
		for _, rank := range cs.careerStats.Ranks {
			if rank.Value == (rnk.Value+1) && rnk.CommisionRequired == true {
				return true
			}
		}
	}
	return false
}

func (cs *careerState) ReEnlisted(dice *dice.Dicepool, manual bool) bool {
	tn := cs.careerStats.ReEnlist
	options := []string{}
	r := dice.Sroll("2d6")
	switch r {
	case 12:
		return true
	default:
		if r >= tn {
			switch manual {
			case false:
				options = append(options, "continue")
				options = append(options, "continue")
			case true:
				options = append(options, "continue")
			}
		}
		options = append(options, "muster out")
	}
	switch manual {
	case true:
		panic("manual not implemented")
	case false:
		switch decidion.Random_One(dice, options...) {
		case "continue":
			return true
		}
	}
	return false
}

func (cs *careerState) MusterOut(dice *dice.Dicepool, gambler bool, manual bool) []string {
	benefits := []string{}
	money := 3
	bdm := 0
	mdm := 0
	rolls := cs.totalTerms + cs.activeRank
	if cs.activeRank > 3 {
		rolls += (cs.activeRank - 3)
	}
	if cs.activeRank > 4 {
		bdm = 1
	}
	if cs.careerStats.Name == CorporateExec || gambler {
		mdm = 1
	}
	fmt.Println("Rolls", rolls, cs.totalTerms, cs.activeRank)
	for i := 1; i <= rolls; i++ {
		label := fmt.Sprintf("Mustering Out roll %v (%v left)", i, rolls-i)
		benefit := ""
		rollType := ""
		options := []string{"Benefit"}
		if money > 0 {
			options = append(options, fmt.Sprintf("Money (%v left)", money))
		}
		fmt.Println("r", i, label, options)

		switch manual {
		case true:
			rollType = decidion.Manual_One(label, false, options...)
		case false:
			rollType = decidion.Random_One(dice, options...)
		}
		fmt.Println(rollType)
		switch rollType {
		case "Benefit":
			options = cs.careerStats.MusterOut
			r := dice.Sroll("1d6") + bdm - 1
			benefit = options[r]
		default:
			money--
			options = []string{"$500", "$1000", "$1000", "$5000", "$8000", "$10000", "$20000"}
			r := dice.Sroll("1d6") + mdm - 1
			benefit = options[r]
		}

		fmt.Println(benefit)
		benefits = append(benefits, benefit)
	}
	return benefits
}

func money2int(s string) int {
	switch s {
	case "$500":
		return 500
	case "$1000":
		return 1000
	case "$5000":
		return 5000
	case "$8000":
		return 8000
	case "$10000":
		return 10000
	case "$20000":
		return 20000
	}
	return 0
}

func (cs *careerState) RankBonus() string {
	rnk, err := cs.careerStats.RankCurrent(cs.activeRank, cs.commisionPassed)
	if err != nil {
		panic("rank not found")
	}
	return rnk.AutoSkill
}
