package screen

import (
	"fmt"
	"reflect"
)

// Processing process the functions from screen's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "GetMouseColor":
		go func() {
			nextID = GetMouseColor(instructionData, finished)
		}()
		<-finished
	case "GetPixelColor":
		go func() {
			nextID = GetPixelColor(instructionData, finished)
		}()
		<-finished
	case "GetScreenSize":
		go func() {
			nextID = GetScreenSize(instructionData, finished)
		}()
		<-finished
	case "PartScreenShot":
		go func() {
			nextID = PartScreenShot(instructionData, finished)
		}()
		<-finished
	case "ScreenShot":
		go func() {
			nextID = ScreenShot(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
