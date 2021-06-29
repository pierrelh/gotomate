package dictionary

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "CreateDictionary":
		go func() {
			nextID = CreateDictionary(instructionData, finished)
		}()
		<-finished
	case "CreateEntry":
		go func() {
			nextID = CreateEntry(instructionData, finished)
		}()
		<-finished
	case "DictionaryToJson":
		go func() {
			nextID = DictionaryToJson(instructionData, finished)
		}()
		<-finished
	case "RemoveEntry":
		go func() {
			nextID = RemoveEntry(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
