package file

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Create":
		return new(CreateDatabinder), CreateTemplate
	case "Delete":
		return new(DeleteDatabinder), DeleteTemplate
	case "Read":
		return new(ReadDatabinder), ReadTemplate
	case "Write":
		return new(WriteDatabinder), WriteTemplate
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
		return nil, nil
	}
}
