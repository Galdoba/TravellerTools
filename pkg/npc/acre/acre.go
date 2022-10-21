package acre

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	StatusDefault = iota
	Ally
	Contact
	Rival
	Enemy
)

type person struct {
	name      string
	relatedPC string
	role      string
	status    int
	affinity  int
	enmity    int
	power     int
	influence int
	special   []string
}

func New(name, relatedPC string, status int) (*person, error) {
	p := person{}
	if name == "" {
		return &p, fmt.Errorf("person was not named")
	}
	p.name = name
	p.relatedPC = relatedPC
	if relatedPC == "" {
		return &p, fmt.Errorf("person was not related to PC")
	}
	err := fmt.Errorf("status code unknown (%v)", status)
	for _, val := range []int{Ally, Contact, Rival, Enemy} {
		if val == status {
			p.status = status
			err = nil
			break
		}
	}
	if err != nil {
		return &p, err
	}
	p.affinity, p.enmity, p.influence, p.power = -1, -1, -1, -1
	dp := dice.New().SetSeed(p.name + p.relatedPC + p.name + p.relatedPC)
	p.affinity = determine("Affinity", p.status, dp)
	p.enmity = determine("Enmity", p.status, dp)
	p.power = determine("Power", p.status, dp)
	p.influence = determine("Influence", p.status, dp)
	if dp.Roll("2d6").Sum() >= 8 {
		p.special = append(p.special, "Creation Event")
	}
	return &p, nil
}

func (p *person) SetRole(role string) {
	p.role = role
}

func determine(valueType string, status int, dp *dice.Dicepool) int {
	val := -1
	diceCode, dm := relashionships(status, valueType)
	r1 := dp.Roll(diceCode).DM(dm).Sum()
	val = getValue(valueType, r1)
	return val
}

func relashionships(status int, value string) (string, int) {
	code := "1d1"
	dm := -2
	switch value {
	default:
		return code, dm
	case "Affinity":
		switch status {
		case Ally:
			code = "2d6"
			dm = 0
		case Contact:
			code = "1d6"
			dm = 1
		case Rival:
			code = "1d6"
			dm = -1
		case Enemy:
			code = "1d1"
			dm = -1
		}
	case "Enmity":
		switch status {
		case Ally:
			code = "1d1"
			dm = -1
		case Contact:
			code = "1d6"
			dm = -1
		case Rival:
			code = "1d6"
			dm = 1
		case Enemy:
			code = "2d6"
			dm = 0
		}
	case "Power", "Influence":
		code = "2d6"
		dm = 0
	}
	return code, dm
}

func getValue(valueType string, r1 int) int {
	val := -5
	switch valueType {
	default:
		fmt.Println("DEBUG: valueType =", valueType)
		return val
	case "Affinity", "Enmity":
		switch r1 {
		default:
			if r1 >= 12 {
				val = 6
			}
			if r1 <= 2 {
				val = 0
			}
		case 3, 4:
			val = 1
		case 5, 6:
			val = 2
		case 7, 8:
			val = 3
		case 9, 10:
			val = 4
		case 11:
			val = 5
		}
	case "Power", "Influence":
		switch r1 {
		default:
			if r1 <= 5 {
				val = 0
			}
		case 6, 7:
			val = 1
		case 8:
			val = 2
		case 9:
			val = 3
		case 10:
			val = 4
		case 11:
			val = 5
		case 12:
			val = 6
		}
	}
	return val
}

func (p *person) String() string {
	str := fmt.Sprintf("Name      : %v\n", p.name)
	str += fmt.Sprintf("Type      : %v\n", p.role)
	str += fmt.Sprintf("Status    : %v\n", describeStatus(p.status))
	str += fmt.Sprintf("Related to: %v\n", p.relatedPC)
	str += fmt.Sprintf("Affinity  : %v\n", p.affinity)
	str += fmt.Sprintf("Enmity    : %v\n", p.enmity)
	str += fmt.Sprintf("Power     : %v\n", p.power)
	str += fmt.Sprintf("Influence : %v\n", p.influence)
	str += fmt.Sprintf("%v\n", describeAffinity(p.affinity))
	str += fmt.Sprintf("%v\n", describeEnmity(p.enmity))
	str += fmt.Sprintf("%v\n", describePower(p.power))
	str += fmt.Sprintf("%v\n", describeInfluence(p.influence))
	return str
}

func describeStatus(i int) string {
	switch i {
	default:
		return fmt.Sprintf("[NO DESCR FOR STATUS CODE %v]", i)
	case Ally:
		return "ALLY"
	case Contact:
		return "CONTACT"
	case Rival:
		return "RIVAL"
	case Enemy:
		return "ENEMY"
	}
}

func describeAffinity(i int) string {
	switch i {
	default:
		return fmt.Sprintf("[NO DESCR FOR INDEX %v]", i)
	case 0:
		return fmt.Sprintf("No positive affinity towards the Traveller. This may be an enemy or just someone who does not care at all what happens to the Traveller, depending on Enmity")
	case 1:
		return fmt.Sprintf("Vaguely well inclined: About as well inclined towards the Traveller as any random stranger with a social conscience is likely to be. They will take minor actions to help, largely out of common courtesy, but not go to much trouble on the Traveller's behalf unless there are benefits to the action.")
	case 2:
		return fmt.Sprintf("Positively inclined: Will probably help in a safe and easy manner if asked, even without reward, but will not take much risk.")
	case 3:
		return fmt.Sprintf("Very positively inclined: Will take modest risks on the Traveller's behalf or offer help without being asked, if they realise their friend could benefit.")
	case 4:
		return fmt.Sprintf("Loyal friend: Will do almost anything to help the Traveller, but may have higher loyalties to their own family, cause or service, or to other close friends.")
	case 5:
		return fmt.Sprintf("Love: Will probably put the Traveller's interests before their own or that of others.")
	case 6:
		return fmt.Sprintf("Fanatical: Will do whatever the Traveller asks of them (or what they think the Traveller would want), no matter what risks are involved. May also expect others to do the same.")
	}
}

func describeEnmity(i int) string {
	switch i {
	default:
		return fmt.Sprintf("[NO DESCR FOR INDEX %v]", i)
	case 0:
		return fmt.Sprintf("No Enmity towards the Traveller. This may be because they do not know who the Traveller is, or because the Traveller has done nothing to offend them.")
	case 1:
		return fmt.Sprintf("Mistrustful: Vaguely ill-disposed towards the Traveller (or perhaps everyone in general) but will not go out of their way to impede them. Someone with an Enmity value of -1 is unlikely to take an action that will have serious consequences for the Traveller unless there is some great benefit.")
	case 2:
		return fmt.Sprintf("Negatively inclined: May engage in acts of petty spite for no gain, just to annoy and upset the Traveller. Someone with an Enmity value of -2 will probably stop short of actions that would seriously harm or kill the Traveller.")
	case 3:
		return fmt.Sprintf("Very negatively inclined: Will go to some trouble to impede the Traveller, just out of spite. Does not care much what happens to the Traveller and will more than likely feel they deserve anything they get.")
	case 4:
		return fmt.Sprintf("Hatred: Will do almost anything to get one over on the Traveller. Might actively plot against the Traveller for the sake of revenge or causing further harm even if there is little or no gain involved.")
	case 5:
		return fmt.Sprintf("Bitter hatred: Will actively plot or take serious risks to cause the Traveller harm at any opportunity")
	case 6:
		return fmt.Sprintf("Blinded by hate: May engage in self-destructive actions in order to harm the Traveller, or put innocents at risk.")
	}
}

func describePower(i int) string {
	switch i {
	default:
		return fmt.Sprintf("[NO DESCR FOR INDEX %v]", i)
	case 0:
		return fmt.Sprintf("Powerless: The individual has virtually no resources they can bring to bear other than their own personal possessions")
	case 1:
		return fmt.Sprintf("Weak: Has a few friends or contacts who might be willing to help; the equivalent of a typical band of Travellers.")
	case 2:
		return fmt.Sprintf("Useful: Has a significant asset such as a small starship and crew, or a small force of skilled mercenaries, high-end lawyers, or the like.")
	case 3:
		return fmt.Sprintf("Moderately Powerful: Has access to very significant assets such as a mercenary unit or a modest sized business entity.")
	case 4:
		return fmt.Sprintf("Powerful: Has powerful assets, equivalent to a small merchant shipping line or major business group.")
	case 5:
		return fmt.Sprintf("Very Powerful: Has enormous power, such as someone in the top echelons of a planetary government or the CEO of a large shipping line.")
	case 6:
		return fmt.Sprintf("Major Player: Is a factor in interstellar politics, such as a navy admiral or an official in an interstellar government.")
	}
}

func describeInfluence(i int) string {
	switch i {
	default:
		return fmt.Sprintf("[NO DESCR FOR INDEX %v]", i)
	case 0:
		return fmt.Sprintf("Has virtually no influence over anyone.")
	case 1:
		return fmt.Sprintf("Little Influence: Owed a couple of favours by minor officials and local notables such as the leader of a street gang or a port authority official.")
	case 2:
		return fmt.Sprintf("Some Influence: Has one or more minor local notables 'in their pocket' and can get them to act illegally or dangerously on the odd occasion.")
	case 3:
		return fmt.Sprintf("Influential: Has some influence over powerful people such as mid-level planetary government officials or rich portside merchant factors.")
	case 4:
		return fmt.Sprintf("Highly Influential: Has some influence at the interplanetary level, with government or underworld figures that owe him a favour or two.")
	case 5:
		return fmt.Sprintf("Extremely Influential: Has very significant influence at the interstellar level, and can lean on lawmakers or officials in interstellar government.")
	case 6:
		return fmt.Sprintf("Kingmaker: Has the ear of extremely powerful people, such as the ruling noble of the local subsector.")
	}
}
