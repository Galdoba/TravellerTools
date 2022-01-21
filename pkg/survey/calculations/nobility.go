package calculations

import (
	"fmt"
	"strings"

	"github.com/Galdoba/utils"
)

const (
	Knight   = "B"
	Baronet  = "c"
	Baron    = "C"
	Marquis  = "D"
	Viscount = "e"
	Count    = "E"
	Duke     = "f"
	Duke2    = "F"
	Archduke = "G"
)

func NobilityErrors(nob string, tc []string, ix int) []error {
	var allErrs []error
	if nob == "Nobl?" {
		allErrs = append(allErrs, fmt.Errorf("Nobility undefined"))
	}
	if nob == "" {
		allErrs = append(allErrs, fmt.Errorf("Nobility undefined"))
	}
	if !strings.Contains(nob, Knight) {
		allErrs = append(allErrs, fmt.Errorf("Knight is valid for all worlds"))
	}
	if strings.Contains(nob, Baronet) {
		switch {
		case utils.ListContains(tc, "Pa"):
		case utils.ListContains(tc, "Pr"):
		default:
			allErrs = append(allErrs, fmt.Errorf("Baronet is valid only for worlds having 'Pa' or 'Pr'"))
		}
	}
	if strings.Contains(nob, Baron) {
		switch {
		case utils.ListContains(tc, "Ag"):
		case utils.ListContains(tc, "Ri"):
		default:
			allErrs = append(allErrs, fmt.Errorf("Baron is valid only for worlds having 'Ag' or 'Ri'"))
		}
	}
	if strings.Contains(nob, Marquis) {
		switch {
		case utils.ListContains(tc, "Pi"):
		default:
			allErrs = append(allErrs, fmt.Errorf("Marquis is valid only for worlds having 'Pi'"))
		}
	}
	if strings.Contains(nob, Viscount) {
		switch {
		case utils.ListContains(tc, "Ph"):
		default:
			allErrs = append(allErrs, fmt.Errorf("Viscount is valid only for worlds having 'Ph'"))
		}
	}
	if strings.Contains(nob, Count) {
		switch {
		case utils.ListContains(tc, "In"):
		case utils.ListContains(tc, "Hi"):
		default:
			allErrs = append(allErrs, fmt.Errorf("Count is valid only for worlds having 'In' or 'Hi'"))
		}
	}
	if strings.Contains(nob, Duke) {
		switch {
		case ix >= 4:
		default:
			allErrs = append(allErrs, fmt.Errorf("Duke is valid only for worlds having Ix 4+"))
		}
	}
	if strings.Contains(nob, Duke2) {
		switch {
		case utils.ListContains(tc, "Cs"):
		case utils.ListContains(tc, "Cp"):
		case utils.ListContains(tc, "Cx"):
		default:
			allErrs = append(allErrs, fmt.Errorf("Duke2 is valid only for Sector Capital worlds"))
		}
	}

	return allErrs
}

func FixNobility(tc []string, ix int) string {
	nob := Knight
	switch {
	case utils.ListContains(tc, "Pa") || utils.ListContains(tc, "Pr"):
		nob += Baronet
	case utils.ListContains(tc, "Ag") || utils.ListContains(tc, "Ri"):
		nob += Baron
	case utils.ListContains(tc, "Pi"):
		nob += Marquis
	case utils.ListContains(tc, "Ph"):
		nob += Viscount
	case utils.ListContains(tc, "In") || utils.ListContains(tc, "Hi"):
		nob += Count
	case ix >= 4:
		nob += Duke
	case utils.ListContains(tc, "Cs") || utils.ListContains(tc, "Cp") || utils.ListContains(tc, "Cx"):
		nob += Duke2
	}
	return nob
}
