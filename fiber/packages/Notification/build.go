package notification

import "gotomate/fiber/template"

// Build Return the right databinder & the right template for a notification instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Create":
		return new(CreateDatabinder), CreateTemplate
	case "Push":
		return new(PushDatabinder), PushTemplate
	}
	return nil, nil
}
