package essc

import "testing"

func inputSeed() []string {
	return []string{
		"test_1",
	}
}

func TestStarSystem(t *testing.T) {
	for _, seed := range inputSeed() {
		ss, err := New("test")
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
		}
		if ss == nil {
			t.Errorf("star system was not created: seed=|%v|", seed)
		}
	}
}
