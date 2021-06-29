package chronometer

import "github.com/lxn/walk/declarative"

// EndTemplate Dialog's End Template
var EndTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Chronometer value",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ChronoVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
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

// StartTemplate Dialog's Start Template
var StartTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
