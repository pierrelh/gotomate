package file

import "github.com/lxn/walk/declarative"

// CreateTemplate Dialog's Create Template
var CreateTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!IsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
}

// DeleteTemplate Dialog's Delete Template
var DeleteTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!IsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
}

// ReadTemplate Dialog's Read Template
var ReadTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!IsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsAVar",
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

// WriteTemplate Dialog's Write Template
var WriteTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("IsPathAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!IsPathAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsPathAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Content",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ContentVarName"),
				Visible:       declarative.Bind("IsContentAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Content"),
				Visible: declarative.Bind("!IsContentAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "IsContentAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ContentIsVar"),
			},
		},
	},
}
