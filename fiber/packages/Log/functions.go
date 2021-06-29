package log

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
)

// Print log a value
func Print(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Logging ...")

	log, err := variable.GetValue(instructionData, "VarName", "LogIsVar", "Log")
	if err != nil {
		finished <- true
		return -1
	}

	fmt.Println("LOG: ", log)
	finished <- true
	return -1
}
