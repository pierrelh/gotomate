package api

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Get":
		return new(GetDatabinder), GetTemplate
	case "Post":
		return new(PostDatabinder), PostTemplate
	}
	return nil, nil
}
