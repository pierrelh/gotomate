package packages

import (

	// DON'T REMOVE ME / New packages inserted here

	algorithmicpack "gotomate-astilectron/fiber/packages/Algorithmic"
	apipack "gotomate-astilectron/fiber/packages/Api"
	arithmeticpack "gotomate-astilectron/fiber/packages/Arithmetic"
	arraypack "gotomate-astilectron/fiber/packages/Array"
	batterypack "gotomate-astilectron/fiber/packages/Battery"
	chronometerpack "gotomate-astilectron/fiber/packages/Chronometer"
	clipboardpack "gotomate-astilectron/fiber/packages/Clipboard"
	conversionpack "gotomate-astilectron/fiber/packages/Conversion"
	definepack "gotomate-astilectron/fiber/packages/Define"
	dictionarypack "gotomate-astilectron/fiber/packages/Dictionary"
	filepack "gotomate-astilectron/fiber/packages/File"
	flowpack "gotomate-astilectron/fiber/packages/Flow"
	inputpack "gotomate-astilectron/fiber/packages/Input"
	jsonpack "gotomate-astilectron/fiber/packages/Json"
	keyboardpack "gotomate-astilectron/fiber/packages/Keyboard"
	logpack "gotomate-astilectron/fiber/packages/Log"
	mousepack "gotomate-astilectron/fiber/packages/Mouse"
	notificationpack "gotomate-astilectron/fiber/packages/Notification"
	processpack "gotomate-astilectron/fiber/packages/Process"
	scrapingpack "gotomate-astilectron/fiber/packages/Scraping"
	screenpack "gotomate-astilectron/fiber/packages/Screen"
	sleeppack "gotomate-astilectron/fiber/packages/Sleep"
	soundpack "gotomate-astilectron/fiber/packages/Sound"
	systimepack "gotomate-astilectron/fiber/packages/Systime"
	"gotomate-astilectron/fiber/template"
	"gotomate-astilectron/log"
)

// PackageDecode Getting the right databinder & the right template needed
func PackageDecode(packageName string, funcName string) (interface{}, *template.InstructionTemplate) {
	var databinder interface{}
	var template *template.InstructionTemplate
	switch packageName {
	case "Flow":
		databinder, template = flowpack.Build(funcName)
	// DON'T REMOVE ME / New Build inserted here
	case "Algorithmic":
		databinder, template = algorithmicpack.Build(funcName)
	case "Api":
		databinder, template = apipack.Build(funcName)
	case "Arithmetic":
		databinder, template = arithmeticpack.Build(funcName)
	case "Array":
		databinder, template = arraypack.Build(funcName)
	case "Battery":
		databinder, template = batterypack.Build(funcName)
	case "Chronometer":
		databinder, template = chronometerpack.Build(funcName)
	case "Clipboard":
		databinder, template = clipboardpack.Build(funcName)
	case "Conversion":
		databinder, template = conversionpack.Build(funcName)
	case "Define":
		databinder, template = definepack.Build(funcName)
	case "Dictionary":
		databinder, template = dictionarypack.Build(funcName)
	case "File":
		databinder, template = filepack.Build(funcName)
	case "Input":
		databinder, template = inputpack.Build(funcName)
	case "Json":
		databinder, template = jsonpack.Build(funcName)
	case "Keyboard":
		databinder, template = keyboardpack.Build(funcName)
	case "Log":
		databinder, template = logpack.Build(funcName)
	case "Mouse":
		databinder, template = mousepack.Build(funcName)
	case "Notification":
		databinder, template = notificationpack.Build(funcName)
	case "Process":
		databinder, template = processpack.Build(funcName)
	case "Scraping":
		databinder, template = scrapingpack.Build(funcName)
	case "Screen":
		databinder, template = screenpack.Build(funcName)
	case "Sleep":
		databinder, template = sleeppack.Build(funcName)
	case "Sound":
		databinder, template = soundpack.Build(funcName)
	case "Systime":
		databinder, template = systimepack.Build(funcName)
	default:
		log.GotomateError("Unable to find the package's dialog", packageName, "for the function", funcName)
		return nil, nil
	}
	if databinder == nil && template == nil && (packageName != "Flow" && packageName != "End") {
		log.GotomateError("Unable to find the function for instruction's building")
	}
	return databinder, template
}
