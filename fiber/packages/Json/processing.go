package json

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "CreateJson":
		go func() {
			nextID = CreateJson(instructionData, finished)
		}()
		<-finished
	case "JsonToDictionary":
		go func() {
			nextID = JsonToDictionary(instructionData, finished)
		}()
		<-finished
	case "SaveJson":
		go func() {
			nextID = SaveJson(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
