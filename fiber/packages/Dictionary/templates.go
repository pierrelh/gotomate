package dictionary

import (
	"gotomate/fiber/template"
)

// CreateDictionaryTemplate Dialog's Define an CreateDictionary Template
var CreateDictionaryTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// CreateEntryTemplate Dialog's Define CreateEntry Template
var CreateEntryTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Dictionary:",
		},
		Input: template.TextInput{
			BindVariable: "DictVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Key:",
		},
		Input: template.TextInput{
			Bind:         "Key",
			BindVariable: "KeyVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "KeyIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Value:",
		},
		Input: template.TextInput{
			Bind:         "Value",
			BindVariable: "ValueVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ValueIsVar",
		},
	}.Build(),
}

// DictionaryToJsonTemplate Dialog's Define DictionaryToJson Template
var DictionaryToJsonTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Dictionary:",
		},
		Input: template.TextInput{
			BindVariable: "DictVarName",
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

// RemoveEntryTemplate Dialog's Define RemoveEntry Template
var RemoveEntryTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Dictionary:",
		},
		Input: template.TextInput{
			BindVariable: "DictVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Key:",
		},
		Input: template.TextInput{
			Bind:         "Key",
			BindVariable: "KeyVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "KeyIsVar",
		},
	}.Build(),
}
