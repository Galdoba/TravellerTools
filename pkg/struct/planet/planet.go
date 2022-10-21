package planet

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/struct/star"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

type planetOrbit struct {
	index       int
	orbitSuffix string
	habitZone   int
	twilight    bool
	comment     string
}

type Planet struct {
	orbit   planetOrbit
	uwp     uwp.UWP
	remarks []string
}

func New() (*Planet, error) {
	return &Planet{}, nil
}

func NewOrbit(orbit int, parentStarStellar string) (planetOrbit, error) {
	po := planetOrbit{}
	po.orbitSuffix = suffixMap(orbit)
	if po.orbitSuffix == "" {
		return po, fmt.Errorf("orbit index '%v' invalid", orbit)
	}
	switch orbit {
	case 0, 1:
		po.twilight = true
	}
	spectral, _, size, err := star.DecodeStellar(parentStarStellar)
	if err != nil {
		return po, fmt.Errorf("parent star decoding: %v", err.Error())
	}
	hz := star.HabitableZone(spectral, size)
	po.habitZone = orbit - hz
	po.index = orbit
	return po, nil
}

func suffixMap(orbit int) string {
	sm := make(map[int]string)
	for i := 0; i < 20; i++ {
		sm[i] = fmt.Sprintf("%v", i)
	}
	return sm[orbit]
}

func (p *planetOrbit) Orbit() int {
	return p.index
}
