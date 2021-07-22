package algorithmic

import (
	"gotomate/fiber/variable"
	"gotomate/log"
)

// For Execute a for loop
func For(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("For Statement")

	valueOne, err := variable.Keys{VarName: "ValueOneVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	valueTwo, err := variable.Keys{VarName: "ValueTwoVarName", Name: "ValueTwo", IsVarName: "ValueTwoIsVar"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	comparator, err := variable.Keys{Name: "Comparator"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	statement := false
	switch comparator {
	case "==":
		if valueOne == valueTwo {
			statement = true
		}
	case "!=":
		if valueOne != valueTwo {
			statement = true
		}
	case ">":
		if variable.GetFloat(valueOne) > variable.GetFloat(valueTwo) {
			statement = true
		}
	case ">=":
		if variable.GetFloat(valueOne) >= variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<":
		if variable.GetFloat(valueOne) < variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<=":
		if variable.GetFloat(valueOne) <= variable.GetFloat(valueTwo) {
			statement = true
		}
	}

	if statement {
		increment, err := variable.Keys{VarName: "IncrementVarName", IsVarName: "IncrementIsVar", Name: "Increment"}.GetValue(instructionData)
		if err != nil {
			finished <- true
			return -1
		}

		variable.SetVariable(instructionData, "ValueOneVarName", variable.GetFloat(valueOne)+float64(increment.(int)))
		finished <- true
		return -1
	} else {
		falseInstruction, err := variable.Keys{VarName: "FalseInstructionVarName", IsVarName: "FalseInstructionIsVar", Name: "FalseInstruction"}.GetValue(instructionData)
		if err != nil {
			finished <- true
			return -1
		}
		finished <- true
		return falseInstruction.(int)
	}
}

// If Compare if a statement is true
func If(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("If Statement")

	valueOne, err := variable.Keys{VarName: "ValueOneVarName", IsVarName: "ValueOneIsVar", Name: "ValueOne"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	valueTwo, err := variable.Keys{VarName: "ValueTwoVarName", IsVarName: "ValueTwoIsVar", Name: "ValueTwo"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	comparator, err := variable.Keys{Name: "Comparator"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	statement := false
	switch comparator {
	case "==":
		if valueOne == valueTwo {
			statement = true
		}
	case "!=":
		if valueOne != valueTwo {
			statement = true
		}
	case ">":
		if variable.GetFloat(valueOne) > variable.GetFloat(valueTwo) {
			statement = true
		}
	case ">=":
		if variable.GetFloat(valueOne) >= variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<":
		if variable.GetFloat(valueOne) < variable.GetFloat(valueTwo) {
			statement = true
		}
	case "<=":
		if variable.GetFloat(valueOne) <= variable.GetFloat(valueTwo) {
			statement = true
		}
	}
	finished <- true
	if statement {
		return -1
	} else {
		falseInstruction, err := variable.Keys{VarName: "FalseInstructionVarName", IsVarName: "FalseInstructionIsVar", Name: "FalseInstruction"}.GetValue(instructionData)
		if err != nil {
			finished <- true
			return -1
		}
		return falseInstruction.(int)
	}
}
