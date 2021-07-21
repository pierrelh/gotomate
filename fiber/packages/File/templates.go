package file

import (
	"gotomate-astilectron/fiber/template"
)

// CreateTemplate Dialog's Create Template
var CreateTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Path:",
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

// DeleteTemplate Dialog's Delete Template
var DeleteTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Path:",
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

// ReadTemplate Dialog's Read Template
var ReadTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Path:",
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

// WriteTemplate Dialog's Write Template
var WriteTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Path:",
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
			Text: "Content:",
		},
		Input: template.TextInput{
			Bind:         "Content",
			BindVariable: "ContentVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ContentIsVar",
		},
	}.Build(),
}
