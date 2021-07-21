package arithmetic

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"math"
)

// Divide make the divide of two numbers
func Divide(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Dividing")

	one, err := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", variable.GetFloat(one)/variable.GetFloat(two))
	finished <- true
	return -1
}

// Multiply make the multiplication of two numbers
func Multiply(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Multiplying")

	one, err := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", variable.GetFloat(one)*variable.GetFloat(two))
	finished <- true
	return -1
}

// Pow make the pow of a number
func Pow(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Pow")

	one, err := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", math.Pow(variable.GetFloat(one), variable.GetFloat(two)))
	finished <- true
	return -1
}

// Sqrt make the Sqrt of a number
func Sqrt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Sqrt")

	value, err := variable.Keys{VarName: "VarName", IsVarName: "ValueIsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", math.Sqrt(variable.GetFloat(value)))
	finished <- true
	return -1
}

// Substract make the substraction of two numbers
func Substract(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Substracting")

	one, err := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", variable.GetFloat(one)-variable.GetFloat(two))
	finished <- true
	return -1
}

// Sum make the sum of two numbers
func Sum(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Additioning")

	one, err := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	two, err := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", variable.GetFloat(one)+variable.GetFloat(two))
	finished <- true
	return -1
}
