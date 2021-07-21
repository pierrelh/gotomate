package systime

import (
	"gotomate-astilectron/fiber/template"
)

//SysTimeTemplate Dialog's SysTime Template
var SysTimeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

//SysClockTemplate Dialog's SysClock Template
var SysClockTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Hours Output:",
		},
		Input: template.TextInput{
			Bind: "HoursOutput",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Minutes Output:",
		},
		Input: template.TextInput{
			Bind: "MinutesOutput",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Seconds Output:",
		},
		Input: template.TextInput{
			Bind: "SecondsOutput",
		},
	}.Build(),
}
