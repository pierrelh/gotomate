package input

import (
	"reflect"
)

// Processing process the functions from input's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Bool":
		go func() {
			nextID = Bool(instructionData, finished)
		}()
		<-finished
	case "Float":
		go func() {
			nextID = Float(instructionData, finished)
		}()
		<-finished
	case "Int":
		go func() {
			nextID = Int(instructionData, finished)
		}()
		<-finished
	case "String":
		go func() {
			nextID = String(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}
