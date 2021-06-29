package conversion

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "BoolToFloat":
		go func() {
			nextID = BoolToFloat(instructionData, finished)
		}()
		<-finished
	case "BoolToInt":
		go func() {
			nextID = BoolToInt(instructionData, finished)
		}()
		<-finished
	case "BoolToString":
		go func() {
			nextID = BoolToString(instructionData, finished)
		}()
		<-finished
	case "FloatToInt":
		go func() {
			nextID = FloatToInt(instructionData, finished)
		}()
		<-finished
	case "FloatToString":
		go func() {
			nextID = FloatToString(instructionData, finished)
		}()
		<-finished
	case "IntToFloat":
		go func() {
			nextID = IntToFloat(instructionData, finished)
		}()
		<-finished
	case "IntToString":
		go func() {
			nextID = IntToString(instructionData, finished)
		}()
		<-finished
	case "StringToBool":
		go func() {
			nextID = StringToBool(instructionData, finished)
		}()
		<-finished
	case "StringToFloat":
		go func() {
			nextID = StringToFloat(instructionData, finished)
		}()
		<-finished
	case "StringToInt":
		go func() {
			nextID = StringToInt(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
