package algorithmic

import (
	"gotomate-astilectron/fiber/template"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "For":
		return new(ForDatabinder), ForTemplate
	case "If":
		return new(IfDatabinder), IfTemplate
	}
	return nil, nil
}
