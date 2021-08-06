package command

import "gotomate/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Dir":
		return new(DirDatabinder), DirTemplate
	case "GetOutput":
		return new(GetOutputDatabinder), GetOutputTemplate
	case "New":
		return new(NewDatabinder), NewTemplate
	case "Run":
		return new(RunDatabinder), RunTemplate
	case "Start":
		return new(StartDatabinder), StartTemplate
	default:
		return nil, nil
	}
}
