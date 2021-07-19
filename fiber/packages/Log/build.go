package log

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a log instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Print":
		return new(PrintDatabinder), PrintTemplate
	}
	return nil, nil
}
