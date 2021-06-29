package clipboard

import "github.com/lxn/walk/declarative"

// WriteTemplate Dialog's ClipboardWrite Template
var WriteTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Clipboard",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Content"),
				Visible: declarative.Bind("!IsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IsVar"),
			},
		},
	},
}

// ReadTemplate Dialog's ClipboardRead Template
var ReadTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}
