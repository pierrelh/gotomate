package log

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"
)

// Print log a value
func Print(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Logging")

	content, err := variable.GetValue(instructionData, "VarName", "LogIsVar", "Log")
	if err != nil {
		finished <- true
		return -1
	}

	log.Plain("LOG:", content)
	finished <- true
	return -1
}
