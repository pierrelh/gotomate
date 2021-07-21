package sound

import (
	"gotomate-astilectron/fiber/template"
)

// GetMutedTemplate Dialog's GetMuted Template
var GetMutedTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// GetVolumeTemplate Dialog's GetVolume Template
var GetVolumeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// SetVolumeTemplate Dialog's SetVolume Template
var SetVolumeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Set Volume to:",
		},
		Input: template.NumberInput{
			Bind:         "Volume",
			BindVariable: "VolumeVarName",
			Decimals:     0,
			Suffix:       "%",
			MaxValue:     100,
			MinValue:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "VolumeIsVar",
		},
	}.Build(),
}
