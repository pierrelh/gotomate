package screen

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"

	"github.com/go-vgo/robotgo"
)

// GetMouseColor Get a pixel color by mouse position
func GetMouseColor(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting the mouse pixel color")

	variable.SetVariable(instructionData, "Output", robotgo.GetMouseColor())
	finished <- true
	return -1
}

// GetPixelColor Get a pixel color by a position
func GetPixelColor(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a pixel color")

	x, err := variable.Keys{VarName: "XVarName", IsVarName: "XIsVar", Name: "X"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	y, err := variable.Keys{VarName: "YVarName", IsVarName: "YIsVar", Name: "Y"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	variable.SetVariable(instructionData, "Output", robotgo.GetPixelColor(x.(int), y.(int)))
	finished <- true
	return -1
}

// GetScreenSize Get the screen size
func GetScreenSize(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting the screen size")

	w, h := robotgo.GetScreenSize()
	variable.SetVariable(instructionData, "HeightOutput", h)
	variable.SetVariable(instructionData, "WidthOutput", w)
	finished <- true
	return -1
}

// PartScreenShot Take a screenshot from a part of the screen
func PartScreenShot(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Taking a screen shot from a part of the screen")

	path, err := variable.Keys{VarName: "PathVarName", IsVarName: "PathIsVar", Name: "Path"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	x, err := variable.Keys{VarName: "XVarName", IsVarName: "XIsVar", Name: "X"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	y, err := variable.Keys{VarName: "YVarName", IsVarName: "YIsVar", Name: "Y"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	w, err := variable.Keys{VarName: "WVarName", IsVarName: "WIsVar", Name: "W"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	h, err := variable.Keys{VarName: "HVarName", IsVarName: "HIsVar", Name: "H"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.SaveCapture(path.(string), x.(int), y.(int), w.(int), h.(int))
	finished <- true
	return -1
}

// ScreenShot Save a screen shot
func ScreenShot(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Taking a screen shot")

	path, err := variable.Keys{VarName: "PathVarName", IsVarName: "PathIsVar", Name: "Path"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.SaveCapture(path.(string))
	finished <- true
	return -1
}
