package star

import (
	"fmt"

	"github.com/Galdoba/devtools/errmaker"
)

type starpair struct {
	primary       string
	companion     string
	temperature   int
	mass          float64
	luminocity    float64
	innerLimit    float64
	habitableLow  float64
	habitableHigh float64
	snowLine      float64
	outerLimit    float64
	habZone       int
	pairDistance  int
}

// G2 V [M6 VI]
// G2 V (with companion M6 VI)

func NewPair(p string, c string) (*starpair, error) {
	sp := starpair{}
	primary := New(p)
	companion := New(c)
	if primary.Code() == "" {
		return nil, errmaker.ErrorFrom(fmt.Errorf("bad input"), p, c)
	}
	sp.primary = primary.Code()
	sp.companion = companion.Code()
	sp.temperature = primary.temperature + companion.temperature
	sp.mass = primary.mass + companion.mass
	sp.luminocity = primary.luminocity + companion.luminocity
	sp.innerLimit = primary.innerLimit + companion.innerLimit
	sp.habitableLow = primary.habitableLow + companion.habitableLow
	sp.habitableHigh = primary.habitableHigh + companion.habitableHigh
	sp.snowLine = primary.snowLine + companion.snowLine
	sp.outerLimit = primary.outerLimit + companion.outerLimit
	sp.habZone = primary.habitableOrbit
	return &sp, nil
}

type StarBody interface {
	Class() string
	Mass() float64
	Luminocity() float64
	InnerLimit() float64
	HabitableLow() float64
	HabitableHigh() float64
	OuterLimit() float64
	HabitableZone() int
}

func (p *starpair) Class() string {
	class := p.primary
	if p.companion != "" {
		class += " (with companion " + p.companion + ")"
	}
	return class
}
func (p *starpair) Mass() float64 {
	return p.mass
}
func (p *starpair) Luminocity() float64 {
	return p.luminocity
}
func (p *starpair) InnerLimit() float64 {
	return p.innerLimit
}
func (p *starpair) HabitableLow() float64 {
	return p.habitableLow
}
func (p *starpair) HabitableHigh() float64 {
	return p.habitableHigh
}
func (p *starpair) OuterLimit() float64 {
	return p.outerLimit
}

func (p *starpair) HabitableZone() int {
	return p.habZone
}
