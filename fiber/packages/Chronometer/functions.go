package chronometer

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"
	"time"
)

// End a chronometer & save the elapsed time
func End(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Getting elapsed chronometer time")

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
	log.FiberInfo("Starting a new chronometer")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Now())
	finished <- true
	return -1
}
