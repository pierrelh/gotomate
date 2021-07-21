package clipboard

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a clipboard instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Write":
		return new(WriteDatabinder), WriteTemplate
	case "Read":
		return new(ReadDatabinder), ReadTemplate
	}
	return nil, nil
}
