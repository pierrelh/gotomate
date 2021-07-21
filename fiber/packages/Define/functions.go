package define

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"strconv"
	"strings"
)

// ArrayOfBool Define an array of bool in a flow
func ArrayOfBool(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining an array of Bools")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	array := strings.Split(value.(string), ",")
	var boolArray []bool
	for i := 0; i < len(array); i++ {
		boolvalue, _ := strconv.ParseBool(array[i])
		boolArray = append(boolArray, boolvalue)
	}

	variable.SetVariable(instructionData, "Name", boolArray)
	finished <- true
	return -1
}

// ArrayOfFloat Define an array of float in a flow
func ArrayOfFloat(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining an array of Floats")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	array := strings.Split(value.(string), ",")
	var floatArray []float64
	for i := 0; i < len(array); i++ {
		floatvalue, _ := strconv.ParseFloat(array[i], 64)
		floatArray = append(floatArray, floatvalue)
	}

	variable.SetVariable(instructionData, "Name", floatArray)
	finished <- true
	return -1
}

// ArrayOfInt Define an array of int in a flow
func ArrayOfInt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining an array of Ints")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	array := strings.Split(value.(string), ",")
	var intArray []int
	for i := 0; i < len(array); i++ {
		intValue, _ := strconv.Atoi(array[i])
		intArray = append(intArray, intValue)
	}

	variable.SetVariable(instructionData, "Name", intArray)
	finished <- true
	return -1
}

// ArrayOfString Define an array of string in a flow
func ArrayOfString(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining an array of Strings")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Name", strings.Split(value.(string), ","))
	finished <- true
	return -1
}

// Bool Define a bool value in a flow
func Bool(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining a Bool value")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Name", value.(bool))
	finished <- true
	return -1
}

// Float Define a float value in a flow
func Float(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining a Float value")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Name", value.(float64))
	finished <- true
	return -1
}

// Int Define an int value in a flow
func Int(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining an Int value")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Name", value.(int))
	finished <- true
	return -1
}

// String Define a string value in a flow
func String(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Defining a String value")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Name", value.(string))
	finished <- true
	return -1
}
