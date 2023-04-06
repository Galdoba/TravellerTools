package pawn

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic/charset"
)

type pawn struct {
	chrSet *charset.CharSet
}

/*
A Characteristics
B Birthworld
C Education
D Careers
E Muster Out


*/

type GeneTemplate interface { //фактически это 2 стринга. база находится в генетике
	Profile() string
	Variations() string
}

//func New(geneTemplate genetics.Base, homeworld ) (*pawn, error) {

func New(gt GeneTemplate) (*pawn, error) {
	chr := pawn{}
	if err := chr.Characteristics(gt); err != nil {
		return &chr, err
	}
	return &chr, nil
}

func (chr *pawn) Characteristics(genetics GeneTemplate) error {
	// geneTemplate, err := genetics.GeneTemplateManual(genetics, sequance)
	// if err != nil {
	// 	return err
	// }
	chr.chrSet = charset.NewCharSet(dice.New(), genetics.Profile(), genetics.Variations())
	return nil
}
