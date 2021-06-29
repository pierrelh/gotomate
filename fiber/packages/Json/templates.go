package json

import (
	"github.com/lxn/walk/declarative"
)

// CreateJsonTemplate Dialog's Define an CreateJson Template
var CreateJsonTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Path to import",
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
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// JsonToDictionaryTemplate Dialog's Define an JsonToDictionary Template
var JsonToDictionaryTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Json to convert",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("JsonVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
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

// SaveJsonTemplate Dialog's Define SaveJson Template
var SaveJsonTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Json to save",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("JsonVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Path to save",
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
