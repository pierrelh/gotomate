package flow

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) bool {
	switch funcName {
	case "End":
		go End(finished)
		<-finished
		return true
	case "Start":
		go Start(finished)
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return false
}
