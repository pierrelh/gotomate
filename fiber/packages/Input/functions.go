package input

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"
)

// Bool Wait for user to set a bool
func Bool(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Waiting for user input of a Bool")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input bool
	log.FiberScan(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}

// Float Wait for user to set a Float
func Float(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Waiting for user input of a Float")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input float64
	log.FiberScan(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}

// Int Wait for user to set a int
func Int(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Waiting for user input of an Int")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input int
	log.FiberScan(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}

// String Wait for user to set a string
func String(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Waiting for user input of a String")

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input string
	log.FiberScan(&input)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), input)
	finished <- true
	return -1
}
