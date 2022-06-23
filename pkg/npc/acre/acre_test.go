package acre

import (
	"fmt"
	"strings"
	"testing"
)

type testInput struct {
	name      string
	relatedPC string
	status    int
}

func inputList() []testInput {
	inp := []testInput{}
	for _, name := range []string{"N1", "N2", "N3", "N4", "N5", "N6", "N7", "N8", "N9", "N0", "N11", "N12"} {
		for _, pc := range []string{"PC1", "PC2", "PC3", "PC4", "PC5", "PC6", "PC7", "PC8", "PC9", "PC0", "PC11", "PC12"} {
			for _, stat := range []int{StatusDefault, Ally, Contact, Rival, Enemy, 42} {
				inp = append(inp, testInput{name, pc, stat})
			}

		}
	}
	inp = append(inp, testInput{"Кайра", "Марк", Contact})
	return inp
}

func errorValid(err error) bool {
	switch {
	default:
		return false
	case strings.Contains(err.Error(), "person was not named"):
		return true
	case strings.Contains(err.Error(), "person was not related to PC"):
		return true
	case strings.Contains(err.Error(), "status code unknown"):
		return true
	}
}

func TestContact(t *testing.T) {
	for tn, input := range inputList() {
		t.Log("test", tn, ":", input)

		npc, err := New(input.name, input.relatedPC, input.status)
		if npc == nil {
			t.Errorf("func returned no object")
			continue
		}
		npc.SetRole("[Role]")
		if err != nil {
			if !errorValid(err) {
				t.Errorf("func returned error: %v", err.Error())
				break
			} else {
				t.Logf("	expected error: %v\n", err.Error())
			}
			continue
		}
		if npc.name == "" {
			t.Errorf("npc.name not filled")
		}
		if npc.relatedPC == "" {
			t.Errorf("npc.relatedPC not filled")
		}
		if npc.name == "" {
			t.Errorf("npc.name not filled")
		}
		switch npc.status {
		case Ally, Contact, Rival, Enemy:
		default:
			t.Errorf("status value unknown (%v)", npc.status)
		}
		if npc.affinity < 0 {
			t.Errorf("affinity check failed")
		}
		if npc.enmity < 0 {
			t.Errorf("enmity check failed")
		}
		if npc.influence < 0 {
			t.Errorf("influence check failed")
		}
		if npc.power < 0 {
			t.Errorf("power check failed")
		}
		if npc.role == "" {
			t.Errorf("role not assighned")
		}
		t.Log(npc)
		fmt.Println(npc)
		fmt.Println(" ")
	}
}
