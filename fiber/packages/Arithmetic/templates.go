package arithmetic

import (
	"gotomate-astilectron/fiber/template"
)

// DivideTemplate Dialog's Divide Template
var DivideTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Value One:",
		},
		Input: template.NumberInput{
			Bind:         "ValueOne",
			BindVariable: "VarOneName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueOneIsVar",
		},
	}.Build(),
	template.Field{
		TextView: template.TextView{
			Text:           "÷",
			FontWeight:     template.FW800,
			TextAlignement: template.TACenter,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value Two:",
		},
		Input: template.NumberInput{
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
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

// MultiplyTemplate Dialog's Multiply Template
var MultiplyTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Value One:",
		},
		Input: template.NumberInput{
			Bind:         "ValueOne",
			BindVariable: "VarOneName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueOneIsVar",
		},
	}.Build(),
	template.Field{
		TextView: template.TextView{
			Text:           "X",
			FontWeight:     template.FW800,
			TextAlignement: template.TACenter,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value Two:",
		},
		Input: template.NumberInput{
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
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

// PowTemplate Dialog's Pow Template
var PowTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Value One:",
		},
		Input: template.NumberInput{
			Bind:         "ValueOne",
			BindVariable: "VarOneName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueOneIsVar",
		},
	}.Build(),
	template.Field{
		TextView: template.TextView{
			Text:           "Pow",
			FontWeight:     template.FW800,
			TextAlignement: template.TACenter,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value Two:",
		},
		Input: template.NumberInput{
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
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

// SqrtTemplate Dialog's Sqrt Template
var SqrtTemplate = &template.InstructionTemplate{
	template.Field{
		TextView: template.TextView{
			Text:           "√",
			FontWeight:     template.FW800,
			TextAlignement: template.TALeft,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value One:",
		},
		Input: template.NumberInput{
			Bind:         "Value",
			BindVariable: "VarName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueIsVar",
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

// SubstractTemplate Dialog's Substract Template
var SubstractTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Value One:",
		},
		Input: template.NumberInput{
			Bind:         "ValueOne",
			BindVariable: "VarOneName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueOneIsVar",
		},
	}.Build(),
	template.Field{
		TextView: template.TextView{
			Text:           "-",
			FontWeight:     template.FW800,
			TextAlignement: template.TACenter,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value Two:",
		},
		Input: template.NumberInput{
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
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

// SumTemplate Dialog's Sum Template
var SumTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Value One:",
		},
		Input: template.NumberInput{
			Bind:         "ValueOne",
			BindVariable: "VarOneName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueOneIsVar",
		},
	}.Build(),
	template.Field{
		TextView: template.TextView{
			Text:           "+",
			FontWeight:     template.FW800,
			TextAlignement: template.TACenter,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value Two:",
		},
		Input: template.NumberInput{
			Bind:         "ValueTwo",
			BindVariable: "VarTwoName",
			Decimals:     5,
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueTwoIsVar",
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
