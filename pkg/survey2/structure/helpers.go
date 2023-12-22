package structure

func disassemble(stellar string) (string, string, string) {
	for _, tp := range []string{Type_O, Type_B, Type_A, Type_F, Type_G, Type_K, Type_M} {
		for _, stp := range []string{SubType_0, SubType_1, SubType_2, SubType_3, SubType_4, SubType_5, SubType_6, SubType_7, SubType_8, SubType_9} {
			for _, cls := range []string{Class_Ia, Class_Ib, Class_II, Class_III, Class_IV, Class_V, Class_VI} {
				if stellar == tp+stp+" "+cls {
					return tp, stp, cls
				}
			}
		}
	}
	switch stellar {
	case Class_BD, Class_D, Class_NS, Class_PSR, Class_MGR, Class_BH:
		return "", "", stellar
	}
	return "", "", ""
}

func boundInt(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}
