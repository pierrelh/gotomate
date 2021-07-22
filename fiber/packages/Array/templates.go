package array

import (
	"gotomate/fiber/template"
)

// GetArrayLengthTemplate Dialog's GetArrayLength Template
var GetArrayLengthTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
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

// GetValueTemplate Dialog's GetValue Template
var GetValueTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Index to get:",
		},
		Input: template.NumberInput{
			Bind:         "Index",
			BindVariable: "IndexVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IndexIsVar",
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

// PopAtTemplate Dialog's PopAt Template
var PopAtTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Index to pop:",
		},
		Input: template.NumberInput{
			Bind:         "Index",
			BindVariable: "IndexVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IndexIsVar",
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

// PopLastTemplate Dialog's PopLast Template
var PopLastTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
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

// PushAtTemplate Dialog's PushAt Template
var PushAtTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Index to push at:",
		},
		Input: template.NumberInput{
			Bind:         "Index",
			BindVariable: "IndexVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IndexIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value to push:",
		},
		Input: template.TextInput{
			BindVariable: "ValueVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// PushLastTemplate Dialog's PushLast Template
var PushLastTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value to push:",
		},
		Input: template.TextInput{
			BindVariable: "ValueVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// RemoveAtTemplate Dialog's RemoveAt Template
var RemoveAtTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Index to remove:",
		},
		Input: template.NumberInput{
			Bind:         "Index",
			BindVariable: "IndexVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IndexIsVar",
		},
	}.Build(),
}

// RemoveLastTemplate Dialog's RemoveLast Template
var RemoveLastTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// ShuffleTemplate Dialog's Shuffle Template
var ShuffleTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// UpdateValueTemplate Dialog's UpdateValue Template
var UpdateValueTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Array:",
		},
		Input: template.TextInput{
			BindVariable: "ArrayVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Index to update:",
		},
		Input: template.NumberInput{
			Bind:         "Index",
			BindVariable: "IndexVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "IndexIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "New value:",
		},
		Input: template.TextInput{
			BindVariable: "ValueVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}
