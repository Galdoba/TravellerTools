package character

import "fmt"

const (
	EVENT_RollCharacteristics = "Roll Characteristics"
	EVENT_DETERMINE_HW        = "Determine Homeworld"
	EVENT_CHOOSE_BGSKILLS     = "Choose Background Skills"
	EVENT_CHOOSE_CAREER       = "Choose Career"
	EVENT_CAREER_CYCLE        = "Career Cycle"
	EVENT_BENEFITS            = "Asign Benefits"
	EVENT_INJURY              = "Injury"
	EVENT_EndGeneration       = "END GENERATION"
)

func (g *generator) Generate_By_Events() (*Character, error) {
	ch := NewCharacter()
	if _, ok := g.options[KeyManual]; ok {
		ch.setAsPC()
	}
	err := fmt.Errorf("Generation not started")
	for ch.nextEvent != EVENT_EndGeneration {
		fmt.Println("Commence: " + ch.nextEvent)
		switch ch.nextEvent {
		default:
			panic(ch.nextEvent + " not implemented")
		case EVENT_RollCharacteristics:
			err = ch.RollCharacteristics(g.options)
			ch.nextEvent = EVENT_DETERMINE_HW
		case EVENT_DETERMINE_HW:
			err = ch.DetermineHomeworld(g.options)
		case EVENT_CHOOSE_BGSKILLS:
			err = ch.ChooseBackgroundSkills(g.options)
		case EVENT_CHOOSE_CAREER:
			err = ch.Injury()
			// err = ch.ChooseAndStartCareer(g.options)
		case EVENT_CAREER_CYCLE:
			err = ch.CareerCycle(g.options)
		case EVENT_BENEFITS:
			err = ch.ConsumeBenefits()
		case EVENT_INJURY:
			err = ch.Injury()
		}
		if err != nil {
			ch.Inform("WARNING: " + err.Error())
			panic(0)
		}
		ch.FlushScreen()
	}

	// ch.ChooseBackgroundSkills(g.options)
	// if err := ch.ChooseAndStartCareer(g.options); err != nil {
	// 	return ch, err
	// }
	// if err := ch.CareerCycle(g.options); err != nil {
	// 	return ch, err
	// }
	// for _, benefit := range ch.Benefits {
	// 	if err := ch.gain(benefit); err != nil {
	// 		return ch, fmt.Errorf("gain benefit: %v", err.Error())
	// 	}
	// }
	// ch.Benefits, ch.Balance = confirmBenefits(ch.Benefits)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	ch.FlushScreen()
	return ch, nil
}
