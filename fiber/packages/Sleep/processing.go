package sleep

import (
	"fmt"
	"reflect"
)

// Processing process the functions from sleep's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "MilliSleep":
		go func() {
			nextID = MilliSleep(instructionData, finished)
		}()
		<-finished
	case "Sleep":
		go func() {
			nextID = Sleep(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
