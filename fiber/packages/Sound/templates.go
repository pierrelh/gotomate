package sound

import "github.com/lxn/walk/declarative"

// GetMutedTemplate Dialog's GetMuted Template
var GetMutedTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// GetVolumeTemplate Dialog's GetVolume Template
var GetVolumeTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// SetVolumeTemplate Dialog's SetVolume Template
var SetVolumeTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Set Volume to",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("VolumeVarName"),
				Visible:       declarative.Bind("IsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("Volume"),
				Visible:  declarative.Bind("!IsAVar.Checked"),
				Decimals: 0,
				MinValue: 0,
				MaxValue: 100,
				Suffix:   " %",
			},
			declarative.CheckBox{
				Name:    "IsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("VolumeIsVar"),
			},
		},
	},
}
