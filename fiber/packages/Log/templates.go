package log

import (
	"gotomate-astilectron/fiber/template"
)

// PrintTemplate Dialog's LogPrint Template
var PrintTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Log:",
		},
		Input: template.TextInput{
			Bind:         "Log",
			BindVariable: "VarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "LogIsVar",
		},
	}.Build(),
}
