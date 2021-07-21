package json

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "CreateJson":
		return new(CreateJsonDatabinder), CreateJsonTemplate
	case "JsonToDictionary":
		return new(JsonToDictionaryDatabinder), JsonToDictionaryTemplate
	case "SaveJson":
		return new(SaveJsonDatabinder), SaveJsonTemplate
	}
	return nil, nil
}
