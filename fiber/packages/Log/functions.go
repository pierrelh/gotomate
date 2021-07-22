package log

import (
	"gotomate/fiber/variable"
	"gotomate/log"
)

// Print log a value
func Print(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Logging")

	content, err := variable.Keys{VarName: "VarName", IsVarName: "LogIsVar", Name: "Log"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	log.Plain("LOG:", content)
	finished <- true
	return -1
}
