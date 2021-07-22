package sleep

import (
	"gotomate/fiber/template"
)

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Sleep For:",
		},
		Input: template.NumberInput{
			Bind:         "Duration",
			BindVariable: "DurationVarName",
			Suffix:       "s",
		},
		VariableToggler: template.VariableToggler{
			Bind: "DurationIsVar",
		},
	}.Build(),
}

// MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Sleep For:",
		},
		Input: template.NumberInput{
			Bind:         "Duration",
			BindVariable: "DurationVarName",
			Suffix:       "ms",
		},
		VariableToggler: template.VariableToggler{
			Bind: "DurationIsVar",
		},
	}.Build(),
}
