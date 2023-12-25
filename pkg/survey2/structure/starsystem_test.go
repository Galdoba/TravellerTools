package structure

import (
	"fmt"
	"testing"
)

func TestStarSystem(t *testing.T) {
	fmt.Println("Go TEST")
	ss, err := SystemFromProfile("5-G7 V-0.929-0.967-0.738-6.336:Ab-0.09-0.11-G8 V-0.907-0.957-0.681:B-6.1-0.8-K8 V-0.626-0.777-0.136:Ca-12.1-0.47-M0 V-0.510-0.728-0.0895:Cb-0.21-0.24-D-0.490-0.017-0.000525")

	fmt.Println("ss created")
	ss.TestOrbital = newVoid()
	bt, err := ss.MarshalStarSystem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bt))

}
