package array

import (
	"gotomate/fiber/variable"
	"gotomate/log"
	"math/rand"
	"time"
)

// GetArrayLength Get the current length of an array
func GetArrayLength(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting an array length")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		variable.SetVariable(instructionData, "Output", len(array.([]bool)))
	case "[]float64":
		variable.SetVariable(instructionData, "Output", len(array.([]float64)))
	case "[]int":
		variable.SetVariable(instructionData, "Output", len(array.([]int)))
	case "[]string":
		variable.SetVariable(instructionData, "Output", len(array.([]string)))
	}
	finished <- true
	return -1
}

// GetValue Get a value from an array by index
func GetValue(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a value from an array")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.Keys{VarName: "IndexVarName", IsVarName: "IndexIsVar", Name: "Index"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		variable.SetVariable(instructionData, "Output", array.([]bool)[index.(int)])
	case "[]float64":
		variable.SetVariable(instructionData, "Output", array.([]float64)[index.(int)])
	case "[]int":
		variable.SetVariable(instructionData, "Output", array.([]int)[index.(int)])
	case "[]string":
		variable.SetVariable(instructionData, "Output", array.([]string)[index.(int)])
	}
	finished <- true
	return -1
}

// PopAt Pop the wanted index of an array
func PopAt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Poping value of array at index")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.Keys{VarName: "IndexVarName", IsVarName: "IndexIsVar", Name: "Index"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		popped := array.([]bool)[index.(int)]
		copy(array.([]bool)[index.(int):], array.([]bool)[index.(int)+1:])
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		popped := array.([]float64)[index.(int)]
		copy(array.([]float64)[index.(int):], array.([]float64)[index.(int)+1:])
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		popped := array.([]int)[index.(int)]
		copy(array.([]int)[index.(int):], array.([]int)[index.(int)+1:])
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		popped := array.([]string)[index.(int)]
		copy(array.([]string)[index.(int):], array.([]string)[index.(int)+1:])
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// PopLast Pop the last index of an array
func PopLast(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Poping value of array at end")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		popped := array.([]bool)[len(array.([]bool))-1]
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		popped := array.([]float64)[len(array.([]float64))-1]
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		popped := array.([]int)[len(array.([]int))-1]
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		popped := array.([]string)[len(array.([]string))-1]
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData, "Output", popped)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// PushAt Push a value at the wanted index of an array
func PushAt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Pushing value in array at index")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.Keys{VarName: "IndexVarName", IsVarName: "IndexIsVar", Name: "Index"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.Keys{VarName: "ValueVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = append(array.([]bool), false)
		copy(array.([]bool)[index.(int)+1:], array.([]bool)[index.(int):])
		array.([]bool)[index.(int)] = value.(bool)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		array = append(array.([]bool), false)
		copy(array.([]float64)[index.(int)+1:], array.([]float64)[index.(int):])
		array.([]float64)[index.(int)] = value.(float64)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		array = append(array.([]int), 0)
		copy(array.([]int)[index.(int)+1:], array.([]int)[index.(int):])
		array.([]int)[index.(int)] = value.(int)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		array = append(array.([]string), "")
		copy(array.([]string)[index.(int)+1:], array.([]string)[index.(int):])
		array.([]string)[index.(int)] = value.(string)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// PushLast Push a value at the end of an array
func PushLast(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Pushing value in array at end")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.Keys{VarName: "ValueVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = append(array.([]bool), value.(bool))
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		array = append(array.([]float64), value.(float64))
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		array = append(array.([]int), value.(int))
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		array = append(array.([]string), value.(string))
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// RemoveAt Remove a value by the index, of an array
func RemoveAt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Removing value from array at index")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.Keys{VarName: "IndexVarName", IsVarName: "IndexIsVar", Name: "Index"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		copy(array.([]bool)[index.(int):], array.([]bool)[index.(int)+1:])
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		copy(array.([]float64)[index.(int):], array.([]float64)[index.(int)+1:])
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		copy(array.([]int)[index.(int):], array.([]int)[index.(int)+1:])
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		copy(array.([]string)[index.(int):], array.([]string)[index.(int)+1:])
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// RemoveLast Remove the last value of an array
func RemoveLast(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Removing value from array at end")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// Shuffle an array
func Shuffle(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Shuffling an array")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	rand.Seed(time.Now().UnixNano())

	switch variable.GetVariableType(array) {
	case "[]bool":
		rand.Shuffle(len(array.([]bool)), func(i, j int) { array.([]bool)[i], array.([]bool)[j] = array.([]bool)[j], array.([]bool)[i] })
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		rand.Shuffle(len(array.([]float64)), func(i, j int) {
			array.([]float64)[i], array.([]float64)[j] = array.([]float64)[j], array.([]float64)[i]
		})
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		rand.Shuffle(len(array.([]int)), func(i, j int) { array.([]int)[i], array.([]int)[j] = array.([]int)[j], array.([]int)[i] })
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		rand.Shuffle(len(array.([]string)), func(i, j int) { array.([]string)[i], array.([]string)[j] = array.([]string)[j], array.([]string)[i] })
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}

// UpdateValue Update a value of an array by index
func UpdateValue(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Updating a value in an array by index")

	array, err := variable.Keys{VarName: "ArrayVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.Keys{VarName: "IndexVarName", IsVarName: "IndexIsVar", Name: "Index"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.Keys{VarName: "ValueVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array.([]bool)[index.(int)] = value.(bool)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]bool))
	case "[]float64":
		array.([]float64)[index.(int)] = value.(float64)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]float64))
	case "[]int":
		array.([]int)[index.(int)] = value.(int)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]int))
	case "[]string":
		array.([]string)[index.(int)] = value.(string)
		variable.SetVariable(instructionData, "ArrayVarName", array.([]string))
	}
	finished <- true
	return -1
}
