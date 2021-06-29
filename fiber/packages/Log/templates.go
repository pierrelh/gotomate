package log

import "github.com/lxn/walk/declarative"

// PrintTemplate Dialog's LogPrint Template
var PrintTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Log",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Log"),
				Visible: declarative.Bind("!IsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("LogIsVar"),
			},
		},
	},
}
