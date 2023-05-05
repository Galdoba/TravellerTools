package profile

// func UPPkeys() []string {
// 	return []string{
// 		KEY_VAL_C1,
// 		KEY_VAL_C2,
// 		KEY_VAL_C3,
// 		KEY_VAL_C4,
// 		KEY_VAL_C5,
// 		KEY_VAL_C6,
// 		KEY_VAL_CP,
// 		KEY_VAL_CS,
// 		KEY_GENE_PRF_2,
// 		KEY_GENE_PRF_3,
// 		KEY_GENE_PRF_5,
// 		KEY_GENE_PRF_6,
// 		KEY_GENE_MAP_2,
// 		KEY_GENE_MAP_3,
// 		KEY_GENE_MAP_5,
// 		KEY_GENE_MAP_6,
// 	}
// }

// func NewUPP(geneprof, genemap string) *universalProfile {
// 	prof := strings.Split(geneprof, "")
// 	if len(prof) != 6 {
// 		return nil
// 	}
// 	gmap := strings.Split(genemap, "")
// 	if len(gmap) != 6 {
// 		return nil
// 	}
// 	prf := universalProfile{}
// 	for i, v := range prof {
// 		switch i {
// 		case 0:
// 			if v != "S" {
// 				return nil
// 			}
// 		case 1:
// 			prf.Inject(KEY_GENE_PRF_2, v)
// 		}
// 	}
// 	return &prf
// }
