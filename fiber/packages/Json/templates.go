package json

import (
	"gotomate-astilectron/fiber/template"
)

// CreateJsonTemplate Dialog's Define an CreateJson Template
var CreateJsonTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Path to import:",
		},
		Input: template.TextInput{
			Bind:         "Path",
			BindVariable: "PathVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "PathIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// JsonToDictionaryTemplate Dialog's Define an JsonToDictionary Template
var JsonToDictionaryTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Json to convert:",
		},
		Input: template.TextInput{
			BindVariable: "JsonVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// SaveJsonTemplate Dialog's Define SaveJson Template
var SaveJsonTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Json to save:",
		},
		Input: template.TextInput{
			BindVariable: "JsonVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Path to save:",
		},
		Input: template.TextInput{
			Bind:         "Path",
			BindVariable: "PathVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "PathIsVar",
		},
	}.Build(),
}
