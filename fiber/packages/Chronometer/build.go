package chronometer

import "gotomate/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "End":
		return new(EndDatabinder), EndTemplate
	case "Start":
		return new(StartDatabinder), StartTemplate
	default:
		return nil, nil
	}
}
