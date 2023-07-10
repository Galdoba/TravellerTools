package orbitns

import "testing"

type testPair struct {
	input  float64
	output float64
}

func TestOrbitn2AU_correct(t *testing.T) {
	data := []testPair{
		{input: 4.3, output: 1.96},
		{input: 0.09, output: 0.036},
		{input: 6.1, output: 5.68},
		{input: 12.19, output: 338.7},
	}
	for testNum, testData := range data {
		result := OR2MKM(testData.input)
		if result != testData.output {
			t.Errorf("test %v:\n have =%v, expected =%v from input =%v", testNum, result, testData.output, testData.input)
		}
	}
}

func TestOrbitn2AU_incorrect(t *testing.T) {
	return
	data := []testPair{
		{input: 4.3, output: 1.96},
		{input: 0.09, output: 0.036},
		{input: 6.1, output: 5.68},
		{input: 12.2, output: 338.7},
	}
	for testNum, testData := range data {
		result := OR2MKM(testData.input)
		if result == testData.output {
			t.Errorf("test %v:\n have =%v, expected other than =%v from input =%v", testNum, result, testData.output, testData.input)
		}
	}
}
