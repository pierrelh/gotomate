package input

import "github.com/lxn/walk/declarative"

// InputTemplate Dialog's Input Template
var InputTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Message",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("MessageVarName"),
				Visible:       declarative.Bind("IsMessageAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Message"),
				Visible: declarative.Bind("!IsMessageAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsMessageAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("MessageIsVar"),
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
