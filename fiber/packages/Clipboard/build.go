package clipboard

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a clipboard instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Write":
		return new(WriteDatabinder), WriteTemplate
	case "Read":
		return new(ReadDatabinder), ReadTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
