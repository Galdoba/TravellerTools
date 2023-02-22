package skillset

import "github.com/Galdoba/TravellerTools/t5/pawn/skill"

type skillSet struct {
	ByName       map[string]*skill.Skill
	incrementLog []int
}

func New() *skillSet {
	ss := skillSet{}
	ss.ByName = make(map[string]*skill.Skill)
	for i := skill.ID_NONE; i < skill.ID_END; i++ {
		skl := skill.New(i)
		if !skl.Default {
			continue
		}
		name := skill.NameByID(i)
		switch name {
		default:
		case "Language", "Instrument":
			continue //пока не создаем язык ибо особая механика
		}
		skl.Name = name
		ss.ByName[name] = skl
	}
	return &ss
}
