package json

import (
	"encoding/json"
	"fmt"
	variable "gotomate/fiber/variable"
	"io/ioutil"
	"os"
	"reflect"
)

// CreateJson Create a new json value with a json file
func CreateJson(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting a json value by path ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	jsonFile, err := os.Open(path.(string))
	if err != nil {
		fmt.Println("FIBER ERROR: Cannot open json file's path")
		finished <- true
		return -1
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), byteValue)
	finished <- true
	return -1
}

// JsonToDictionary Convert a json to a dictionary
func JsonToDictionary(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Converting json to dictionary ...")

	input, err := variable.GetValue(instructionData, "JsonVarName")
	if err != nil {
		finished <- true
		return -1
	}

	m := make(map[string][]interface{})
	json.Unmarshal(input.([]byte), &m)

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), m)
	finished <- true
	return -1
}

// SaveJson Save a json value in a file
func SaveJson(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Saving json to path ...")

	json, err := variable.GetValue(instructionData, "JsonVarName")
	if err != nil {
		finished <- true
		return -1
	}

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	err = ioutil.WriteFile(path.(string), json.([]byte), 0644)
	if err != nil {
		fmt.Println("FIBER ERROR: Unable to write the new json file")
		finished <- true
		return -1
	}
	finished <- true
	return -1
}
