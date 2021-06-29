package screen

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// GetMouseColor Get a pixel color by mouse position
func GetMouseColor(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting mouse pixel color ...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetMouseColor())
	finished <- true
	return -1
}

// GetPixelColor Get a pixel color by a position
func GetPixelColor(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting pixel color ...")

	x, err := variable.GetValue(instructionData, "XVarName", "XIsVar", "X")
	if err != nil {
		finished <- true
		return -1
	}

	y, err := variable.GetValue(instructionData, "YVarName", "YIsVar", "Y")
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), robotgo.GetPixelColor(x.(int), y.(int)))
	finished <- true
	return -1
}

// GetScreenSize Get the screen size
func GetScreenSize(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting screen size ...")

	w, h := robotgo.GetScreenSize()
	variable.SetVariable(instructionData.FieldByName("HeightOutput").Interface().(string), h)
	variable.SetVariable(instructionData.FieldByName("WidthOutput").Interface().(string), w)
	finished <- true
	return -1
}

func PartScreenShot(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Saving screen shot ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	x, err := variable.GetValue(instructionData, "XVarName", "XIsVar", "X")
	if err != nil {
		finished <- true
		return -1
	}

	y, err := variable.GetValue(instructionData, "YVarName", "YIsVar", "Y")
	if err != nil {
		finished <- true
		return -1
	}

	w, err := variable.GetValue(instructionData, "WVarName", "WIsVar", "W")
	if err != nil {
		finished <- true
		return -1
	}

	h, err := variable.GetValue(instructionData, "HVarName", "HIsVar", "H")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.SaveCapture(path.(string), x.(int), y.(int), w.(int), h.(int))
	finished <- true
	return -1
}

// SaveCapture Save a screen shot
func ScreenShot(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Saving screen shot ...")

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.SaveCapture(path.(string))
	finished <- true
	return -1
}
