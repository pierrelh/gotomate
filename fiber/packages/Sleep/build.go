package sleep

import (
	"fmt"
	"gotomate-astilectron/fiber/template"
)

// Build Return the right databinder & the right template for a sleep instruction
func Build(function string) (interface{}, template.Template) {
	switch function {
	case "Sleep":
		return new(Databinder), SleepTemplate
	case "MilliSleep":
		return new(Databinder), SleepTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
