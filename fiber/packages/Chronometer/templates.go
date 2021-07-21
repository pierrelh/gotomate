package chronometer

import (
	"gotomate-astilectron/fiber/template"
)

// EndTemplate Dialog's End Template
var EndTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Chronometer:",
		},
		Input: template.TextInput{
			BindVariable: "ChronoVarName",
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
			BindVariable: "Output",
		},
	}.Build(),
}

// StartTemplate Dialog's Start Template
var StartTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			BindVariable: "Output",
		},
	}.Build(),
}
