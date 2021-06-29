package keyboard

import (
	"fmt"
	"gotomate-astilectron/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Tap Simulate a tap on the keyboard
func Tap(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Keyboard tap ...")

	special := []string{}
	if err := instructionData.FieldByName("Special1").Interface().(string); err != "" {
		special = append(special, err)
	}
	if err := instructionData.FieldByName("Special2").Interface().(string); err != "" {
		special = append(special, err)
	}
	if len(special) == 0 {
		robotgo.KeyTap(instructionData.FieldByName("Input").Interface().(string))
	} else {
		robotgo.KeyTap(instructionData.FieldByName("Input").Interface().(string), special)
	}
	finished <- true
	return -1
}

// Type Simulate a type on the keyboard
func Type(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Keyboard type ...")

	input, err := variable.GetValue(instructionData, "VarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.TypeStr(input.(string))
	finished <- true
	return -1
}
