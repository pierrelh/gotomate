package packages

import (

	// DON'T REMOVE ME / New packages inserted here

	algorithmicpack "gotomate/fiber/packages/Algorithmic"
	apipack "gotomate/fiber/packages/Api"
	arithmeticpack "gotomate/fiber/packages/Arithmetic"
	arraypack "gotomate/fiber/packages/Array"
	batterypack "gotomate/fiber/packages/Battery"
	chronometerpack "gotomate/fiber/packages/Chronometer"
	clipboardpack "gotomate/fiber/packages/Clipboard"
	conversionpack "gotomate/fiber/packages/Conversion"
	definepack "gotomate/fiber/packages/Define"
	dictionarypack "gotomate/fiber/packages/Dictionary"
	filepack "gotomate/fiber/packages/File"
	flowpack "gotomate/fiber/packages/Flow"
	inputpack "gotomate/fiber/packages/Input"
	jsonpack "gotomate/fiber/packages/Json"
	keyboardpack "gotomate/fiber/packages/Keyboard"
	logpack "gotomate/fiber/packages/Log"
	mousepack "gotomate/fiber/packages/Mouse"
	notificationpack "gotomate/fiber/packages/Notification"
	processpack "gotomate/fiber/packages/Process"
	scrapingpack "gotomate/fiber/packages/Scraping"
	screenpack "gotomate/fiber/packages/Screen"
	sleeppack "gotomate/fiber/packages/Sleep"
	soundpack "gotomate/fiber/packages/Sound"
	systimepack "gotomate/fiber/packages/Systime"
	"gotomate/fiber/template"
	"gotomate/log"
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
