package screen

import (
	"gotomate-astilectron/fiber/template"
)

// PixelColorTemplate Dialog's PixelColor Template
var PixelColorTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Horizontal:",
		},
		Input: template.NumberInput{
			Bind:         "X",
			BindVariable: "XVarName",
			Decimals:     0,
			Suffix:       "px",
		},
		VariableToggler: template.VariableToggler{
			Bind: "XIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Vertical:",
		},
		Input: template.NumberInput{
			Bind:         "Y",
			BindVariable: "YVarName",
			Decimals:     0,
			Suffix:       "px",
		},
		VariableToggler: template.VariableToggler{
			Bind: "YIsVar",
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

// MouseColorTemplate Dialog's MouseColor Template
var MouseColorTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// ScreenSizeTemplate Dialog's ScreenSize Template
var ScreenSizeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Height Output:",
		},
		Input: template.TextInput{
			Bind: "HeightOutput",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Width Output:",
		},
		Input: template.TextInput{
			Bind: "WidthOutput",
		},
	}.Build(),
}

// PartScreenShotTemplate Dialog's PartScreenShot Template
var PartScreenShotTemplate = &template.InstructionTemplate{
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
			Text: "Horizontal:",
		},
		Input: template.NumberInput{
			Bind:         "X",
			BindVariable: "XVarName",
			Decimals:     0,
			Suffix:       "px",
		},
		VariableToggler: template.VariableToggler{
			Bind: "XIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Vertical:",
		},
		Input: template.NumberInput{
			Bind:         "Y",
			BindVariable: "YVarName",
			Decimals:     0,
			Suffix:       "px",
		},
		VariableToggler: template.VariableToggler{
			Bind: "YIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Height:",
		},
		Input: template.NumberInput{
			Bind:         "H",
			BindVariable: "HVarName",
			Decimals:     0,
			Suffix:       "px",
		},
		VariableToggler: template.VariableToggler{
			Bind: "HIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Width:",
		},
		Input: template.NumberInput{
			Bind:         "W",
			BindVariable: "WVarName",
			Decimals:     0,
			Suffix:       "px",
		},
		VariableToggler: template.VariableToggler{
			Bind: "WIsVar",
		},
	}.Build(),
}

// ScreenShotTemplate Dialog's ScreenShot Template
var ScreenShotTemplate = &template.InstructionTemplate{
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
}
