package screen

import "github.com/lxn/walk/declarative"

// PixelColorTemplate Dialog's PixelColor Template
var PixelColorTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "X",
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
		Title:  "Y",
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
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// MouseColorTemplate Dialog's MouseColor Template
var MouseColorTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// ScreenSizeTemplate Dialog's ScreenSize Template
var ScreenSizeTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Height Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("HeightOutput"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Width Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("WidthOutput"),
		CompactHeight: true,
	},
}

// PartScreenShotTemplate Dialog's PartScreenShot Template
var PartScreenShotTemplate = []declarative.Widget{
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
	declarative.GroupBox{
		Title:  "X",
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
		Title:  "Y",
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
	declarative.GroupBox{
		Title:  "Height",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("HVarName"),
				Visible:       declarative.Bind("HIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("H"),
				Visible:  declarative.Bind("!HIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "HIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("HIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Width",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("WVarName"),
				Visible:       declarative.Bind("WIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("W"),
				Visible:  declarative.Bind("!WIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "WIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("WIsVar"),
			},
		},
	},
}

// ScreenShotTemplate Dialog's ScreenShot Template
var ScreenShotTemplate = []declarative.Widget{
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
}
