package clipboard

import (
	"gotomate/fiber/variable"
	"gotomate/log"

	"github.com/go-vgo/robotgo"
)

// Read read string from clipboard
func Read(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Reading clipboard")

	content, _ := robotgo.ReadAll()
	variable.SetVariable(instructionData, "Output", content)
	finished <- true
	return -1
}

// Write write string to clipboard
func Write(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Writing clipboard")

	content, err := variable.Keys{VarName: "VarName", IsVarName: "IsVar", Name: "Content"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.WriteAll(content.(string))
	finished <- true
	return -1
}
