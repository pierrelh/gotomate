package process

import (
	"gotomate-astilectron/fiber/template"
)

// GetPidTemplate Dialog's GetPid Template
var GetPidTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// GetTitleTemplate Dialog's GetTitle Template
var GetTitleTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Process's PID:",
		},
		Input: template.NumberInput{
			Bind:         "PID",
			BindVariable: "PIDVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "PIDIsVar",
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

// KillProcessTemplate Dialog's KillProcess Template
var KillProcessTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Process's PID:",
		},
		Input: template.NumberInput{
			Bind:         "PID",
			BindVariable: "PIDVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "PIDIsVar",
		},
	}.Build(),
}

// StartProcessTemplate Dialog's StartProcess Template
var StartProcessTemplate = &template.InstructionTemplate{
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
			Text: "PID Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}
