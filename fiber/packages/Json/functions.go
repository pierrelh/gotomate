package json

import (
	"encoding/json"
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"io/ioutil"
	"os"
	"reflect"
)

// CreateJson Create a new json value with a json file
func CreateJson(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Creating a Json from a File")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	jsonFile, err := os.Open(path.(string))
	if err != nil {
		log.FiberError("Cannot open Json file's path")
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
	log.FiberInfo("Converting a Json to a Dictionary")

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
	log.FiberInfo("Saving a Json to a path")

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
		log.FiberError("Unable to write the new json file")
		finished <- true
		return -1
	}
	finished <- true
	return -1
}
