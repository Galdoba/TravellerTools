package language

type Sound struct {
	sType          int
	sound          string
	freq           int
	pronuansiation string
}

type Language struct {
	Name                       string
	ConsonantsInitial          []Sound
	Viwels                     []Sound
	ConsonantsFinal            []Sound
	BasicGenerationTable       map[string]string
	AlternativeGenerationTable map[string]string
}
