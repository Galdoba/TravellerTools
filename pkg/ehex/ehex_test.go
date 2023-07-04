package ehex

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	for i := -5; i < 65; i++ {
		ehex := New().Set(i)
		if ehex.code == "?" && ehex.value == -2 && ehex.comment == "unknown" {
			continue
		}
		fmt.Println(ehex.code, ehex.value)

		//fmt.Printf("i:=%v Comment (%v) is expected\n%v\n", i, ehex.Meaning(), ehex.printStruct())
	}

}

func codeUnexpected(code string) bool {
	result := true
	for _, val := range []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
		"L", "M", "N", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z", "?",
	} {
		if val == code {
			result = false
		}
	}
	return result
}

func valueUnexpected(v int) bool {
	result := true
	for _, val := range []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
		22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		33, -2, -3} {
		if val == v {
			result = false
		}
	}
	return result
}

type testTableSetCodes struct {
	inputData string
	expected  ehex
}

func TestSetCodes(t *testing.T) {
	ehexObj := New()
	testTable := []testTableSetCodes{
		{
			inputData: "a",
			expected:  ehex{value: 10, code: "A"},
		},
		{
			inputData: "A",
			expected:  ehex{value: 10, code: "A"},
		},
		{
			inputData: "ф",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "%",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "AB",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "*",
			expected:  ehex{value: -3, code: "*", comment: "any value"},
		},
		{
			inputData: "?",
			expected:  ehex{value: -2, code: "?", comment: "unknown"},
		},
	}
	////////////
	for i, val := range testTable {
		ehexObj.Set(val.inputData)
		if !match(*ehexObj, val.expected) {
			t.Errorf("input %v: (%v) No Match: \nexecuted: %v\nexpected: %v\n", i+1, val.inputData, ehexObj.printStruct(), val.expected.printStruct())
		}
	}
}

func (ehex *ehex) printStruct() string {
	return fmt.Sprintf("ehex{code: '%v', value: '%v', comment: '%v'}", ehex.code, ehex.value, ehex.comment)
}

func match(ehex1, ehex2 ehex) bool {
	return ehex1.printStruct() == ehex2.printStruct()
}

type testStruct struct {
	Str string
	Val int
}

func TestContained(t *testing.T) {

	eh := New().Set("A")
	eh.Put(&testStruct{"aaa", 111})
	fmt.Println("eh", eh)
	eh.SetContainedValFunc(func() *int {
		return &eh.Container.(*testStruct).Val
	})
	fmt.Println("eh.Value()", eh.Value())

	fmt.Println("eh.ContainedVal()", eh.ContainedVal())
	fmt.Println("eh.Container.(*testStruct).Val", eh.Container.(*testStruct).Val)

}
