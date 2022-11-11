package language

import (
	"fmt"
	"strings"
)

func callTables(name string) ([]Sound, []Sound, []Sound, map[string]string, map[string]string, map[string]string, map[string]string, map[string]string) {
	nameUp := strings.ToUpper(name)
	switch nameUp {
	default:
		fmt.Printf("Language '%v' is not implemented", name)
	case "ZHODANI":
		return zhodaniInitialConsonants(),
			zhodaniVowels(),
			zhodaniFinalConsonants(),
			zhodaniBasicGenerationTable(),
			zhodaniAlternativeGenerationTable(),
			zhodaniInitialConsonantTable(),
			zhodaniVowelTable(),
			zhodaniFinalConsonantTable()
	case "VILANI":
		return vilaniInitialConsonants(),
			vilaniVowels(),
			vilaniFinalConsonants(),
			vilaniBasicGenerationTable(),
			vilaniAlternativeGenerationTable(),
			vilaniInitialConsonantTable(),
			vilaniVowelTable(),
			vilaniFinalConsonantTable()
	}
	//must not happen
	return []Sound{}, []Sound{}, []Sound{}, nil, nil, nil, nil, nil
}

func zhodaniInitialConsonants() []Sound {
	return []Sound{
		{InitialConsonant, "B", 3, "b"},
		{InitialConsonant, "BL", 2, "bl"},
		{InitialConsonant, "BR", 3, "br"},
		{InitialConsonant, "CH", 3, "ch"},
		{InitialConsonant, "CHT", 7, "cht"},
		{InitialConsonant, "D", 6, "d"},
		{InitialConsonant, "DL", 4, "ddle"},
		{InitialConsonant, "DR", 3, "dr"},
		{InitialConsonant, "F", 3, "f"},
		{InitialConsonant, "FL", 2, "fl"},
		{InitialConsonant, "FR", 2, "fr"},
		{InitialConsonant, "J", 4, "j"},
		{InitialConsonant, "JD", 3, "ged"},
		{InitialConsonant, "K", 3, "k"},
		{InitialConsonant, "KL", 1, "ckle"},
		{InitialConsonant, "KR", 1, "cker"},
		{InitialConsonant, "L", 2, "l"},
		{InitialConsonant, "M", 1, "m"},
		{InitialConsonant, "N", 5, "n"},
		{InitialConsonant, "P", 4, "p"},
		{InitialConsonant, "PL", 4, "pl"},
		{InitialConsonant, "PR", 2, "pr"},
		{InitialConsonant, "Q", 1, "k"},
		{InitialConsonant, "QL", 1, "kl"},
		{InitialConsonant, "QR", 1, "kr"},
		{InitialConsonant, "R", 3, "r"},
		{InitialConsonant, "S", 4, "s"},
		{InitialConsonant, "SH", 4, "sh"},
		{InitialConsonant, "SHT", 4, "sht"},
		{InitialConsonant, "ST", 4, "st"},
		{InitialConsonant, "T", 3, "t"},
		{InitialConsonant, "TL", 6, "tl"},
		{InitialConsonant, "TS", 2, "ts"},
		{InitialConsonant, "V", 3, "v"},
		{InitialConsonant, "VL", 1, "vl"},
		{InitialConsonant, "VR", 1, "vr"},
		{InitialConsonant, "Y", 2, "y"},
		{InitialConsonant, "Z", 3, "z"},
		{InitialConsonant, "ZD", 6, "zd"},
		{InitialConsonant, "ZH", 4, "zh"},
		{InitialConsonant, "ZHD", 6, "zhd"},
	}
}

func zhodaniVowels() []Sound {
	return []Sound{
		{Vowel, "A", 7, "o"},
		{Vowel, "E", 8, "e"},
		{Vowel, "I", 5, "i"},
		{Vowel, "IA", 4, "ya"},
		{Vowel, "IE", 4, "ye"},
		{Vowel, "O", 2, "o"},
		{Vowel, "R", 1, "rz"},
	}
}

func zhodaniFinalConsonants() []Sound {
	return []Sound{
		{FinalConsonant, "B", 1, "b"},
		{FinalConsonant, "BL", 4, "bl"},
		{FinalConsonant, "BR", 4, "bor"},
		{FinalConsonant, "CH", 3, "ch"},
		{FinalConsonant, "D", 2, "d"},
		{FinalConsonant, "DL", 4, "ddle"},
		{FinalConsonant, "DR", 4, "dder"},
		{FinalConsonant, "F", 3, "ff"},
		{FinalConsonant, "FL", 3, "ffle"},
		{FinalConsonant, "FR", 3, "fr"},
		{FinalConsonant, "J", 2, "ge"},
		{FinalConsonant, "K", 1, "k"},
		{FinalConsonant, "KL", 2, "ckle"},
		{FinalConsonant, "KR", 1, "cker"},
		{FinalConsonant, "L", 7, "ll"},
		{FinalConsonant, "M", 1, "m"},
		{FinalConsonant, "N", 1, "n"},
		{FinalConsonant, "NCH", 4, "nch"},
		{FinalConsonant, "NJ", 3, "nj"},
		{FinalConsonant, "NS", 3, "ns"},
		{FinalConsonant, "NSH", 4, "nsh"},
		{FinalConsonant, "NT", 2, "nt"},
		{FinalConsonant, "NTS", 2, "nts"},
		{FinalConsonant, "NZ", 3, "nz"},
		{FinalConsonant, "NZH", 4, "nzh"},
		{FinalConsonant, "P", 1, "p"},
		{FinalConsonant, "PL", 4, "ppl"},
		{FinalConsonant, "PR", 4, "pr"},
		{FinalConsonant, "Q", 1, "k"},
		{FinalConsonant, "QL", 1, "kl"},
		{FinalConsonant, "QR", 1, "kr"},
		{FinalConsonant, "R", 3, "r"},
		{FinalConsonant, "SH", 4, "sh"},
		{FinalConsonant, "T", 2, "t"},
		{FinalConsonant, "TS", 4, "ts"},
		{FinalConsonant, "TL", 5, "tl"},
		{FinalConsonant, "V", 3, "ve"},
		{FinalConsonant, "VL", 2, "vl"},
		{FinalConsonant, "VR", 3, "vr"},
		{FinalConsonant, "Z", 5, "z"},
		{FinalConsonant, "ZH", 4, "zh"},
		{FinalConsonant, "'", 4, "'"},
	}
}

func zhodaniBasicGenerationTable() map[string]string {
	tab := make(map[string]string)
	codes := valid2()
	syl := []string{
		"V", "V", "V", "CV", "CV", "CV",
		"VC", "VC", "VC", "VC", "VC", "VC",
		"VC", "VC", "VC", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
	}
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func zhodaniAlternativeGenerationTable() map[string]string {
	tab := make(map[string]string)
	codes := valid2()
	syl := []string{
		"V", "V", "V", "V", "V", "V",
		"CV", "CV", "CV", "CV", "CV", "CV",
		"VC", "VC", "VC", "VC", "VC", "VC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
	}
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func zhodaniInitialConsonantTable() map[string]string {
	tab := make(map[string]string)
	codes := valid3()
	syl := []string{}
	syl = append(syl, []string{
		"B", "B", "B", "B", "B", "BL",
		"BL", "BL", "BR", "BR", "BR", "BR",
		"BR", "CH", "CH", "CH", "CH", "CH",
		"CH", "CH", "CH", "CH", "CH", "CH",
		"CH", "CHT", "CHT", "CHT", "CHT", "CHT",
		"CHT", "CHT", "D", "D", "D", "D",
	}...)
	syl = append(syl, []string{
		"D", "D", "D", "D", "D", "DL",
		"DL", "DL", "DL", "DL", "DL", "DL",
		"DR", "DR", "DR", "DR", "DR", "F",
		"F", "F", "F", "F", "FL", "FL",
		"FL", "FR", "FR", "FR", "J", "J",
		"J", "J", "J", "J", "J", "JD",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func zhodaniVowelTable() map[string]string {
	tab := make(map[string]string)
	codes := valid3()
	syl := []string{}
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func zhodaniFinalConsonantTable() map[string]string {
	tab := make(map[string]string)
	codes := valid3()
	syl := []string{}
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	syl = append(syl, []string{
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "", "", "",
	}...)
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

/////////////////////////////////

func vilaniInitialConsonants() []Sound {
	return []Sound{
		{InitialConsonant, "B", 3, "b"},
		{InitialConsonant, "BL", 2, "bl"},
		{InitialConsonant, "BR", 3, "br"},
		{InitialConsonant, "CH", 3, "ch"},
		{InitialConsonant, "CHT", 7, "cht"},
		{InitialConsonant, "D", 6, "d"},
		{InitialConsonant, "DL", 4, "ddle"},
		{InitialConsonant, "DR", 3, "dr"},
		{InitialConsonant, "F", 3, "f"},
		{InitialConsonant, "FL", 2, "fl"},
		{InitialConsonant, "FR", 2, "fr"},
		{InitialConsonant, "J", 4, "j"},
		{InitialConsonant, "JD", 3, "ged"},
		{InitialConsonant, "K", 3, "k"},
		{InitialConsonant, "KL", 1, "ckle"},
		{InitialConsonant, "KR", 1, "cker"},
		{InitialConsonant, "L", 2, "l"},
		{InitialConsonant, "M", 1, "m"},
		{InitialConsonant, "N", 5, "n"},
		{InitialConsonant, "P", 4, "p"},
		{InitialConsonant, "PL", 4, "pl"},
		{InitialConsonant, "PR", 2, "pr"},
		{InitialConsonant, "Q", 1, "k"},
		{InitialConsonant, "QL", 1, "kl"},
		{InitialConsonant, "QR", 1, "kr"},
		{InitialConsonant, "R", 3, "r"},
		{InitialConsonant, "S", 4, "s"},
		{InitialConsonant, "SH", 4, "sh"},
		{InitialConsonant, "SHT", 4, "sht"},
		{InitialConsonant, "ST", 4, "st"},
		{InitialConsonant, "T", 3, "t"},
		{InitialConsonant, "TL", 6, "tl"},
		{InitialConsonant, "TS", 2, "ts"},
		{InitialConsonant, "V", 3, "v"},
		{InitialConsonant, "VL", 1, "vl"},
		{InitialConsonant, "VR", 1, "vr"},
		{InitialConsonant, "Y", 2, "y"},
		{InitialConsonant, "Z", 3, "z"},
		{InitialConsonant, "ZD", 6, "zd"},
		{InitialConsonant, "ZH", 4, "zh"},
		{InitialConsonant, "ZHD", 6, "zhd"},
	}
}

func vilaniVowels() []Sound {
	return []Sound{
		{Vowel, "A", 7, "o"},
		{Vowel, "E", 8, "e"},
		{Vowel, "I", 5, "i"},
		{Vowel, "IA", 4, "ya"},
		{Vowel, "IE", 4, "ye"},
		{Vowel, "O", 2, "o"},
		{Vowel, "R", 1, "rz"},
	}
}

func vilaniFinalConsonants() []Sound {
	return []Sound{
		{FinalConsonant, "B", 1, "b"},
		{FinalConsonant, "BL", 4, "bl"},
		{FinalConsonant, "BR", 4, "bor"},
		{FinalConsonant, "CH", 3, "ch"},
		{FinalConsonant, "D", 2, "d"},
		{FinalConsonant, "DL", 4, "ddle"},
		{FinalConsonant, "DR", 4, "dder"},
		{FinalConsonant, "F", 3, "ff"},
		{FinalConsonant, "FL", 3, "ffle"},
		{FinalConsonant, "FR", 3, "fr"},
		{FinalConsonant, "J", 2, "ge"},
		{FinalConsonant, "K", 1, "k"},
		{FinalConsonant, "KL", 2, "ckle"},
		{FinalConsonant, "KR", 1, "cker"},
		{FinalConsonant, "L", 7, "ll"},
		{FinalConsonant, "M", 1, "m"},
		{FinalConsonant, "N", 1, "n"},
		{FinalConsonant, "NCH", 4, "nch"},
		{FinalConsonant, "NJ", 3, "nj"},
		{FinalConsonant, "NS", 3, "ns"},
		{FinalConsonant, "NSH", 4, "nsh"},
		{FinalConsonant, "NT", 2, "nt"},
		{FinalConsonant, "NTS", 2, "nts"},
		{FinalConsonant, "NZ", 3, "nz"},
		{FinalConsonant, "NZH", 4, "nzh"},
		{FinalConsonant, "P", 1, "p"},
		{FinalConsonant, "PL", 4, "ppl"},
		{FinalConsonant, "PR", 4, "pr"},
		{FinalConsonant, "Q", 1, "k"},
		{FinalConsonant, "QL", 1, "kl"},
		{FinalConsonant, "QR", 1, "kr"},
		{FinalConsonant, "R", 3, "r"},
		{FinalConsonant, "SH", 4, "sh"},
		{FinalConsonant, "T", 2, "t"},
		{FinalConsonant, "TS", 4, "ts"},
		{FinalConsonant, "TL", 5, "tl"},
		{FinalConsonant, "V", 3, "ve"},
		{FinalConsonant, "VL", 2, "vl"},
		{FinalConsonant, "VR", 3, "vr"},
		{FinalConsonant, "Z", 5, "z"},
		{FinalConsonant, "ZH", 4, "zh"},
		{FinalConsonant, "'", 4, "'"},
	}
}

func vilaniBasicGenerationTable() map[string]string {
	tab := make(map[string]string)
	codes := valid2()
	syl := []string{
		"V", "V", "V", "V", "V", "V",
		"CV", "CV", "CV", "CV", "CV", "CV",
		"CV", "CV", "CV", "CV", "CV", "CV",
		"CV", "CV", "CV", "VC", "VC", "VC",
		"VC", "VC", "VC", "VC", "VC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
	}
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func vilaniAlternativeGenerationTable() map[string]string {
	tab := make(map[string]string)
	codes := valid2()
	syl := []string{
		"CV", "CV", "CV", "CV", "CV", "CV",
		"CV", "CV", "CV", "CV", "CV", "CV",
		"CV", "CV", "CV", "CV", "CV", "CV",
		"CV", "CV", "CV", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
		"CVC", "CVC", "CVC", "CVC", "CVC", "CVC",
	}
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func vilaniInitialConsonantTable() map[string]string {
	tab := make(map[string]string)
	codes := valid3()
	syl := []string{}
	syl = append(syl, []string{
		"K", "K", "K", "K", "K", "K",
		"K", "K", "K", "K", "K", "K",
		"K", "K", "K", "K", "K", "K",
		"K", "K", "K", "K", "K", "K",
		"K", "K", "K", "K", "K", "K",
		"K", "K", "K", "K", "K", "K",
	}...)
	syl = append(syl, []string{
		"K", "K", "K", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
	}...)
	syl = append(syl, []string{
		"G", "G", "G", "G", "G", "G",
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "D", "D", "D",
		"D", "D", "D", "D", "D", "D",
	}...)
	syl = append(syl, []string{
		"D", "D", "D", "D", "D", "D",
		"D", "D", "D", "D", "D", "D",
		"L", "L", "L", "L", "L", "L",
		"L", "L", "L", "L", "L", "L",
		"L", "L", "L", "L", "L", "L",
		"L", "L", "L", "SH", "SH", "SH",
	}...)
	syl = append(syl, []string{
		"SH", "SH", "SH", "SH", "SH", "SH",
		"SH", "SH", "SH", "SH", "SH", "SH",
		"SH", "SH", "SH", "SH", "SH", "SH",
		"KH", "KH", "KH", "KH", "KH", "KH",
		"KH", "KH", "KH", "KH", "KH", "KH",
		"KH", "KH", "KH", "KH", "KH", "KH",
	}...)
	syl = append(syl, []string{
		"N", "N", "N", "N", "N", "N",
		"N", "N", "N", "N", "S", "S",
		"S", "S", "S", "S", "S", "S",
		"S", "S", "P", "P", "P", "P",
		"B", "B", "B", "B", "Z", "Z",
		"Z", "Z", "R", "R", "R", "R",
	}...)
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func vilaniVowelTable() map[string]string {
	tab := make(map[string]string)
	codes := valid3()
	syl := []string{}
	syl = append(syl, []string{
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
	}...)
	syl = append(syl, []string{
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "A", "A", "A", "A", "A",
		"A", "E", "E", "E", "E", "E",
	}...)
	syl = append(syl, []string{
		"E", "E", "E", "E", "E", "E",
		"E", "E", "E", "E", "E", "E",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
	}...)
	syl = append(syl, []string{
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "I",
		"I", "I", "I", "I", "I", "U",
	}...)
	syl = append(syl, []string{
		"U", "U", "U", "U", "U", "U",
		"U", "U", "U", "U", "U", "U",
		"U", "U", "U", "U", "U", "U",
		"U", "U", "U", "U", "U", "U",
		"U", "U", "U", "U", "U", "U",
		"U", "U", "U", "U", "U", "U",
	}...)
	syl = append(syl, []string{
		"U", "U", "U", "U", "AA", "AA",
		"AA", "AA", "AA", "AA", "AA", "AA",
		"II", "II", "II", "II", "II", "II",
		"II", "II", "II", "II", "II", "II",
		"II", "II", "II", "II", "UU", "UU",
		"UU", "UU", "UU", "UU", "UU", "UU",
	}...)
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

func vilaniFinalConsonantTable() map[string]string {
	tab := make(map[string]string)
	codes := valid3()
	syl := []string{}
	syl = append(syl, []string{
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
	}...)
	syl = append(syl, []string{
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
		"R", "R", "R", "R", "R", "R",
	}...)
	syl = append(syl, []string{
		"R", "R", "R", "R", "N", "N",
		"N", "N", "N", "N", "N", "N",
		"N", "N", "N", "N", "N", "N",
		"N", "N", "N", "N", "N", "N",
		"N", "N", "N", "N", "N", "N",
		"M", "M", "M", "M", "M", "M",
	}...)
	syl = append(syl, []string{
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "M", "M", "M",
		"M", "M", "M", "M", "M", "M",
		"M", "SH", "SH", "SH", "SH", "SH",
	}...)
	syl = append(syl, []string{
		"SH", "SH", "SH", "SH", "SH", "SH",
		"SH", "SH", "SH", "SH", "SH", "SH",
		"SH", "SH", "SH", "SH", "SH", "SH",
		"SH", "SH", "SH", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
		"G", "G", "G", "G", "G", "G",
	}...)
	syl = append(syl, []string{
		"S", "S", "S", "S", "S", "S",
		"S", "S", "S", "S", "S", "D",
		"D", "D", "D", "D", "D", "D",
		"D", "D", "D", "D", "D", "D",
		"P", "P", "P", "P", "P", "P",
		"K", "K", "K", "K", "K", "K",
	}...)
	for i, code := range codes {
		tab[code] = syl[i]
	}
	return tab
}

/////////////HELPERS
func valid2() []string {
	return []string{
		"11", "12", "13", "14", "15", "16",
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
		"61", "62", "63", "64", "65", "66",
	}
}

func valid3() []string {
	v3 := []string{}
	for _, v := range []string{"1", "2", "3", "4", "5", "6"} {
		for _, v2 := range valid2() {
			vc := v + v2
			v3 = append(v3, vc)
		}
	}
	return v3
}
