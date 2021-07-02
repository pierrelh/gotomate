package algorithmic

import (
	"fmt"
	"gotomate-astilectron/fiber/template"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, template.Template) {
	switch function {
	case "For":
		return new(ForDatabinder), ForTemplate
	case "If":
		return new(IfDatabinder), IfTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
