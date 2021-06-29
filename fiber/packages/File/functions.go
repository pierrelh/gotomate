package file

import (
	"fmt"
	"gotomate/fiber/variable"
	"io/ioutil"
	"os"
	"reflect"
)

// Create a new file
func Create(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Creating File ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	f, _ := os.Create(path.(string))
	f.Close()
	finished <- true
	return -1
}

// Delete an existing file
func Delete(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Deleting File ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	os.Remove(path.(string))
	finished <- true
	return -1
}

// Read a file
func Read(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Reading File ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	content, _ := ioutil.ReadFile(path.(string))
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), string(content))
	finished <- true
	return -1
}

// Write an existing file
func Write(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Reading File ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	content, err := variable.GetValue(instructionData, "ContentVarName", "ContentIsVar", "Content")
	if err != nil {
		finished <- true
		return -1
	}

	ioutil.WriteFile(path.(string), []byte(content.(string)), 0644)
	finished <- true
	return -1
}
