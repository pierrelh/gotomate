package process

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a process instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
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
	return nil, nil
}
