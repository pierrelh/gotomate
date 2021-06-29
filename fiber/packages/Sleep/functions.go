package sleep

import (
	"fmt"
	"gotomate/fiber/variable"
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

	fmt.Println("FIBER INFO: Sleeping for: ", duration.(float64), "ms")
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

	fmt.Println("FIBER INFO: Sleeping for: ", duration.(float64), "s")
	time.Sleep(time.Duration(duration.(float64)) * time.Second)
	finished <- true
	return -1
}
