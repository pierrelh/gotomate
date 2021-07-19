package input

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a input instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Bool":
		return new(InputDatabinder), InputTemplate
	case "Float":
		return new(InputDatabinder), InputTemplate
	case "Int":
		return new(InputDatabinder), InputTemplate
	case "String":
		return new(InputDatabinder), InputTemplate
	}
	return nil, nil
}
