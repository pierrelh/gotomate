package dictionary

import (
	"github.com/lxn/walk/declarative"
)

// CreateDictionaryTemplate Dialog's Define an CreateDictionary Template
var CreateDictionaryTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// CreateEntryTemplate Dialog's Define CreateEntry Template
var CreateEntryTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Dictionary",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("DictVarName"),
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
		Title:  "Key",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("KeyVarName"),
				Visible:       declarative.Bind("KeyIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("Key"),
				Visible:       declarative.Bind("!KeyIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "KeyIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("KeyIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Value",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ValueVarName"),
				Visible:       declarative.Bind("ValueIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Value"),
				Visible: declarative.Bind("!ValueIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ValueIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ValueIsVar"),
			},
		},
	},
}

// DictionaryToJsonTemplate Dialog's Define DictionaryToJson Template
var DictionaryToJsonTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Dictionary",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("DictVarName"),
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

// RemoveEntryTemplate Dialog's Define RemoveEntry Template
var RemoveEntryTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Dictionary",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("DictVarName"),
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
		Title:  "Key",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("KeyVarName"),
				Visible:       declarative.Bind("KeyIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:          declarative.Bind("Key"),
				Visible:       declarative.Bind("!KeyIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Name:    "KeyIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("KeyIsVar"),
			},
		},
	},
}
