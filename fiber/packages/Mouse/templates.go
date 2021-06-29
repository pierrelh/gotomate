package mouse

import "github.com/lxn/walk/declarative"

// ClickTemplate Dialog's MouseClic Template
var ClickTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Click:",
	},
	declarative.ComboBox{
		Value:         declarative.Bind("MouseButtonName", declarative.SelRequired{}),
		BindingMember: "ID",
		DisplayMember: "Name",
		Model:         Buttons(),
	},
}

// DragTemplate Dialog's Drag Template
var DragTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Horizontal",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("XVarName"),
				Visible:       declarative.Bind("XIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Visible:  declarative.Bind("!XIsAVar.Checked"),
				Value:    declarative.Bind("X"),
				Decimals: 0,
				Suffix:   " px",
			},
			declarative.CheckBox{
				Name:    "XIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("XIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Vertical",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("YVarName"),
				Visible:       declarative.Bind("YIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Y"),
				Visible:  declarative.Bind("!YIsAVar.Checked"),
				Decimals: 0,
				Suffix:   " px",
			},
			declarative.CheckBox{
				Name:    "YIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("YIsVar"),
			},
		},
	},
}

// MoveTemplate Dialog's MouseMove Template
var MoveTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Horizontal",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("XVarName"),
				Visible:       declarative.Bind("XIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("X"),
				Visible:  declarative.Bind("!XIsAVar.Checked"),
				Suffix:   " px",
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "XIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("XIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Vertical",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("YVarName"),
				Visible:       declarative.Bind("YIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Y"),
				Visible:  declarative.Bind("!YIsAVar.Checked"),
				Suffix:   " px",
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "YIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("YIsVar"),
			},
		},
	},
}

// ScrollTemplate Dialog's MouseScroll Template
var ScrollTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Horizontal",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("XVarName"),
				Visible:       declarative.Bind("XIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Visible:  declarative.Bind("!XIsAVar.Checked"),
				Value:    declarative.Bind("X"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "XIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("XIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Vertical",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("YVarName"),
				Visible:       declarative.Bind("YIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Y"),
				Visible:  declarative.Bind("!YIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "YIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("YIsVar"),
			},
		},
	},
}
