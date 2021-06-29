package process

import "github.com/lxn/walk/declarative"

// GetPidTemplate Dialog's GetPid Template
var GetPidTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// GetTitleTemplate Dialog's GetTitle Template
var GetTitleTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "PID of process",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PIDVarName"),
				Visible:       declarative.Bind("PIDIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("PID"),
				Visible:  declarative.Bind("!PIDIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "PIDIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PIDIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// KillProcessTemplate Dialog's KillProcess Template
var KillProcessTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "PID of the process",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PIDVarName"),
				Visible:       declarative.Bind("PIDIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("PID"),
				Visible:  declarative.Bind("!PIDIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "PIDIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PIDIsVar"),
			},
		},
	},
}

// StartProcessTemplate Dialog's StartProcess Template
var StartProcessTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("PathIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!PathIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:      "PathIsAVar",
				Text:      "Is a Var",
				Alignment: declarative.AlignHFarVCenter,
				Checked:   declarative.Bind("PathIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "PID Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
