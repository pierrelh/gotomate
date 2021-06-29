package chronometer

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"time"
)

// End a chronometer & save the elapsed time
func End(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting elapsed chronometer time ...")

	chrono, err := variable.GetValue(instructionData, "ChronoVarName")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Since(chrono.(time.Time)))
	finished <- true
	return -1
}

// Start a chronometer
func Start(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting a new chronometer ...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Now())
	finished <- true
	return -1
}
