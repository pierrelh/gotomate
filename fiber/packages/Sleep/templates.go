package sleep

import "github.com/lxn/walk/declarative"

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Sleep For",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("DurationVarName"),
				Visible:       declarative.Bind("DurationIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Duration"),
				Visible:  declarative.Bind("!DurationIsAVar.Checked"),
				Suffix:   " s",
				Decimals: 2,
			},
			declarative.CheckBox{
				Name:    "DurationIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("DurationIsVar"),
			},
		},
	},
}

// MilliSleepTemplate Dialog's MilliSleep Template
var MilliSleepTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Sleep For",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("DurationVarName"),
				Visible:       declarative.Bind("DurationIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Duration"),
				Visible:  declarative.Bind("!DurationIsAVar.Checked"),
				Suffix:   " ms",
				Decimals: 2,
			},
			declarative.CheckBox{
				Name:    "DurationIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("DurationIsVar"),
			},
		},
	},
}
