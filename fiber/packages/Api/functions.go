package api

import (
	"fmt"
	variable "gotomate/fiber/variable"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Get Send a Get Request
func Get(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting path ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	resp, err := http.Get(path.(string))
	if err != nil {
		fmt.Println("FIBER ERROR: Error while getting response from Get")
		finished <- true
		return -1
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FIBER ERROR: Error while reading response from Get")
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), string(body))
	finished <- true
	return -1
}

// Post Send a Post request
func Post(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Post Statement ...")

	data, err := variable.GetValue(instructionData, "DataVarName")
	if err != nil {
		finished <- true
		return -1
	}

	datas := data.(map[string][]interface{})
	values := make(map[string][]string)
	for key, element := range datas {
		values[key] = variable.ToArrayOfString(element)
	}

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	resp, err := http.PostForm(path.(string), values)
	if err != nil {
		fmt.Println("FIBER ERROR: Error while posting data")
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), resp)
	finished <- true
	return -1
}
