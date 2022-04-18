package trvdb

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func emulateInput() []string {
	return []string{
		"drinax",
		"driNax",
		"asim",
	}
}

func Test_WorldByName(t *testing.T) {
	for _, inp := range emulateInput() {
		ssr, err := survey.Search(inp)
		fmt.Println(len(ssr), err)
		res, err := WorldByName(inp)
		if err != nil {
			t.Errorf("input (%v): internal error: %v", inp, err.Error())
		}
		if res == nil {
			t.Errorf("input (%v): search gave no results", inp)
		}
	}
}
