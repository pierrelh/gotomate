package chronometer

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "End":
		return new(EndDatabinder), EndTemplate
	case "Start":
		return new(StartDatabinder), StartTemplate
	default:
		return nil, nil
	}
}
