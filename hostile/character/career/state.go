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
	Report() string
	Train(*dice.Dicepool, bool) string
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
	fmt.Println("table", key, "selected")
	bonus := decidion.Random_One(dice, cs.careerStats.SkillTable[key]...)
	fmt.Println(bonus, "== received")
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
	dm := charSet.Chars[chrCode].Mod()
	r := dice.Sroll("2d6") + dm
	if r >= tn {
		return true
	}
	check.Char(cs.careerStats.Survival, dice, charSet)

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
		return true
	}
	return false
}

func (cs *careerState) AdvancementReceived(dice *dice.Dicepool, charSet *characteristic.CharSet) bool {

	switch cs.commisionPassed {
	case true:
	case false:
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

func (cs *careerState) Report() string {
	str := cs.careerStats.Name
	// str := fmt.Sprintf("%v", cs.careerStats)
	// str += fmt.Sprintf("%v", cs)
	return str
}

func (cs *careerState) CanAdvance(nco bool) bool {
	rnk := cs.
}
