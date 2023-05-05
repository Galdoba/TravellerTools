package pawn

import (
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/genetics"
)

func (p *pawn2) GenomeTemplate() (genetics.GeneProfile, error) {
	gp := profile.New()
	nodata := 0
	for _, key := range []string{
		KEY_GENE_PRF_1,
		KEY_GENE_PRF_2,
		KEY_GENE_PRF_3,
		KEY_GENE_PRF_4,
		KEY_GENE_PRF_5,
		KEY_GENE_PRF_6,
		KEY_GENE_MAP_1,
		KEY_GENE_MAP_2,
		KEY_GENE_MAP_3,
		KEY_GENE_MAP_4,
		KEY_GENE_MAP_5,
		KEY_GENE_MAP_6,
	} {
		data := p.profile.Data(key)
		if data == nil {
			nodata++
			continue
		}
		gp.Inject(key, data.Code())
	}
	if nodata == 12 {
		return genetics.NewGeneData(genetics.HumanGeneData()), nil
	}
	return gp, nil
}
