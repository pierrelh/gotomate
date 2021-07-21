package log

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a log instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Print":
		return new(PrintDatabinder), PrintTemplate
	}
	return nil, nil
}
