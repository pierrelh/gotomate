package flow

import (
	"gotomate/fiber/template"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Start":
		return new(StartDatabinder), StartTemplate
	case "End":
		return nil, nil
	}
	return nil, nil
}
