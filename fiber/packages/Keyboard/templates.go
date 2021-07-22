package keyboard

import (
	"gotomate/fiber/template"
)

// Inputs Define the possibles values of keyboard specials inputs
func Inputs() []template.Option {
	return []template.Option{
		{Name: "", Value: ""},
		{Name: "alt", Value: "alt"},
		{Name: "cmd", Value: "cmd"},
		{Name: "shift", Value: "shift"},
		{Name: "ctrl", Value: "ctrl"},
		{Name: "enter", Value: "enter"},
	}
}

// TapTemplate Dialog's KeyboardTap Template
var TapTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Input:",
		},
		Input: template.TextInput{
			Bind:      "Input",
			MaxLength: 1,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Special Input 1:",
		},
		Input: template.Select{
			Bind:    "Special1",
			Options: Inputs(),
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Special Input 2:",
		},
		Input: template.Select{
			Bind:    "Special2",
			Options: Inputs(),
		},
	}.Build(),
}

// TypeTemplate Dialog's KeyboardType Template
var TypeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Input:",
		},
		Input: template.TextInput{
			Bind:         "Input",
			BindVariable: "VarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "InputIsVar",
		},
	}.Build(),
}
