package keyboard

import (
	"gotomate/fiber/variable"
	"gotomate/log"

	"github.com/go-vgo/robotgo"
)

// Tap Simulate a tap on the keyboard
func Tap(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Taping on keyboard")

	special := []string{}

	spe1, _ := variable.Keys{Name: "Special1"}.GetValue(instructionData)
	if spe1 != nil {
		special = append(special, spe1.(string))
	}

	spe2, _ := variable.Keys{Name: "Special2"}.GetValue(instructionData)
	if spe2 != nil {
		special = append(special, spe2.(string))
	}

	input, err := variable.Keys{Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	if len(special) == 0 {
		robotgo.KeyTap(input.(string))
	} else {
		robotgo.KeyTap(input.(string), special)
	}
	finished <- true
	return -1
}

// Type Simulate a type on the keyboard
func Type(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Typing on keyboard")

	input, err := variable.Keys{VarName: "VarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.TypeStr(input.(string))
	finished <- true
	return -1
}
