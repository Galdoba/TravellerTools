package skill

import (
	"fmt"
	"testing"
)

func TestSkillSet(t *testing.T) {
	sklst := NewSkillSet()
	sklst.AddBackGroundSkill(Administration)
	sklst.AddBackGroundSkill(Mechanical)
	sklst.Increase(Gunnery)
	sklst.Increase(Gunnery)
	sklst.Increase(Gunnery)
	sklst.Increase(Gunnery)
	fmt.Println(sklst)
	fmt.Println(sklst.SkillVal(Gunnery))
}
