package api

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"io/ioutil"
	"net/http"
)

// Get Send a Get Request
func Get(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Sending GET request")

	path, err := variable.Keys{VarName: "PathVarName", IsVarName: "PathIsVar", Name: "Path"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	resp, err := http.Get(path.(string))
	if err != nil {
		log.FiberError("Error while getting response from GET request")
		finished <- true
		return -1
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.FiberError("Error while reading response from GET request")
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", string(body))
	finished <- true
	return -1
}

// Post Send a Post request
func Post(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Sending POST request")

	data, err := variable.Keys{VarName: "DataVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	datas := data.(map[string][]interface{})
	values := make(map[string][]string)
	for key, element := range datas {
		values[key] = variable.ToArrayOfString(element)
	}

	path, err := variable.Keys{VarName: "PathVarName", IsVarName: "PathIsVar", Name: "Path"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	resp, err := http.PostForm(path.(string), values)
	if err != nil {
		log.FiberError("Error while posting data to POST request")
	}

	variable.SetVariable(instructionData, "Output", resp)
	finished <- true
	return -1
}
