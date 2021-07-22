package define

import (
	"gotomate/fiber/template"
)

// ArrayTemplate Dialog's Define an Array Template
var ArrayTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Name:",
		},
		Input: template.TextInput{
			Bind: "Name",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Values (separated by ',' without spaces):",
		},
		Input: template.TextInput{
			Bind:         "Value",
			BindVariable: "VarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "IsVar",
		},
	}.Build(),
}

// BoolTemplate Dialog's Define Bool Template
var BoolTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Name:",
		},
		Input: template.TextInput{
			Bind: "Name",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value:",
		},
		Input: template.CheckboxInput{
			Bind:         "Value",
			BindVariable: "VarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "IsVar",
		},
	}.Build(),
}

// FloatTemplate Dialog's Define Float Template
var FloatTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Name:",
		},
		Input: template.TextInput{
			Bind: "Name",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value:",
		},
		Input: template.NumberInput{
			Bind:         "Value",
			BindVariable: "VarName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IsVar",
		},
	}.Build(),
}

// IntTemplate Dialog's Define Int Template
var IntTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Name:",
		},
		Input: template.TextInput{
			Bind: "Name",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value:",
		},
		Input: template.NumberInput{
			Bind:         "Value",
			BindVariable: "VarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IsVar",
		},
	}.Build(),
}

// StringTemplate Dialog's Define String Template
var StringTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Name:",
		},
		Input: template.TextInput{
			Bind: "Name",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value:",
		},
		Input: template.TextInput{
			Bind:         "Value",
			BindVariable: "VarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "IsVar",
		},
	}.Build(),
}
