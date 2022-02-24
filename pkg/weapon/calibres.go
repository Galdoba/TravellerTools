package weapon

type calibre struct {
	cType                 string
	baseDamageDie         int
	perDieMod             int
	costPer100            int
	baseCapacityVariation float64
	baseRange             float64
	unreliable            int
	slowLoader            int
	inacurate             int
	bulky                 int
	zeroG                 int
	physSignature         int
	emissionSignature     int
}
