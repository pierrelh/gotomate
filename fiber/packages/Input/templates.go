package input

import (
	"gotomate-astilectron/fiber/template"
)

// InputTemplate Dialog's Input Template
var InputTemplate = &template.Template{
	template.Field{
		Label: template.Label{
			Text: "Message:",
		},
		Input: template.TextInput{
			Bind:         "Message",
			BindVariable: "MessageVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "MessageIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Output var:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}
