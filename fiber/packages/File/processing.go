package file

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Create":
		go func() {
			nextID = Create(instructionData, finished)
		}()
		<-finished
	case "Delete":
		go func() {
			nextID = Delete(instructionData, finished)
		}()
		<-finished
	case "Read":
		go func() {
			nextID = Read(instructionData, finished)
		}()
		<-finished
	case "Write":
		go func() {
			nextID = Write(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
