package mouse

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/go-vgo/robotgo"
)

// Click Simulate a user click
func Click(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clicking mouse  ...")

	robotgo.Click(instructionData.FieldByName("MouseButtonName").Interface().(string))
	finished <- true
	return -1
}

// Drag Simulate a user drag
func Drag(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clicking mouse  ...")

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

	robotgo.MouseToggle("down")
	robotgo.Drag(x.(int), y.(int))
	robotgo.MouseToggle("up")
	finished <- true
	return -1
}

// DragSmooth Simulate a user draggin smoothly
func DragSmooth(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Clicking mouse  ...")

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

	robotgo.MouseToggle("down")
	robotgo.DragSmooth(x.(int), y.(int))
	robotgo.MouseToggle("up")
	finished <- true
	return -1
}

// Move move the mouse
func Move(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Moving mouse ...")

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

	robotgo.MoveMouse(x.(int), y.(int))
	finished <- true
	return -1
}

// MoveSmooth move the mouse smoothly
func MoveSmooth(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Moving mouse ...")

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

	robotgo.MoveSmooth(x.(int), y.(int))
	finished <- true
	return -1
}

// Scroll scroll the mouse
func Scroll(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Scrolling mouse ...")

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

	robotgo.Scroll(x.(int), y.(int))
	finished <- true
	return -1
}
