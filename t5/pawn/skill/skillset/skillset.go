package skillset

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

type SkillSet struct {
	ID           map[int]*skill.Skill
	incrementLog []int
}

func NewSkillSet() (*SkillSet, error) {
	ss := SkillSet{}
	ss.ID = make(map[int]*skill.Skill)
	for _, id := range skillIDs() {
		if id == skill.ID_NONE {
			continue
		}
		skl, err := skill.New(id)
		if err != nil {
			return nil, fmt.Errorf("skillset.NewSkillSet(): can not create skillset: \n%v", err.Error())
		}
		if !skl.Default {
			continue
		}
		name := skill.NameByID(id)
		switch name {
		default:
		case "Language", "Instrument":
			continue //пока не создаем язык ибо особая механика
		}
		skl.Name = name
		ss.ID[id] = skl
	}
	return &ss, nil
}

func skillIDs() []int {
	id := []int{}
	for i := skill.ID_NONE; i < skill.ID_END; i++ {
		id = append(id, i)
	}
	return id
}

func skillNames() []string {
	names := []string{}
	for _, id := range skillIDs() {
		names = append(names, skill.NameByID(id))
	}
	return names
}

func (ssr *SkillSet) String() string {
	str := ""
	for _, id := range skillIDs() {
		if skl, ok := ssr.ID[id]; ok {
			str += fmt.Sprintf("%v %v\n", spacedStr(skl.Name, skill.LongestNameLength), spacedIntStr(skl.Value()))
		}
	}
	return str
}

func spacedStr(str string, n int) string {
	for len(str) < n {
		str += " "
	}
	return str
}

func spacedIntStr(i int) string {
	if i >= 0 && i <= 9 {
		return fmt.Sprintf(" %v", i)
	}
	return fmt.Sprintf("%v", i)
}

func have(sst *SkillSet, id int) bool {
	if _, ok := sst.ID[id]; ok {
		return true
	}
	return false
}

func (sst *SkillSet) AddSkill(id int) error {
	if have(sst, id) {
		return fmt.Errorf("skill already added")
	}
	skl, err := skill.New(id)
	if err != nil {
		return fmt.Errorf("skillset.AddSkill(id int): %v\n", err.Error())
	}
	sst.ID[id] = skl
	return nil
}

func KKSruleAllow(sst *SkillSet, sID int) bool {
	knlArr := []int{}
	sklVal := -1
	sklData, err := skill.New(sID)
	if err != nil {
		fmt.Println(1)
		return false
	}
	if sklData.SType() != skill.TYPE_SKILL {
		return true
	}
	switch have(sst, sID) {
	case false:
		if len(sklData.AssociatedKnowledge) == 0 {
			sst.AddSkill(sID)
			return true
		}
		fmt.Println(1)
		return false
	case true:
	}
	skl := sst.ID[sID]
	// if !skl.KKSrule {
	// 	return true
	// }
	if len(skl.AssociatedKnowledge) == 0 {
		return true
	}
	sklVal = skl.Value()
	for _, kID := range skl.AssociatedKnowledge {
		if have(sst, kID) {
			knl := sst.ID[kID].Value()
			knlArr = append(knlArr, knl)
		}
	}
	if sklVal < sum(knlArr)/2 {
		return true
	}
	return false
}

func sum(iArr []int) int {
	s := 0
	for _, v := range iArr {
		s += v
	}
	return s
}

var MustChooseErr = fmt.Errorf("must choose exact skill")

func (sst *SkillSet) Increase(id int) error {
	switch id {
	case skill.One_Art, skill.One_Trade:
		return MustChooseErr
	}
	if !KKSruleAllow(sst, id) {
		return fmt.Errorf("skillset.Increase(%v): KKS rule not allow", id)
	}
	if !have(sst, id) {
		sst.AddSkill(id)
	}
	if !have(sst, sst.ID[id].ParentSkl) && sst.ID[id].ParentSkl != skill.ID_NONE {
		sst.AddSkill(sst.ID[id].ParentSkl)
	}
	return sst.ID[id].Learn()
}

func (sst *SkillSet) IncreaseByKKSrule(id int) error {

	//skl := skill.New(id)
	//knlSet := skl.AssociatedKnowledge
	return nil
}

func (sst *SkillSet) CanAdd(id int) bool {
	return false
}
