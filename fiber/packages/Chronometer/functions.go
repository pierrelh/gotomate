package chronometer

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"time"
)

// End a chronometer & save the elapsed time
func End(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting elapsed chronometer time")

	chrono, err := variable.Keys{VarName: "ChronoVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", time.Since(chrono.(time.Time)))
	finished <- true
	return -1
}

// Start a chronometer
func Start(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Starting a new chronometer")

	variable.SetVariable(instructionData, "Output", time.Now())
	finished <- true
	return -1
}
