package survey

import (
	"fmt"
	"strings"
)

type constructorInstruction struct {
	field int
	data  string
}

func NewSecondSurvey(data ...constructorInstruction) (*SecondSurveyData, error) {
	rawData := []string{}
	for len(rawData) < END_OF_SURVEY_DATA+1 {
		rawData = append(rawData, "")
	}
	for _, instruction := range data {
		if instruction.field < reserved || instruction.field > END_OF_SURVEY_DATA {
			return &SecondSurveyData{}, fmt.Errorf("incorect instruction: '%v-%v'", instruction.field, instruction.data)
		}
		rawData[instruction.field] = instruction.data
	}
	rawString := strings.Join(rawData, "|")
	ssd := Parse(rawString)
	return ssd, nil
}

func Instruction(field int, data string) constructorInstruction {
	return constructorInstruction{field, data}
}
