package profile2

import "github.com/Galdoba/TravellerTools/pkg/ehex"

type uniProfile struct {
	data          map[string]ehex.Ehex
	profileType   string
	exclusiveKeys []string
}

type Profile interface {
	Create(string)
	Read(string) ehex.Ehex
	Update(string, ehex.Ehex)
	Delete(string)
	Keys() []string
	Type() string
}

func matchByKeys(bigger, smaller Profile) bool {
	if len(bigger.Keys()) == 0 {
		return true
	}
	for _, keyS := range smaller.Keys() {
		if !inSlice(bigger.Keys(), keyS) {
			return false
		}
	}
	return true
}

func inSlice(sl []string, s string) bool {
	for _, ss := range sl {
		if s == ss {
			return true
		}
	}
	return false
}
