package worldgen

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type generator struct {
	dice    *dice.Dicepool
	options map[string]string
}

func NewGenerator(seed string, options ...worldGenOption) *generator {
	g := generator{}
	g.dice = dice.New().SetSeed(seed)
	g.options = make(map[string]string)
	for _, opt := range options {
		g.options[opt.key] = opt.val
	}
	return &g
}

type worldGenOption struct {
	key, val string
}

func Option(key, val string) worldGenOption {
	return worldGenOption{key, val}
}

type WorldGen interface {
	Reset(string, ...worldGenOption)
}

func (g *generator) Reset(seed string, options ...worldGenOption) {
	g = NewGenerator(seed, options...)
}
