package conversion

import (
	"gotomate/fiber/template"
)

// BoolConversionTemplate Dialog's BoolConversion Template
var BoolConversionTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Input:",
		},
		Input: template.CheckboxInput{
			Bind:         "Input",
			BindVariable: "InputVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "InputIsVar",
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

// FloatConversionTemplate Dialog's FloatConversion Template
var FloatConversionTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Input:",
		},
		Input: template.NumberInput{
			Bind:         "Input",
			BindVariable: "InputVarName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "InputIsVar",
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

// IntConversionTemplate Dialog's IntConversion Template
var IntConversionTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Input:",
		},
		Input: template.NumberInput{
			Bind:         "Input",
			BindVariable: "InputVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "InputIsVar",
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

// StringConversionTemplate Dialog's StringConversion Template
var StringConversionTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Input:",
		},
		Input: template.TextInput{
			Bind:         "Input",
			BindVariable: "InputVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "InputIsVar",
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
