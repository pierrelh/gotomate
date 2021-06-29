package conversion

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"
	"strconv"
)

// BoolToFloat Convert a bool to a float
func BoolToFloat(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Bool to Float ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
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

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), value)
	finished <- true
	return -1
}

// BoolToInt Convert a bool to an int
func BoolToInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Bool to Int ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
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

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), value)
	finished <- true
	return -1
}

// BoolToString Convert a bool to a string
func BoolToString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Bool to String ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), strconv.FormatBool(input.(bool)))
	finished <- true
	return -1
}

// FloatToInt Convert a float to an int
func FloatToInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Float to Int ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), int(input.(float64)))
	finished <- true
	return -1
}

// FloatToString Convert a float to a string
func FloatToString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Float to String ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), strconv.FormatFloat(input.(float64), 'E', -1, 64))
	finished <- true
	return -1
}

// IntToFloat Convert an int to a float
func IntToFloat(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Int to Float ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	name := instructionData.FieldByName("Output").Interface().(string)
	value := float64(input.(int))
	variable.SetVariable(name, value)
	finished <- true
	return -1
}

// IntToString Convert an int to a string
func IntToString(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting Int to String ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), strconv.Itoa(input.(int)))
	finished <- true
	return -1
}

// StringToBool Convert a string to a bool
func StringToBool(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting String to Bool ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	value, _ := strconv.ParseBool(input.(string))
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), value)
	finished <- true
	return -1
}

// StringToFloat Convert a string to a float
func StringToFloat(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting String to Float ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	value, _ := strconv.ParseFloat(input.(string), 64)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), value)
	finished <- true
	return -1
}

// StringToInt Convert a string to an int
func StringToInt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting String to Int ...")

	input, err := variable.GetValue(instructionData, "InputVarName", "InputIsVar", "Input")
	if err != nil {
		finished <- true
		return -1
	}

	value, _ := strconv.Atoi(input.(string))
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), value)
	finished <- true
	return -1
}
