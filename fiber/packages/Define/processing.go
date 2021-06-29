package define

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "ArrayOfBool":
		go func() {
			nextID = ArrayOfBool(instructionData, finished)
		}()
		<-finished
	case "ArrayOfFloat":
		go func() {
			nextID = ArrayOfFloat(instructionData, finished)
		}()
		<-finished
	case "ArrayOfInt":
		go func() {
			nextID = ArrayOfInt(instructionData, finished)
		}()
		<-finished
	case "ArrayOfString":
		go func() {
			nextID = ArrayOfString(instructionData, finished)
		}()
		<-finished
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
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
