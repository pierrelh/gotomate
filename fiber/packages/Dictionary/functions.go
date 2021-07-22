package dictionary

import (
	"encoding/json"
	"gotomate/fiber/variable"
	"gotomate/log"
)

// CreateDictionary Create a new Dictionary
func CreateDictionary(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Creating a new Dictionary")

	variable.SetVariable(instructionData, "Output", make(map[string][]interface{}))
	finished <- true
	return -1
}

// CreateEntry Create a new entry in a dictionary
func CreateEntry(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Create an entry in a Dictionary")

	dict, err := variable.Keys{VarName: "DictVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	key, err := variable.Keys{VarName: "KeyVarName", IsVarName: "KeyIsVar", Name: "Key"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	value, err := variable.Keys{VarName: "ValueVarName", IsVarName: "ValueIsVar", Name: "Value"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	newMap := dict.(map[string][]interface{})
	newMap[key.(string)] = append(newMap[key.(string)], value)
	variable.SetVariable(instructionData, "DictVarName", newMap)
	finished <- true
	return -1
}

// DictionaryToJson Convert a dictionary to Json
func DictionaryToJson(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Converting a Dictionary to Json")

	dict, err := variable.Keys{VarName: "DictVarName"}.GetValue(instructionData)
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

	variable.SetVariable(instructionData, "Output", json)
	finished <- true
	return -1
}

// RemoveEntry Remove an entry from a Dictionary
func RemoveEntry(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Removing an entry from a Dictionay")

	dict, err := variable.Keys{VarName: "DictVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	key, err := variable.Keys{VarName: "KeyVarName", IsVarName: "KeyIsVar", Name: "Key"}.GetValue(instructionData)
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

	variable.SetVariable(instructionData, "DictVarName", newMap)
	finished <- true
	return -1
}
