package command

import (
	"gotomate/fiber/variable"
	"gotomate/log"
	"os/exec"
)

// Dir set the dir of a command
func Dir(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting command's dir")

	command, err1 := variable.Keys{VarName: "CommandVarName"}.GetValue(instructionData)
	dir, err2 := variable.Keys{Name: "Dir", IsVarName: "DirIsVar", VarName: "DirVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	command.(*exec.Cmd).Dir = dir.(string)

	variable.SetVariable(instructionData, "CommandVarName", command)
	finished <- true
	return -1
}

// GetOutput get the output of a command
func GetOutput(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting command's output")

	command, err := variable.Keys{VarName: "CommandVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	out, _ := command.(*exec.Cmd).Output()

	variable.SetVariable(instructionData, "Output", out)
	finished <- true
	return -1
}

// New create a new command tool
func New(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Creating a new command tool")

	parameter, err1 := variable.Keys{Name: "Parameter", IsVarName: "ParameterIsVar", VarName: "ParameterVarName"}.GetValue(instructionData)
	progname, err2 := variable.Keys{Name: "Program", IsVarName: "ProgramIsVar", VarName: "ProgramVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	command := exec.Command(progname.(string), parameter.(string))

	variable.SetVariable(instructionData, "Output", command)
	finished <- true
	return -1
}

// Run run a command
func Run(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Running a command")

	command, err := variable.Keys{VarName: "CommandVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	command.(*exec.Cmd).Run()
	finished <- true
	return -1
}

// Start start a command
func Start(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Starting a command")

	command, err := variable.Keys{VarName: "CommandVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	command.(*exec.Cmd).Start()
	finished <- true
	return -1
}
