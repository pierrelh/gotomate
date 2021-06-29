package process

import (
	"fmt"
	"gotomate/fiber/variable"
	"os"
	"os/exec"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// GetTitle get the title of a process
func GetTitle(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting title of process ...")

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
	fmt.Println("FIBER INFO: Getting process pid ...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetPID())
	finished <- true
	return -1
}

// KillProcess Kill a process by his pid
func KillProcess(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Killing process ...")

	pid, err := variable.GetValue(instructionData, "PIDVarName", "PIDIsVar", "PID")
	if err != nil {
		finished <- true
		return -1
	}

	proc, err := os.FindProcess(pid.(int))
	if err != nil {
		fmt.Println("FIBER WARNING: Didn't find process with pid", pid)
		finished <- true
		return -1
	}
	proc.Kill()
	finished <- true
	return -1
}

// MaxSize set the max size for a process
func MaxSize(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Maximizing process size ...")

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
	fmt.Println("FIBER INFO: Reducing a process ...")

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
	fmt.Println("FIBER INFO: Starting new process ...")

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
