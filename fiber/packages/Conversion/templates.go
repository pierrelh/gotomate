package conversion

import "github.com/lxn/walk/declarative"

// BoolConversionTemplate Dialog's BoolConversion Template
var BoolConversionTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:               declarative.Bind("InputVarName"),
				Visible:            declarative.Bind("IsAVar.Checked"),
				AlwaysConsumeSpace: true,
				CompactHeight:      true,
			},
			declarative.CheckBox{
				Text:    "True ?",
				Visible: declarative.Bind("!IsAVar.Checked"),
				Checked: declarative.Bind("Input"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
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

// FloatConversionTemplate Dialog's FloatConversion Template
var FloatConversionTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("InputVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Input"),
				Visible:  declarative.Bind("!IsAVar.Checked"),
				Decimals: 5,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
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

// IntConversionTemplate Dialog's IntConversion Template
var IntConversionTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("InputVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Input"),
				Visible:  declarative.Bind("!IsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
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

// StringConversionTemplate Dialog's StringConversion Template
var StringConversionTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Input to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("InputVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("Input"),
				Visible:       declarative.Bind("!IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("InputIsVar"),
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
