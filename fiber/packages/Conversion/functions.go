package conversion

import (
	"gotomate/fiber/variable"
	"gotomate/log"
	"strconv"
)

// BoolToFloat Convert a bool to a float
func BoolToFloat(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Bool to Float")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	var value float64
	if input.(bool) {
		value = 1
	} else {
		value = 0
	}

	variable.SetVariable(instructionData, "Output", value)
	finished <- true
	return -1
}

// BoolToInt Convert a bool to an int
func BoolToInt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Bool to Int")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	var value int
	if input.(bool) {
		value = 1
	} else {
		value = 0
	}

	variable.SetVariable(instructionData, "Output", value)
	finished <- true
	return -1
}

// BoolToString Convert a bool to a string
func BoolToString(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Bool to String")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", strconv.FormatBool(input.(bool)))
	finished <- true
	return -1
}

// FloatToInt Convert a float to an int
func FloatToInt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Float to Int")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", int(input.(float64)))
	finished <- true
	return -1
}

// FloatToString Convert a float to a string
func FloatToString(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Float to String")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", strconv.FormatFloat(input.(float64), 'E', -1, 64))
	finished <- true
	return -1
}

// IntToFloat Convert an int to a float
func IntToFloat(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Int to Float")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value := float64(input.(int))
	variable.SetVariable(instructionData, "Output", value)
	finished <- true
	return -1
}

// IntToString Convert an int to a string
func IntToString(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting Int to String")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", strconv.Itoa(input.(int)))
	finished <- true
	return -1
}

// StringToBool Convert a string to a bool
func StringToBool(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting String to Bool")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, _ := strconv.ParseBool(input.(string))
	variable.SetVariable(instructionData, "Output", value)
	finished <- true
	return -1
}

// StringToFloat Convert a string to a float
func StringToFloat(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting String to Float")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, _ := strconv.ParseFloat(input.(string), 64)
	variable.SetVariable(instructionData, "Output", value)
	finished <- true
	return -1
}

// StringToInt Convert a string to an int
func StringToInt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting String to Int")

	input, err := variable.Keys{VarName: "InputVarName", IsVarName: "InputIsVar", Name: "Input"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, _ := strconv.Atoi(input.(string))
	variable.SetVariable(instructionData, "Output", value)
	finished <- true
	return -1
}
