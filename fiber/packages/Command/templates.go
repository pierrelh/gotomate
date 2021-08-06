package command

import (
	"gotomate/fiber/template"
)

// DirTemplate Dialog's Dir Template
var DirTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Command:",
		},
		Input: template.TextInput{
			BindVariable: "CommandVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Directory:",
		},
		Input: template.TextInput{
			Bind:         "Dir",
			BindVariable: "DirVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "DirIsVar",
		},
	}.Build(),
}

// GetOutputTemplate Dialog's GetOutput Template
var GetOutputTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Command:",
		},
		Input: template.TextInput{
			BindVariable: "CommandVarName",
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

// NewTemplate Dialog's New Template
var NewTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Program:",
		},
		Input: template.TextInput{
			Bind:         "Program",
			BindVariable: "ProgramVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ProgramIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Parameter:",
		},
		Input: template.TextInput{
			Bind:         "Parameter",
			BindVariable: "ParameterVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ParameterIsVar",
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

// RunTemplate Dialog's Run Template
var RunTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Command:",
		},
		Input: template.TextInput{
			BindVariable: "CommandVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// StartTemplate Dialog's Start Template
var StartTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Command:",
		},
		Input: template.TextInput{
			BindVariable: "CommandVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}
