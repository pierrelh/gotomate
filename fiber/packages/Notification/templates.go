package notification

import "github.com/lxn/walk/declarative"

// CreateTemplate Dialog's NotificationCreate Template
var CreateTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Title",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("TitleVarName"),
				Visible:       declarative.Bind("TitleIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("Title"),
				Visible:       declarative.Bind("!TitleIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "TitleIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("TitleIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Message",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("MessageVarName"),
				Visible:       declarative.Bind("MessageIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Message"),
				Visible: declarative.Bind("!MessageIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "MessageIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("MessageIsVar"),
			},
		},
	},
}
