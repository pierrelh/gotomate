package sleep

import (
	"gotomate/fiber/variable"
	"gotomate/log"
	"time"
)

// MilliSleep sleep tm milli second
func MilliSleep(instructionData interface{}, finished chan bool) int {
	duration, err := variable.Keys{VarName: "DurationVarName", IsVarName: "DurationIsVar", Name: "Duration"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}
	log.FiberInfo("Sleeping for: ", duration.(float64), "ms")

	time.Sleep(time.Duration(duration.(float64)) * time.Millisecond)
	finished <- true
	return -1
}

// Sleep time.Sleep tm second
func Sleep(instructionData interface{}, finished chan bool) int {
	duration, err := variable.Keys{VarName: "DurationVarName", IsVarName: "DurationIsVar", Name: "Duration"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	log.FiberInfo("Sleeping for: ", duration.(float64), "s")

	time.Sleep(time.Duration(duration.(float64)) * time.Second)
	finished <- true
	return -1
}
