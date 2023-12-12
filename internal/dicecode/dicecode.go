package dicecode

import (
	"fmt"
	"strconv"
	"strings"
)

//DetectValue(val int, codes []string) string
// val
// val-
// val+
// va1...val2  (-val1)...(+val2)
// val1 val2 val3...

const (
	methodUndefined = iota
	methodSingle
	methodAndMore
	methodAndLess
	methodBounds
	methodMany
)

func DetectMatch_Unsafe(r int, codes []string) string {
	code, err := DetectMatch(r, codes)
	if err != nil {
		panic(err.Error())
	}
	return code
}

// DetectValue - returns dicecode that matches value i
func DetectMatch(rollValue int, codes []string) (string, error) {
	out := ""
	vals := []int{}
	for _, code := range codes {
		switch detectMethod(code) {
		case methodUndefined:
			return "", fmt.Errorf("detection method undetermined: [%v]", code)
		case methodSingle:
			vals = parseValidValues_m1(code)
		case methodAndMore:
			vals = parseValidValues_m2(code)
		case methodAndLess:
			vals = parseValidValues_m3(code)
		case methodBounds:
			vals = parseValidValues_m4(code)
		case methodMany:
			vals = parseValidValues_m5(code)
		}
		if len(vals) == 0 {
			return "", fmt.Errorf("failed to parse code [%v]", code)
		}
		for _, validValue := range vals {
			if validValue != rollValue {
				continue
			}
			if out != "" {
				return "", fmt.Errorf("multiple matches detected: [%v] [%v]", out, code)
			}
			out = code

		}
	}
	if out == "" {
		return "", fmt.Errorf("value not detected")
	}
	return out, nil
}

func detectMethod(code string) int {
	if _, err := strconv.Atoi(code); err == nil {
		return methodSingle
	}
	if len(strings.Fields(code)) > 1 {
		return methodMany
	}
	if strings.HasSuffix(code, "+") {
		return methodAndMore
	}
	if strings.HasSuffix(code, "-") {
		return methodAndLess
	}
	if strings.Contains(code, "...") {
		return methodBounds
	}

	return methodUndefined
}

func parseValidValues_m1(s string) []int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return []int{}
	}
	vals := append([]int{}, v)
	return vals
}

func parseValidValues_m2(s string) []int {
	s = strings.TrimSuffix(s, "+")
	v, err := strconv.Atoi(s)
	if err != nil {
		return []int{}
	}
	vals := append([]int{}, v)
	for i := 1; i <= 60; i++ {
		vals = append(vals, v+i)
	}
	return vals
}

func parseValidValues_m3(s string) []int {
	s = strings.TrimSuffix(s, "-")
	v, err := strconv.Atoi(s)
	if err != nil {
		return []int{}
	}
	vals := append([]int{}, v)
	for i := 1; i <= 60; i++ {
		vals = append(vals, v-i)
	}
	return vals
}

func parseValidValues_m4(s string) []int {
	data := strings.Split(s, "...")
	if len(data) != 2 {
		return []int{}
	}
	v1, err := strconv.Atoi(data[0])
	if err != nil {
		return []int{}
	}
	v2, err := strconv.Atoi(data[1])
	if err != nil {
		return []int{}
	}
	if v1 >= v2 {
		v1, v2 = v2, v1
	}
	vals := []int{}
	for i := v1; i <= v2; i++ {
		vals = append(vals, i)
	}
	return vals
}

func parseValidValues_m5(s string) []int {
	data := strings.Fields(s)
	vals := []int{}
	rslt := []int{}
	for _, dat := range data {
		switch detectMethod(dat) {
		case methodUndefined, methodMany:
			return nil
		case methodSingle:
			vals = parseValidValues_m1(dat)
		case methodAndMore:
			vals = parseValidValues_m2(dat)
		case methodAndLess:
			vals = parseValidValues_m3(dat)
		case methodBounds:
			vals = parseValidValues_m4(dat)
		}
		if len(vals) == 0 {
			return nil
		}
		rslt = append(rslt, vals...)
	}
	return rslt
}
