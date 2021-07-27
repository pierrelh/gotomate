package battery

import (
	"gotomate/fiber/template"
)

// UserBatteryTemplate Dialog's UserBattery Template
var UserBatteryTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// ParametersTemplate Dialog's BatteryParameters Template
var ParametersTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Battery Name:",
		},
		Input: template.TextInput{
			BindVariable: "BatteryName",
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
			Bind: "Output",
		},
	}.Build(),
}
