package packages

import (

	// DON'T REMOVE ME / New packages inserted here

	"fmt"
	algorithmic "gotomate-astilectron/fiber/packages/Algorithmic"
	flow "gotomate-astilectron/fiber/packages/Flow"
	sleep "gotomate-astilectron/fiber/packages/Sleep"
	"gotomate-astilectron/fiber/template"
)

// PackageDecode Getting the right databinder & the right template needed
func PackageDecode(packageName string, funcName string) (interface{}, template.Template) {
	switch packageName {
	case "Flow":
		return flow.Build(funcName)
	// DON'T REMOVE ME / New Build inserted here
	case "Algorithmic":
		return algorithmic.Build(funcName)
	case "Sleep":
		return sleep.Build(funcName)
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the dialog's package")
		return nil, nil
	}
}
