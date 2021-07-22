package notification

import (
	"gotomate/fiber/template"
)

// CreateTemplate Dialog's NotificationCreate Template
var CreateTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Title:",
		},
		Input: template.TextInput{
			Bind:         "Title",
			BindVariable: "TitleVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "TitleIsVar",
		},
	}.Build(),
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
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// PushTemplate Dialog's NotificationCreate Template
var PushTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Notification:",
		},
		Input: template.TextInput{
			BindVariable: "NotificationVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}
