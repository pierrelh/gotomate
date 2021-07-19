package file

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"io/ioutil"
	"os"
	"reflect"
)

// Create a new file
func Create(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Creating a File")

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
	log.FiberInfo("Deleting a File")

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
	log.FiberInfo("Reading a File")

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
	log.FiberInfo("Writing a File")

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
