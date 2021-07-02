package sleep

import (
	"gotomate-astilectron/fiber/template"
)

// SleepTemplate Dialog's Sleep Template
var SleepTemplate = template.Template{
	template.Field{
		Label: template.Label{
			Text:     "Sleep For (s):",
			AssignTo: "Sleep",
		}.NewElement(),
		Input: template.Input{
			ID:           "Sleep",
			Type:         template.Number,
			Bind:         "Duration",
			BindVariable: "DurationVarName",
		}.NewElement(),
		VariableToggler: template.VariableToggler{
			AssignTo: "Sleep",
			Bind:     "DurationIsVar",
		}.NewElement(),
	}.Build(),
}

// MilliSleepTemplate Dialog's MilliSleep Template
// var MilliSleepTemplate = []declarative.Widget{
// 	declarative.GroupBox{
// 		Title:  "Sleep For",
// 		Layout: declarative.HBox{},
// 		Children: []declarative.Widget{
// 			declarative.TextEdit{
// 				Text:          declarative.Bind("DurationVarName"),
// 				Visible:       declarative.Bind("DurationIsAVar.Checked"),
// 				CompactHeight: true,
// 			},
// 			declarative.NumberEdit{
// 				Value:    declarative.Bind("Duration"),
// 				Visible:  declarative.Bind("!DurationIsAVar.Checked"),
// 				Suffix:   " ms",
// 				Decimals: 2,
// 			},
// 			declarative.CheckBox{
// 				Name:    "DurationIsAVar",
// 				Text:    "Is a Var",
// 				Checked: declarative.Bind("DurationIsVar"),
// 			},
// 		},
// 	},
// }
