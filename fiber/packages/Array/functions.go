package array

import (
	"fmt"
	"gotomate/fiber/variable"
	"math/rand"
	"reflect"
	"time"
)

// GetArrayLength Get the current length of an array
func GetArrayLength(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Get the lenth of an array ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]bool)))
	case "[]float64":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]float64)))
	case "[]int":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]int)))
	case "[]string":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), len(array.([]string)))
	}
	finished <- true
	return -1
}

// GetValue Get a value from an array by index
func GetValue(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting a value from array ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.GetValue(instructionData, "IndexVarName", "IndexIsVar", "Index")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]bool)[index.(int)])
	case "[]float64":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]float64)[index.(int)])
	case "[]int":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]int)[index.(int)])
	case "[]string":
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), array.([]string)[index.(int)])
	}
	finished <- true
	return -1
}

// PopAt Pop the wanted index of an array
func PopAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pop array at index ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.GetValue(instructionData, "IndexVarName", "IndexIsVar", "Index")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		popped := array.([]bool)[index.(int)]
		copy(array.([]bool)[index.(int):], array.([]bool)[index.(int)+1:])
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		popped := array.([]float64)[index.(int)]
		copy(array.([]float64)[index.(int):], array.([]float64)[index.(int)+1:])
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		popped := array.([]int)[index.(int)]
		copy(array.([]int)[index.(int):], array.([]int)[index.(int)+1:])
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		popped := array.([]string)[index.(int)]
		copy(array.([]string)[index.(int):], array.([]string)[index.(int)+1:])
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// PopLast Pop the last index of an array
func PopLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Pop array at end ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		popped := array.([]bool)[len(array.([]bool))-1]
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		popped := array.([]float64)[len(array.([]float64))-1]
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		popped := array.([]int)[len(array.([]int))-1]
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		popped := array.([]string)[len(array.([]string))-1]
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), popped)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// PushAt Push a value at the wanted index of an array
func PushAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Push value in array at index ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.GetValue(instructionData, "IndexVarName", "IndexIsVar", "Index")
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.GetValue(instructionData, "ValueVarName")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = append(array.([]bool), false)
		copy(array.([]bool)[index.(int)+1:], array.([]bool)[index.(int):])
		array.([]bool)[index.(int)] = value.(bool)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		array = append(array.([]bool), false)
		copy(array.([]float64)[index.(int)+1:], array.([]float64)[index.(int):])
		array.([]float64)[index.(int)] = value.(float64)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		array = append(array.([]int), 0)
		copy(array.([]int)[index.(int)+1:], array.([]int)[index.(int):])
		array.([]int)[index.(int)] = value.(int)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		array = append(array.([]string), "")
		copy(array.([]string)[index.(int)+1:], array.([]string)[index.(int):])
		array.([]string)[index.(int)] = value.(string)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// PushLast Push a value at the end of an array
func PushLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Push value in array at end ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.GetValue(instructionData, "ValueVarName")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = append(array.([]bool), value.(bool))
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		array = append(array.([]float64), value.(float64))
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		array = append(array.([]int), value.(int))
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		array = append(array.([]string), value.(string))
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// RemoveAt Remove a value by the index, of an array
func RemoveAt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Remove value from array at index ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.GetValue(instructionData, "IndexVarName", "IndexIsVar", "Index")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		copy(array.([]bool)[index.(int):], array.([]bool)[index.(int)+1:])
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		copy(array.([]float64)[index.(int):], array.([]float64)[index.(int)+1:])
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		copy(array.([]int)[index.(int):], array.([]int)[index.(int)+1:])
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		copy(array.([]string)[index.(int):], array.([]string)[index.(int)+1:])
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// RemoveLast Remove the last value of an array
func RemoveLast(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Remove value from array at end ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array = array.([]bool)[:len(array.([]bool))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		array = array.([]float64)[:len(array.([]float64))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		array = array.([]int)[:len(array.([]int))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		array = array.([]string)[:len(array.([]string))-1]
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// Shuffle an array
func Shuffle(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Shuffling array ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	rand.Seed(time.Now().UnixNano())

	switch variable.GetVariableType(array) {
	case "[]bool":
		rand.Shuffle(len(array.([]bool)), func(i, j int) { array.([]bool)[i], array.([]bool)[j] = array.([]bool)[j], array.([]bool)[i] })
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		rand.Shuffle(len(array.([]float64)), func(i, j int) {
			array.([]float64)[i], array.([]float64)[j] = array.([]float64)[j], array.([]float64)[i]
		})
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		rand.Shuffle(len(array.([]int)), func(i, j int) { array.([]int)[i], array.([]int)[j] = array.([]int)[j], array.([]int)[i] })
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		rand.Shuffle(len(array.([]string)), func(i, j int) { array.([]string)[i], array.([]string)[j] = array.([]string)[j], array.([]string)[i] })
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}

// UpdateValue Update a value of an array by index
func UpdateValue(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Updating value in array by index ...")

	array, err := variable.GetValue(instructionData, "ArrayVarName")
	if err != nil {
		finished <- true
		return -1
	}

	index, err := variable.GetValue(instructionData, "IndexVarName", "IndexIsVar", "Index")
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.GetValue(instructionData, "ValueVarName")
	if err != nil {
		finished <- true
		return -1
	}

	switch variable.GetVariableType(array) {
	case "[]bool":
		array.([]bool)[index.(int)] = value.(bool)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]bool))
	case "[]float64":
		array.([]float64)[index.(int)] = value.(float64)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]float64))
	case "[]int":
		array.([]int)[index.(int)] = value.(int)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]int))
	case "[]string":
		array.([]string)[index.(int)] = value.(string)
		variable.SetVariable(instructionData.FieldByName("ArrayVarName").Interface().(string), array.([]string))
	}
	finished <- true
	return -1
}
