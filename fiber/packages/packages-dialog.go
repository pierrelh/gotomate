package packages

import (

	// DON'T REMOVE ME / New packages inserted here

	algorithmic "gotomate-astilectron/fiber/packages/Algorithmic"
	flow "gotomate-astilectron/fiber/packages/Flow"
	input "gotomate-astilectron/fiber/packages/Input"
	sleep "gotomate-astilectron/fiber/packages/Sleep"
	"gotomate-astilectron/fiber/template"
	"gotomate-astilectron/log"
)

// PackageDecode Getting the right databinder & the right template needed
func PackageDecode(packageName string, funcName string) (interface{}, *template.Template) {
	var databinder interface{}
	var template *template.Template
	switch packageName {
	case "Flow":
		databinder, template = flow.Build(funcName)
	// DON'T REMOVE ME / New Build inserted here
	case "Algorithmic":
		databinder, template = algorithmic.Build(funcName)
	case "Input":
		databinder, template = input.Build(funcName)
	case "Sleep":
		databinder, template = sleep.Build(funcName)
	default:
		log.GotomateError("Unable to find the package's dialog", packageName, "for the function", funcName)
		return nil, nil
	}
	if databinder == nil && template == nil && (packageName != "Flow" && packageName != "End") {
		log.GotomateError("Unable to find the function for instruction's building")
	}
	return databinder, template
}
