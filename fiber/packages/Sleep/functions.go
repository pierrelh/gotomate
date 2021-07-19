package sleep

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"
	"time"
)

// MilliSleep sleep tm milli second
func MilliSleep(instructionData reflect.Value, finished chan bool) int {

	duration, err := variable.GetValue(instructionData, "DurationVarName", "DurationIsVar", "Duration")
	if err != nil {
		finished <- true
		return -1
	}
	log.FiberInfo("Sleeping for: " + duration.(string) + "ms")

	time.Sleep(time.Duration(duration.(float64)) * time.Millisecond)
	finished <- true
	return -1
}

// Sleep time.Sleep tm second
func Sleep(instructionData reflect.Value, finished chan bool) int {
	duration, err := variable.GetValue(instructionData, "DurationVarName", "DurationIsVar", "Duration")
	if err != nil {
		finished <- true
		return -1
	}

	log.FiberInfo("Sleeping for: " + duration.(string) + "s")

	time.Sleep(time.Duration(duration.(float64)) * time.Second)
	finished <- true
	return -1
}
