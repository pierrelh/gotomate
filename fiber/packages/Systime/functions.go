package systime

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"time"
)

// GetCurrentSysClock return the current sys clock
func GetCurrentSysClock(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting sys clock...")

	h, m, s := time.Now().Clock()
	variable.SetVariable(instructionData.FieldByName("HoursOutput").Interface().(string), h)
	variable.SetVariable(instructionData.FieldByName("MinutesOutput").Interface().(string), m)
	variable.SetVariable(instructionData.FieldByName("SecondsOutput").Interface().(string), s)
	finished <- true
	return -1
}

// GetCurrentSysTime return the current sys time
func GetCurrentSysTime(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting sys time...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Now())
	finished <- true
	return -1
}
