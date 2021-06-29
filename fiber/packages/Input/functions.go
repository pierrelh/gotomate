package input

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
)

// Bool Wait for user to set a bool
func Bool(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Waiting for user input ...")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	fmt.Println(msg)
	var input bool
	fmt.Scanln(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}

// Float Wait for user to set a Float
func Float(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Waiting for user input ...")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	fmt.Println(msg)
	var input float64
	fmt.Scanln(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}

// Int Wait for user to set a int
func Int(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Waiting for user input ...")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	fmt.Println(msg)
	var input int
	fmt.Scanln(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}

// String Wait for user to set a string
func String(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Waiting for user input ...")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	fmt.Println(msg)
	var input string
	fmt.Scanln(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}
