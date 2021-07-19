package systime

import (
	"reflect"
)

// Processing process the functions from systime's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "GetCurrentSysClock":
		go func() {
			nextID = GetCurrentSysClock(instructionData, finished)
		}()
		<-finished
	case "GetCurrentSysTime":
		go func() {
			nextID = GetCurrentSysTime(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}
