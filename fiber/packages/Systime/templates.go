package systime

import "github.com/lxn/walk/declarative"

//SysTimeTemplate Dialog's SysTime Template
var SysTimeTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

//SysClockTemplate Dialog's SysClock Template
var SysClockTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Hours Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("HoursOutput"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Minutes Output:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("MinutesOutput"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "SecondsOutput:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("SecondsOutput"),
		CompactHeight: true,
	},
}
