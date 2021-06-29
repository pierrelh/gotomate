package arithmetic

import (
	"fmt"
	"gotomate/fiber/variable"
	"math"
	"reflect"
)

// Divide make the divide of two numbers
func Divide(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Dividing ...")

	one, err := variable.GetValue(instructionData, "VarOneName", "ValueOneIsVar", "ValueOne")
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.GetValue(instructionData, "VarTwoName", "ValueTwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), variable.GetFloat(one)/variable.GetFloat(two))
	finished <- true
	return -1
}

// Multiply make the multiplication of two numbers
func Multiply(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Multiplying ...")

	one, err := variable.GetValue(instructionData, "VarOneName", "ValueOneIsVar", "ValueOne")
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.GetValue(instructionData, "VarTwoName", "ValueTwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), variable.GetFloat(one)*variable.GetFloat(two))
	finished <- true
	return -1
}

// Pow make the pow of a number
func Pow(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pow ...")

	one, err := variable.GetValue(instructionData, "VarOneName", "ValueOneIsVar", "ValueOne")
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.GetValue(instructionData, "VarTwoName", "ValueTwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), math.Pow(variable.GetFloat(one), variable.GetFloat(two)))
	finished <- true
	return -1
}

// Sqrt make the Sqrt of a number
func Sqrt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Sqrt ...")

	value, err := variable.GetValue(instructionData, "VarName", "ValueIsVar", "Value")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), math.Sqrt(variable.GetFloat(value)))
	finished <- true
	return -1
}

// Substract make the substraction of two numbers
func Substract(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Substracting ...")

	one, err := variable.GetValue(instructionData, "VarOneName", "ValueOneIsVar", "ValueOne")
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.GetValue(instructionData, "VarTwoName", "ValueTwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), variable.GetFloat(one)-variable.GetFloat(two))
	finished <- true
	return -1
}

// Sum make the sum of two numbers
func Sum(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Additioning ...")

	one, err := variable.GetValue(instructionData, "VarOneName", "ValueOneIsVar", "ValueOne")
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.GetValue(instructionData, "VarTwoName", "ValueTwoIsVar", "ValueTwo")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), variable.GetFloat(one)+variable.GetFloat(two))
	finished <- true
	return -1
}
