package mouse

import (
	"gotomate/fiber/variable"
	"gotomate/log"

	"github.com/go-vgo/robotgo"
)

// Click Simulate a user click
func Click(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Clicking the mouse")

	btnName, err := variable.Keys{Name: "MouseButtonName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	robotgo.Click(btnName.(string))
	finished <- true
	return -1
}

// Drag Simulate a user drag
func Drag(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Draging the mouse")

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

	robotgo.MouseToggle("down")
	robotgo.Drag(x.(int), y.(int))
	robotgo.MouseToggle("up")
	finished <- true
	return -1
}

// DragSmooth Simulate a user draggin smoothly
func DragSmooth(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Draging smoothly the mouse")

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

	robotgo.MouseToggle("down")
	robotgo.DragSmooth(x.(int), y.(int))
	robotgo.MouseToggle("up")
	finished <- true
	return -1
}

// Move move the mouse
func Move(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Moving the mouse")

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

	robotgo.MoveMouse(x.(int), y.(int))
	finished <- true
	return -1
}

// MoveSmooth move the mouse smoothly
func MoveSmooth(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Moving smoothly the mouse")

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

	robotgo.MoveSmooth(x.(int), y.(int))
	finished <- true
	return -1
}

// Scroll scroll the mouse
func Scroll(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Scrolling the mouse")

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

	robotgo.Scroll(x.(int), y.(int))
	finished <- true
	return -1
}
