package keyboard

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a keyboard instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Tap":
		return new(TapDatabinder), TapTemplate
	case "Type":
		return new(TypeDatabinder), TypeTemplate
	}
	return nil, nil
}
