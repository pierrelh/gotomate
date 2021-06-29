package define

import (
	"github.com/lxn/walk/declarative"
)

// ArrayTemplate Dialog's Define an Array Template
var ArrayTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.GroupBox{
		Title:  "Values (separated by ',' without spaces)",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Value"),
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

// BoolTemplate Dialog's Define Bool Template
var BoolTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.GroupBox{
		Title:  "Value",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:               declarative.Bind("VarName"),
				Visible:            declarative.Bind("IsAVar.Checked"),
				AlwaysConsumeSpace: true,
				CompactHeight:      true,
			},
			declarative.CheckBox{
				Text:    "True ?",
				Visible: declarative.Bind("!IsAVar.Checked"),
				Checked: declarative.Bind("Value"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IsVar"),
			},
		},
	},
}

// FloatTemplate Dialog's Define Float Template
var FloatTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.GroupBox{
		Title:  "Value",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Value"),
				Visible:  declarative.Bind("!IsAVar.Checked"),
				Decimals: 5,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IsVar"),
			},
		},
	},
}

// IntTemplate Dialog's Define Int Template
var IntTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.GroupBox{
		Title:  "Value",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Value"),
				Visible:  declarative.Bind("!IsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IsVar"),
			},
		},
	},
}

// StringTemplate Dialog's Define String Template
var StringTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.GroupBox{
		Title:  "Value",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Value"),
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
