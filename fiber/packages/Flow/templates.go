package flow

import "gotomate-astilectron/fiber/template"

// StartTemplate Dialog's FlowStart Template
var StartTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Flow Name",
		},
		Input: template.TextInput{
			Bind: "FlowName",
		},
	}.Build(),
}
