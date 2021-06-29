package keyboard

import "github.com/lxn/walk/declarative"

// TapTemplate Dialog's KeyboardTap Template
var TapTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Input:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Input"),
		CompactHeight: true,
		MaxLength:     1,
	},
	declarative.Label{
		Text: "Special Input 1:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("Special1", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Inputs(),
	},
	declarative.Label{
		Text: "Special Input 2:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("Special2", declarative.SelRequired{}),
		BindingMember: "Name",
		DisplayMember: "Name",
		Model:         Inputs(),
	},
}

// TypeTemplate Dialog's KeyboardType Template
var TypeTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Input"),
				Visible: declarative.Bind("!IsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
			},
		},
	},
}
