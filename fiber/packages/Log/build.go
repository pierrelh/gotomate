package log

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a log instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Print":
		return new(PrintDatabinder), PrintTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
