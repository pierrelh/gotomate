package clipboard

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Read read string from clipboard
func Read(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clipboard Reading ...")
	content, _ := robotgo.ReadAll()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), content)
	finished <- true
	return -1
}

// Write write string to clipboard
func Write(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clipboard Writing ...")

	content, err := variable.GetValue(instructionData, "VarName", "IsVar", "Content")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.WriteAll(content.(string))
	finished <- true
	return -1
}
