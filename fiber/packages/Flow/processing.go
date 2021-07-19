package flow

import (
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "End":
		go End(finished)
		<-finished
		return -99
	case "Start":
		go Start(finished)
		<-finished
	default:
		return -2
	}
	return nextID
}
