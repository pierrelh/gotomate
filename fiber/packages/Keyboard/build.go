package keyboard

import (
	"fmt"

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
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
