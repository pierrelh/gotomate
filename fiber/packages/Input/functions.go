package input

import (
	"fmt"
	"gotomate/fiber/variable"
	"gotomate/log"
)

// Bool Wait for user to set a bool
func Bool(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Waiting for user input of a Bool")

	msg, err := variable.Keys{VarName: "MessageVarName", IsVarName: "MessageIsVar", Name: "Message"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input bool
	fmt.Scanln(&input)
	variable.SetVariable(instructionData, "Output", input)
	finished <- true
	return -1
}

// Float Wait for user to set a Float
func Float(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Waiting for user input of a Float")

	msg, err := variable.Keys{VarName: "MessageVarName", IsVarName: "MessageIsVar", Name: "Message"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input float64
	fmt.Scanln(&input)
	variable.SetVariable(instructionData, "Output", input)
	finished <- true
	return -1
}

// Int Wait for user to set a int
func Int(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Waiting for user input of an Int")

	msg, err := variable.Keys{VarName: "MessageVarName", IsVarName: "MessageIsVar", Name: "Message"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input int
	fmt.Scanln(&input)
	variable.SetVariable(instructionData, "Output", input)
	finished <- true
	return -1
}

// String Wait for user to set a string
func String(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Waiting for user input of a String")

	msg, err := variable.Keys{VarName: "MessageVarName", IsVarName: "MessageIsVar", Name: "Message"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}
	log.Plain(msg)
	var input string
	fmt.Scanln(&input)
	variable.SetVariable(instructionData, "Output", input)
	finished <- true
	return -1
}
