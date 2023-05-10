package skill

import (
	"fmt"
	"testing"
)

func TestSkill(t *testing.T) {
	for i := ID_NONE; i < ID_END; i++ {
		skl, err := New(i)
		fmt.Println(i, NameByID(i), skl)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}

func TestSkillNameLenLongest(t *testing.T) {
	lMax := LongestNameLen()
	if lMax != LongestNameLength {
		t.Errorf("lMax != LongestNameLength (%v != %v)\nconst 'LongestNameLength' must be updated", lMax, LongestNameLength)
	}
}
