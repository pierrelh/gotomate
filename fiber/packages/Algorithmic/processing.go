package algorithmic

import (
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "For":
		go func() {
			nextID = For(instructionData, finished)
		}()
		<-finished
	case "If":
		go func() {
			nextID = If(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}
