package cmd

import (
	"fmt"
	"testing"
)

func TestSystemCreation(t *testing.T) {
	for i := 200; i < 220; i++ {
		name := fmt.Sprintf("name_%v", i)
		ssd, err := NewSystem(name)
		if err != nil {
			t.Errorf("func err: %v", err.Error())
		}
		fmt.Printf("tested %v\r", i)
		fmt.Println(ssd)
	}

}
