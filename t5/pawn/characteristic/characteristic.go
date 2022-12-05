package characteristic

const (
	Human = "SDEIES"
)

type characteristic struct {
	positionCode   string
	abb            string
	name           string
	human          string
	description    string
	geneticProfile string
	charType       string
	dice           int
}

type profileOption struct {
	key  string
	valS string
	valI int
}

type personalityProfile struct {
	c1 characteristic
	c2 characteristic
	c3 characteristic
	c4 characteristic
	c5 characteristic
	c6 characteristic
	cs characteristic
	cp characteristic
}

func SetupProfile(template string) (*personalityProfile, error) {
	pp := personalityProfile{}

	return &pp, nil
}
