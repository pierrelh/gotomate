package file

import "gotomate/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Create":
		return new(CreateDatabinder), CreateTemplate
	case "Delete":
		return new(DeleteDatabinder), DeleteTemplate
	case "Read":
		return new(ReadDatabinder), ReadTemplate
	case "Write":
		return new(WriteDatabinder), WriteTemplate
	default:
		return nil, nil
	}
}
