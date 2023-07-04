package profile2

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func TestProfile(t *testing.T) {
	keys := []string{}
	ptype := ""
	prf := New()
	prf.Narrow(ptype, keys)
	keys = append(keys, "C3")
	for i, k := range keys {
		if err := prf.Create(k); err != nil {
			t.Errorf("%v %v %v", i, k, err.Error())
		}
		eh, err := prf.Read(k)
		if err != nil {
			t.Errorf("%v %v %v", i, k, err.Error())
		}
		fmt.Println(eh)
		eh2 := ehex.New().Set(i + 3)
		if err := prf.Update(k, eh2); err != nil {
			t.Errorf("%v %v %v", i, k, err.Error())
		}
		eh3, err2 := prf.Read(k)
		fmt.Println(eh3, err2)
		if err := prf.Delete(k); err != nil {
			t.Errorf("%v %v %v", i, k, err.Error())
		}
	}

}
