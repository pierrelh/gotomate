package keyboard

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a keyboard instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Tap":
		return new(TapDatabinder), TapTemplate
	case "Type":
		return new(TypeDatabinder), TypeTemplate
	}
	return nil, nil
}
