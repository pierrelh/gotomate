package sleep

import (
	"gotomate-astilectron/fiber/template"
)

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = &template.Template{
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
var MilliSleepTemplate = &template.Template{
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
