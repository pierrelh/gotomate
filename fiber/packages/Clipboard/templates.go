package clipboard

import (
	"gotomate/fiber/template"
)

// WriteTemplate Dialog's ClipboardWrite Template
var WriteTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Clipboard:",
		},
		Input: template.TextInput{
			Bind:         "Content",
			BindVariable: "VarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "IsVar",
		},
	}.Build(),
}

// ReadTemplate Dialog's ClipboardRead Template
var ReadTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}
