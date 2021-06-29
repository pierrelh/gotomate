package api

import "github.com/lxn/walk/declarative"

// GetTemplate Dialog's Get Template
var GetTemplate = []declarative.Widget{
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
				Name:    "PathIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// PostTemplate Dialog's Post Template
var PostTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Dictionary to Post",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("DataVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
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
				Name:    "PathIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
