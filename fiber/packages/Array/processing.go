package array

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "GetArrayLength":
		go func() {
			nextID = GetArrayLength(instructionData, finished)
		}()
		<-finished
	case "GetValue":
		go func() {
			nextID = GetValue(instructionData, finished)
		}()
		<-finished
	case "PopAt":
		go func() {
			nextID = PopAt(instructionData, finished)
		}()
		<-finished
	case "PopLast":
		go func() {
			nextID = PopLast(instructionData, finished)
		}()
		<-finished
	case "PushAt":
		go func() {
			nextID = PushAt(instructionData, finished)
		}()
		<-finished
	case "PushLast":
		go func() {
			nextID = PushLast(instructionData, finished)
		}()
		<-finished
	case "RemoveAt":
		go func() {
			nextID = RemoveAt(instructionData, finished)
		}()
		<-finished
	case "RemoveLast":
		go func() {
			nextID = RemoveLast(instructionData, finished)
		}()
		<-finished
	case "Shuffle":
		go func() {
			nextID = Shuffle(instructionData, finished)
		}()
		<-finished
	case "UpdateValue":
		go func() {
			nextID = UpdateValue(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
