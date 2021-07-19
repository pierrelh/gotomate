package dictionary

import (
	"encoding/json"
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"
)

// CreateDictionary Create a new Dictionary
func CreateDictionary(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Creating a new Dictionary")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), make(map[string][]interface{}))
	finished <- true
	return -1
}

// CreateEntry Create a new entry in a dictionary
func CreateEntry(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Create an entry in a Dictionary")

	dict, err := variable.GetValue(instructionData, "DictVarName")
	if err != nil {
		finished <- true
		return -1
	}

	key, err := variable.GetValue(instructionData, "KeyVarName", "KeyIsVar", "Key")
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.GetValue(instructionData, "ValueVarName", "ValueIsVar", "Value")
	if err != nil {
		finished <- true
		return -1
	}

	newMap := dict.(map[string][]interface{})
	newMap[key.(string)] = append(newMap[key.(string)], value)
	variable.SetVariable(instructionData.FieldByName("DictVarName").Interface().(string), newMap)
	finished <- true
	return -1
}

// DictionaryToJson Convert a dictionary to Json
func DictionaryToJson(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Converting a Dictionary to Json")

	dict, err := variable.GetValue(instructionData, "DictVarName")
	if err != nil {
		finished <- true
		return -1
	}

	json, err := json.Marshal(dict.(map[string][]interface{}))
	if err != nil {
		log.FiberError("Cannot convert the Dictionnary to a Json")
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), json)
	finished <- true
	return -1
}

// RemoveEntry Remove an entry from a Dictionary
func RemoveEntry(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Removing an entry from a Dictionay")

	dict, err := variable.GetValue(instructionData, "DictVarName")
	if err != nil {
		finished <- true
		return -1
	}

	key, err := variable.GetValue(instructionData, "KeyVarName", "KeyIsVar", "Key")
	if err != nil {
		finished <- true
		return -1
	}

	newMap := dict.(map[string][]interface{})
	newMap[key.(string)] = nil
	_, ok := newMap[key.(string)]
	if ok {
		delete(newMap, key.(string))
	}

	variable.SetVariable(instructionData.FieldByName("DictVarName").Interface().(string), newMap)
	finished <- true
	return -1
}
