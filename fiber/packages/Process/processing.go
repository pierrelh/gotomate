package process

import (
	"fmt"
	"reflect"
)

// Processing process the functions from process's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "GetTitle":
		go func() {
			nextID = GetTitle(instructionData, finished)
		}()
		<-finished
	case "GetPid":
		go func() {
			nextID = GetPid(instructionData, finished)
		}()
		<-finished
	case "KillProcess":
		go func() {
			nextID = KillProcess(instructionData, finished)
		}()
		<-finished
	case "MaxSize":
		go func() {
			nextID = MaxSize(instructionData, finished)
		}()
		<-finished
	case "Reduce":
		go func() {
			nextID = Reduce(instructionData, finished)
		}()
		<-finished
	case "StartProcess":
		go func() {
			nextID = StartProcess(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
