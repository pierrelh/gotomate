package process

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"os"
	"os/exec"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// GetTitle get the title of a process
func GetTitle(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Getting the title of a process")

	pid, err := variable.GetValue(instructionData, "PIDVarName", "PIDIsVar", "PID")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetTitle(int32(pid.(int))))
	finished <- true
	return -1
}

// GetPid get the pid of the active window
func GetPid(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Getting a process's pid")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetPID())
	finished <- true
	return -1
}

// KillProcess Kill a process by his pid
func KillProcess(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Killing a process")

	pid, err := variable.GetValue(instructionData, "PIDVarName", "PIDIsVar", "PID")
	if err != nil {
		finished <- true
		return -1
	}

	proc, err := os.FindProcess(pid.(int))
	if err != nil {
		log.FiberError("Unable to find the process with pid", pid)
		finished <- true
		return -1
	}
	proc.Kill()
	finished <- true
	return -1
}

// MaxSize set the max size for a process
func MaxSize(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Maximizing a process's size")

	pid, err := variable.GetValue(instructionData, "PIDVarName", "PIDIsVar", "PID")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.MaxWindow(int32(pid.(int)))
	finished <- true
	return -1
}

// Reduce a process
func Reduce(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Reducing a process")

	pid, err := variable.GetValue(instructionData, "PIDVarName", "PIDIsVar", "PID")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.MinWindow(int32(pid.(int)))
	finished <- true
	return -1
}

// StartProcess Create a process with a given path & return the process's pid
func StartProcess(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Starting a new process")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	cmd := exec.Command(path.(string))
	cmd.Env = os.Environ()
	cmd.Start()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), cmd.Process.Pid)
	finished <- true
	return -1
}
