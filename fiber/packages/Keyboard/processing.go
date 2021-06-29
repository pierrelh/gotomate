package keyboard

import (
	"fmt"
	"reflect"
)

// Processing process the functions from keyboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Tap":
		go func() {
			nextID = Tap(instructionData, finished)
		}()
		<-finished
	case "Type":
		go func() {
			nextID = Type(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
