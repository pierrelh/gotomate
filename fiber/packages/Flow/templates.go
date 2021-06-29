package flow

import (
	"gotomate-astilectron/fiber/template"
)

// StartTemplate Dialog's FlowStart Template
var StartTemplate = []template.Field{
	{
		Label: template.Label{
			Text:     "Flow Name",
			AssignTo: "FlowNameInput",
		}.NewElement(),
		Input: template.Input{
			Type: template.Text,
			Bind: "FlowName",
			ID:   "FlowNameInput",
		}.NewElement(),
	},
}
