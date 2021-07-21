package mouse

import (
	"gotomate-astilectron/fiber/template"
)

// Buttons Define the possibles values of MouseButton
func Buttons() []template.Option {
	return []template.Option{
		{Name: "left", Value: "Left"},
		{Name: "right", Value: "Right"},
	}
}

// ClickTemplate Dialog's MouseClic Template
var ClickTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Click:",
		},
		Input: template.Select{
			Bind:    "MouseButtonName",
			Options: Buttons(),
		},
	}.Build(),
}

// DragTemplate Dialog's Drag Template
var DragTemplate = &template.InstructionTemplate{
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
}

// MoveTemplate Dialog's MouseMove Template
var MoveTemplate = &template.InstructionTemplate{
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
}

// ScrollTemplate Dialog's MouseScroll Template
var ScrollTemplate = &template.InstructionTemplate{
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
}
