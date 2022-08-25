package language

const (
	InitialConsonant = iota
	Vowel
	FinalConsonant
)

type Sound struct {
	sType          int
	sound          string
	freq           int
	pronuansiation string
}

type Language struct {
	Name                       string
	ConsonantsInitial          []Sound
	Vowels                     []Sound
	ConsonantsFinal            []Sound
	BasicGenerationTable       map[string]string
	AlternativeGenerationTable map[string]string
	InitialConsonantTable      map[string]string
	VowelTable                 map[string]string
	FinalConsonantTable        map[string]string
}

func New(name string) (*Language, error) {
	l := Language{}
	l.Name = name
	l.ConsonantsInitial, l.Vowels, l.ConsonantsFinal,
		l.BasicGenerationTable, l.AlternativeGenerationTable,
		l.InitialConsonantTable, l.VowelTable, l.FinalConsonantTable = callTables(name)
	return &l, nil
}
