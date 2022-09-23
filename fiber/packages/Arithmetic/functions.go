package arithmetic

import (
	"gotomate/fiber/variable"
	"gotomate/log"
	"math"
)

// Divide make the divide of two numbers
func Divide(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Dividing")

	one, err1 := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	two, err2 := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
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

	one, err1 := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	two, err2 := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
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

	one, err1 := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	two, err2 := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
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

	one, err1 := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	two, err2 := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
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

	one, err1 := variable.Keys{VarName: "VarOneName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	two, err2 := variable.Keys{VarName: "VarTwoName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", variable.GetFloat(one)+variable.GetFloat(two))
	finished <- true
	return -1
}
