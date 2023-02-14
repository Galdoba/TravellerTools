package task

func Phrase(s string) instruction {
	return instruction{
		key:  instKey_TaskPhrase,
		valS: "To " + s,
	}
}

func Duration(durType int, units int64) instruction {
	in := instruction{}
	in.key = "Duration"
	in.valI = durType
	switch durType {
	default:
		in.key = ignoreInstruction
	case Ignored:
		in.valInt64 = -1
	case Absolute, Variable, Randomized:
		in.valInt64 = units
	}
	return in
}

func Difficulty(d int) instruction {
	return instruction{
		key:  instKey_BaseDifficulty,
		valI: d,
	}
}

func Char(characteristic string) instruction {
	return instruction{
		key:  instKey_Chararteristic,
		valS: characteristic,
	}
}

const (
	USAGE_NoSkill      = "No Skill"
	USAGE_SkillOnly    = "Skill Only"
	USAGE_SkillOptinal = "Optional"
)

func Skill(skill, skillUsage string) instruction {
	inst := instruction{key: instKey_Skill}
	switch skillUsage {
	default:
		inst.valS = skill
	case USAGE_NoSkill:
		inst.valS = "3 (" + skillUsage + ")"
	case USAGE_SkillOnly, USAGE_SkillOptinal:
		inst.valS = skill + " (" + skillUsage + ")"
	}
	return inst
}

func Modifier(mod string, val int, req bool) instruction {
	return instruction{
		key:  instKey_Mod,
		valS: mod,
		valI: val,
		valB: req,
	}
}
