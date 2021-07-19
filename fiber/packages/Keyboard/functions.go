package keyboard

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Tap Simulate a tap on the keyboard
func Tap(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Taping on keyboard")

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
	log.FiberInfo("Typing on keyboard")

	input, err := variable.GetValue(instructionData, "VarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.TypeStr(input.(string))
	finished <- true
	return -1
}
