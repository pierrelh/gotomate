package algorithmic

import (
	"gotomate-astilectron/fiber/template"
)

// Comparators Define the possibles values of algorithmic comparators
func Comparators() []template.Option {
	return []template.Option{
		{Name: "", Value: ""},
		{Name: "==", Value: "=="},
		{Name: "!=", Value: "!="},
		{Name: ">", Value: ">"},
		{Name: ">=", Value: ">="},
		{Name: "<", Value: "<"},
		{Name: "<=", Value: "<="},
	}
}

// ForTemplate Dialog's For Template
var ForTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Variable to increment",
		},
		Input: template.TextInput{
			BindVariable: "ValueOneVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Comparator",
		},
		Input: template.Select{
			Bind:    "Comparator",
			Options: Comparators(),
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value to compare",
		},
		Input: template.NumberInput{
			Bind:         "ValueTwo",
			BindVariable: "ValueTwoVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Increment",
		},
		Input: template.NumberInput{
			Bind:         "Increment",
			BindVariable: "IncrementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "IncrementIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Else instruction ID",
		},
		Input: template.NumberInput{
			Bind:         "FalseInstruction",
			BindVariable: "FalseInstructionVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "FalseInstructionIsVar",
		},
	}.Build(),
}

// IfTemplate Dialog's If Template
var IfTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Value 1",
		},
		Input: template.TextInput{
			Bind:         "ValueOne",
			BindVariable: "ValueOneVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueOneIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Comparator",
		},
		Input: template.Select{
			Bind:    "Comparator",
			Options: Comparators(),
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value 2",
		},
		Input: template.TextInput{
			Bind:         "ValueTwo",
			BindVariable: "ValueTwoVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Else instruction ID",
		},
		Input: template.NumberInput{
			Bind:         "FalseInstruction",
			BindVariable: "FalseInstructionVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "FalseInstructionIsVar",
		},
	}.Build(),
}
