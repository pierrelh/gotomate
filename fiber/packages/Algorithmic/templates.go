package algorithmic

import (
	"gotomate-astilectron/fiber/template"
)

// Comparators Define the possibles values of algorithmic comparators
func Comparators() []template.Option {
	return []template.Option{
		template.Option{Name: "", Value: ""}.NewElement(),
		template.Option{Name: "==", Value: "=="}.NewElement(),
		template.Option{Name: "!=", Value: "!="}.NewElement(),
		template.Option{Name: ">", Value: ">"}.NewElement(),
		template.Option{Name: ">=", Value: ">="}.NewElement(),
		template.Option{Name: "<", Value: "<"}.NewElement(),
		template.Option{Name: "<=", Value: "<="}.NewElement(),
	}
}

// ForTemplate Dialog's For Template
var ForTemplate = template.Template{
	template.Field{
		Label: template.Label{
			Text:     "Variable to increment",
			AssignTo: "Var",
		}.NewElement(),
		Input: template.Input{
			Type:         template.Text,
			BindVariable: "VarOneName",
			ID:           "Var",
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "Var",
			Checked:  true,
			Disabled: true,
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Comparator",
			AssignTo: "ComparatorSelect",
		}.NewElement(),
		Input: template.Select{
			Bind:    "Comparator",
			ID:      "ComparatorSelect",
			Options: Comparators(),
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Value to compare",
			AssignTo: "Compare",
		}.NewElement(),
		Input: template.Input{
			Type:         template.Number,
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			ID:           "Compare",
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "Compare",
			Bind:     "TwoIsVar",
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Increment",
			AssignTo: "Increment",
		}.NewElement(),
		Input: template.Input{
			Type:         template.Number,
			Bind:         "Increment",
			BindVariable: "IncrementVarName",
			ID:           "Increment",
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "Increment",
			Bind:     "IncrementIsVar",
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Else instruction ID",
			AssignTo: "FalseInstruction",
		}.NewElement(),
		Input: template.Input{
			Type:         template.Number,
			Bind:         "FalseInstruction",
			BindVariable: "FalseInstructionVarName",
			ID:           "FalseInstruction",
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "FalseInstruction",
			Bind:     "FalseInstructionIsVar",
		}.NewElement(),
	}.Build(),
}

// IfTemplate Dialog's If Template
var IfTemplate = template.Template{
	template.Field{
		Label: template.Label{
			Text:     "Value 1",
			AssignTo: "ValueOneID",
		}.NewElement(),
		Input: template.Input{
			Bind:         "ValueOne",
			BindVariable: "VarOneName",
			ID:           "ValueOneID",
			Type:         template.Text,
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "ValueOneID",
			Bind:     "OneIsVar",
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Comparator",
			AssignTo: "ComparatorSelect",
		}.NewElement(),
		Input: template.Select{
			Bind:    "Comparator",
			ID:      "ComparatorSelect",
			Options: Comparators(),
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Value 2",
			AssignTo: "ValueTwoID",
		}.NewElement(),
		Input: template.Input{
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			ID:           "ValueTwoID",
			Type:         template.Text,
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "ValueTwoID",
			Bind:     "TwoIsVar",
		}.NewElement(),
	}.Build(),
	template.Field{
		Label: template.Label{
			Text:     "Else instruction ID",
			AssignTo: "FalseInstruction",
		}.NewElement(),
		Input: template.Input{
			Type:         template.Number,
			Bind:         "FalseInstruction",
			BindVariable: "FalseInstructionVarName",
			ID:           "FalseInstruction",
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "FalseInstruction",
			Bind:     "FalseInstructionIsVar",
		}.NewElement(),
	}.Build(),
}
