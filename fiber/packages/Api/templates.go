package api

import (
	"gotomate/fiber/template"
)

// GetTemplate Dialog's Get Template
var GetTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Path:",
		},
		Input: template.UrlInput{
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

// PostTemplate Dialog's Post Template
var PostTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Dictionary to Post:",
		},
		Input: template.TextInput{
			BindVariable: "DataVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Path:",
		},
		Input: template.UrlInput{
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
