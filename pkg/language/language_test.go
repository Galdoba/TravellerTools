package language

import "testing"

func TestNewLanguage(t *testing.T) {
	l, err := New("Zhodani")
	if l == nil {
		t.Errorf("func returned no object")
	}
	if err != nil {
		t.Errorf("func returned err = '%v'", err.Error())
	}
	if l.Name == "" {
		t.Errorf("language Name was not set")
	}
	if len(l.ConsonantsInitial) == 0 {
		t.Errorf("language ConsonantsInitial Sounds not set")
	}
	if len(l.Vowels) == 0 {
		t.Errorf("language Vowels Sounds not set")
	}
	if len(l.ConsonantsFinal) == 0 {
		t.Errorf("language ConsonantsFinal Sounds not set")
	}
	if l.BasicGenerationTable == nil {
		t.Errorf("language BasicGenerationTable not set")
	}
	if l.AlternativeGenerationTable == nil {
		t.Errorf("language AlternativeGenerationTable not set")
	}
	if l.InitialConsonantTable == nil {
		t.Errorf("language InitialConsonantTable not set")
	}
	if l.VowelTable == nil {
		t.Errorf("language VowelTable not set")
	}
	if l.FinalConsonantTable == nil {
		t.Errorf("language FinalConsonantTable not set")
	}
	if len(l.BasicGenerationTable) != 36 {
		t.Errorf("BasicGenerationTable expected to have 36 entries (have %v)", len(l.BasicGenerationTable))
	}
	if len(l.AlternativeGenerationTable) != 36 {
		t.Errorf("AlternativeGenerationTable expected to have 36 entries (have %v)", len(l.AlternativeGenerationTable))
	}

	for _, v := range valid2() {
		if l.BasicGenerationTable[v] == "" {
			t.Errorf("InitialConsonantTable[%v] is not set", v)
		}
		if l.AlternativeGenerationTable[v] == "" {
			t.Errorf("AlternativeGenerationTable[%v] is not set", v)
		}
	}

	if len(l.InitialConsonantTable) != 216 {
		t.Errorf("InitialConsonantTable expected to have 216 entries (have %v)", len(l.InitialConsonantTable))
	}
	if len(l.VowelTable) != 216 {
		t.Errorf("VowelTable expected to have 216 entries (have %v)", len(l.VowelTable))
	}
	if len(l.FinalConsonantTable) != 216 {
		t.Errorf("FinalConsonantTable expected to have 216 entries (have %v)", len(l.FinalConsonantTable))
	}

	for _, v := range valid3() {
		if l.InitialConsonantTable[v] == "" {
			t.Errorf("InitialConsonantTable[%v] is not set", v)
		}
		if l.VowelTable[v] == "" {
			t.Errorf("VowelTable[%v] is not set", v)
		}
		if l.FinalConsonantTable[v] == "" {
			t.Errorf("FinalConsonantTable[%v] is not set", v)
		}
	}

}
