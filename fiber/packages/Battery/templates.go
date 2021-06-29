package battery

import "github.com/lxn/walk/declarative"

// UserBatteryTemplate Dialog's UserBattery Template
var UserBatteryTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// ParametersTemplate Dialog's BatteryParameters Template
var ParametersTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Battery Name",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("BatteryName"),
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
