package process

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a process instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "GetTitle":
		return new(GetTitleDatabinder), GetTitleTemplate
	case "GetPid":
		return new(GetPidDatabinder), GetPidTemplate
	case "KillProcess":
		return new(KillProcessDatabinder), KillProcessTemplate
	case "MaxSize":
		return new(KillProcessDatabinder), KillProcessTemplate
	case "Reduce":
		return new(KillProcessDatabinder), KillProcessTemplate
	case "StartProcess":
		return new(StartProcessDatabinder), StartProcessTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
