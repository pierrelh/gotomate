package flow

import (
	"fmt"
	"gotomate-astilectron/fiber/template"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.Template) {
	switch function {
	case "Start":
		return new(StartDatabinder), StartTemplate
	case "End":
		return nil, nil
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
